#!/bin/bash

FAILED=0

# Función para hacer una solicitud GET y mostrar los resultados
run_get_test() {
    URL="$2"
    EXPECTED_STATUS="$3"
    COMMAND="curl -s -o /dev/null -w '%{http_code}' '$URL'"
    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

    if [[ ! "$HTTP_STATUS" -eq "$EXPECTED_STATUS" ]]; then
        echo "Test: $1"
        echo "FALLÓ - Se esperaba $EXPECTED_STATUS ❌"
        echo "Comando: $COMMAND"
        echo "Resultado: $HTTP_STATUS"
        echo "----------------------------------------------------"
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

# Función para eliminar una celebridad
delete_celebrity() {
    CELEBRITY_ID="$1"
    echo "----------------------------------------------------"
    echo "Acción: Eliminando celebridad con ID $CELEBRITY_ID"
    URL="$APP_TEST_URL/api/celebrities/$CELEBRITY_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Celebridad eliminada exitosamente"
        echo "----------------------------------------------------" >&2
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA GET /api/celebrities/{id} ====="

# Caso 1: ID no es un número
run_get_test "ID no es un número" "$APP_TEST_URL/api/celebrities/abc" 500

# Caso 2: Celebridad no encontrada
run_get_test "Celebridad no encontrada" "$APP_TEST_URL/api/celebrities/999999" 404

# Caso 3: Obtención exitosa
CELEBRITY_ID=$(create_celebrity "Crear celebridad para obtener" '{"name": "Get Test Celebrity", "birth_date": "1992-02-02T00:00:00Z"}')
if [ ! -z "$CELEBRITY_ID" ]; then
    run_get_test "Obtención exitosa" "$APP_TEST_URL/api/celebrities/$CELEBRITY_ID" 200
    delete_celebrity "$CELEBRITY_ID"
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para GET /api/celebrities/{id} pasaron exitosamente. ✅
"
fi
