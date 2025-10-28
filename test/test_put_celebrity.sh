#!/bin/bash

FAILED=0

# Función para hacer una solicitud PUT y mostrar los resultados
run_put_test() {
    TEST_NAME="$1"
    URL="$2"
    DATA="$3"
    EXPECTED_STATUS="$4"
    COMMAND="curl -s -o /dev/null -w '%{http_code}' -X PUT -H 'Content-Type: application/json' -d '$DATA' '$URL'"
    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X PUT -H "Content-Type: application/json" -d "$DATA" "$URL")

    if [[ ! "$HTTP_STATUS" -eq "$EXPECTED_STATUS" ]]; then
        echo "----------------------------------------------------"
        echo "Test: $TEST_NAME"
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

# Función para eliminar una celebridad
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
===== INICIANDO PRUEBAS PARA PUT /api/celebrities/{id} ====="

# Caso 1: ID no es un número
run_put_test "ID no es un número" "$APP_TEST_URL/api/celebrities/abc" '{}' 500

# Caso 2: Celebridad no encontrada
run_put_test "Celebridad no encontrada" "$APP_TEST_URL/api/celebrities/999999" '{"name": "Not Found Test", "birth_date": "1999-01-01T00:00:00Z"}' 404

# Caso 3: JSON Inválido
CELEBRITY_ID_INV_JSON=$(create_celebrity "Crear celebridad para JSON inválido" '{"name": "Invalid JSON Celebrity", "birth_date": "1993-03-03T00:00:00Z"}')
if [ ! -z "$CELEBRITY_ID_INV_JSON" ]; then
    run_put_test "JSON Inválido" "$APP_TEST_URL/api/celebrities/$CELEBRITY_ID_INV_JSON" '{"name": "Updated Name",,}' 500
    delete_celebrity "$CELEBRITY_ID_INV_JSON"
fi

# Caso 4: Falla de validación (campo faltante)
CELEBRITY_ID_VALIDATION=$(create_celebrity "Crear celebridad para falla de validación" '{"name": "Validation Fail Celebrity", "birth_date": "1994-04-04T00:00:00Z"}')
if [ ! -z "$CELEBRITY_ID_VALIDATION" ]; then
    run_put_test "Falla de validación" "$APP_TEST_URL/api/celebrities/$CELEBRITY_ID_VALIDATION" '{"birth_date": "1994-04-04T00:00:00Z"}' 400
    delete_celebrity "$CELEBRITY_ID_VALIDATION"
fi

# Caso 5: Actualización exitosa
CELEBRITY_ID_SUCCESS=$(create_celebrity "Crear celebridad para actualizar" '{"name": "Update Success Original", "birth_date": "1995-05-05T00:00:00Z"}')
if [ ! -z "$CELEBRITY_ID_SUCCESS" ]; then
    run_put_test "Actualización exitosa" "$APP_TEST_URL/api/celebrities/$CELEBRITY_ID_SUCCESS" '{"name": "Update Success Final", "birth_date": "1996-06-06T00:00:00Z"}' 200
    delete_celebrity "$CELEBRITY_ID_SUCCESS"
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo "----------------------------------------------------" >&2
    echo -e "Todas las pruebas para PUT /api/celebrities/{id} pasaron exitosamente. ✅
"
fi
