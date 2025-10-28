#!/bin/bash

FAILED=0

run_post_test() {
    URL="$2"
    DATA="$3"
    EXPECTED_STATUS="$4"
    COMMAND="curl -s -o /dev/null -w '%{http_code}' -X POST -H 'Content-Type: application/json' -d '$DATA' $URL"
    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")

    if [[ ! "$HTTP_STATUS" -eq "$EXPECTED_STATUS" ]]; then
        echo "----------------------------------------------------"
        echo "Test: $1"
        echo "Comando: $COMMAND"
        echo "FALLÓ - Se esperaba $EXPECTED_STATUS ❌"
        echo "Resultado: $HTTP_STATUS"
        FAILED=1
    fi
}

create_celebrity() {
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="$APP_TEST_URL/api/celebrities"
    DATA="$2"

    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
       CELEBRITY_ID=$(echo "$BODY" | jq -r '.id')

       if [ -z "$CELEBRITY_ID" ] || [ "$CELEBRITY_ID" == "null" ]; then
            echo "Error: Creación OK (201), pero ID no encontrado en la respuesta JSON. Body: $BODY" >&2
            FAILED=1
       fi

       echo "Celebridad creada con ID: $CELEBRITY_ID" >&2
       echo "$CELEBRITY_ID"
       return 0
    else
       echo "Error al crear la celebridad. Status: $HTTP_STATUS, Body: $BODY" >&2
       FAILED=1
    fi
}

delete_celebrity() {
    CELEBRITY_ID="$1"
    echo "----------------------------------------------------"
    echo "Acción: Eliminando celebridad con ID $CELEBRITY_ID"
    URL="$APP_TEST_URL/api/celebrities/$CELEBRITY_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Celebridad eliminada exitosamente"
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA POST /api/celebrities ===="

# Caso 1: JSON Inválido
run_post_test "JSON Inválido" "$APP_TEST_URL/api/celebrities" '{"name": "Test celebrity",,}' 500

# Caso 2: Campo requerido faltante (name)
run_post_test "Campo requerido faltante (name)" "$APP_TEST_URL/api/celebrities" '{"birth_date": "1990-01-01T00:00:00Z"}' 400

# Caso 3: Crear una celebridad para probar duplicados
CELEBRITY_DATA='{"name": "duplicate test celebrity", "birth_date": "1985-05-15T00:00:00Z"}'
CELEBRITY_ID_1=$(create_celebrity "Crear celebridad para prueba de duplicados" "$CELEBRITY_DATA")
CELEBRITY_ID_2=$(create_celebrity "Crear celebridad para prueba de duplicados" "$CELEBRITY_DATA")

# Caso 4: Celebridad duplicada
run_post_test "Celebridad duplicada" "$APP_TEST_URL/api/celebrities" "$CELEBRITY_DATA" 409

if [ ! -z "$CELEBRITY_ID" ]; then
    delete_celebrity "$CELEBRITY_ID"
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo "----------------------------------------------------" >&2
    echo -e "Todas las pruebas para POST /api/celebrities pasaron exitosamente. ✅
"
fi
