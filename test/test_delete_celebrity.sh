#!/bin/bash

FAILED=0

# Función para hacer una solicitud DELETE y mostrar los resultados
run_delete_test() {
    TEST_NAME="$1"
    URL="$2"
    EXPECTED_STATUS="$3"
    COMMAND="curl -s -o /dev/null -w '%{http_code}' -X DELETE '$URL'"
    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [[ "$HTTP_STATUS" -ne "$EXPECTED_STATUS" ]]; then
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
    URL="http://localhost:8080/api/celebrities"
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

echo -e "
===== INICIANDO PRUEBAS PARA DELETE /api/celebrities/{id} ====="

# Caso 1: ID no es un número
run_delete_test "ID no es un número" "http://localhost:8080/api/celebrities/abc" 500

# Caso 2: Celebridad no encontrada
run_delete_test "Celebridad no encontrada" "http://localhost:8080/api/celebrities/999999" 404

# Caso 3: Borrado exitoso
CELEBRITY_ID=$(create_celebrity "Crear celebridad para borrar" '{"name": "Delete Test Celebrity", "birth_date": "1998-08-08T00:00:00Z"}')
if [ ! -z "$CELEBRITY_ID" ]; then
    run_delete_test "Borrado exitoso" "http://localhost:8080/api/celebrities/$CELEBRITY_ID" 204
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para DELETE /api/celebrities/{id} pasaron exitosamente. ✅
"
fi
