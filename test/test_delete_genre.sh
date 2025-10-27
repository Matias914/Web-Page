#!/bin/bash

FAILED=0

run_delete_test() {

    TEST_NAME="$1"
    URL="$2"
    EXPECTED_STATUS="$3"
    COMMAND="curl -s -o /dev/null -w '%{http_code}' -X DELETE '$URL'"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [[ "$HTTP_STATUS" -ne "$EXPECTED_STATUS" ]]; then
        echo "----------------------------------------------------"
        echo "Test: $TEST_NAME"
        echo "Comando: $COMMAND" # Ahora se imprime el comando correctamente
        echo "FALLÓ - Se esperaba $EXPECTED_STATUS ❌"
        echo "Resultado: $HTTP_STATUS"
        FAILED=1
    fi
}

create_genre() {
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="http://localhost:8080/api/genres"
    DATA="$2"

    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
       GENRE_ID=$(echo "$BODY" | jq -r '.id')

       if [ -z "$GENRE_ID" ] || [ "$GENRE_ID" == "null" ]; then
            echo "Error: Creación OK (201), pero ID no encontrado en la respuesta JSON. Body: $BODY" >&2
            FAILED=1
       fi

       echo "Género creado con ID: $GENRE_ID" >&2
       echo "$GENRE_ID"
       echo "----------------------------------------------------" >&2
       return 0
    else
       echo "Error al crear el género. Status: $HTTP_STATUS, Body: $BODY" >&2
       echo "----------------------------------------------------" >&2
       FAILED=1
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA DELETE /api/genres/{id} ====="

# Caso 1: ID no es un número
run_delete_test "ID no es un número" "http://localhost:8080/api/genres/abc" 400

# Caso 2: Género no encontrado
run_delete_test "Género no encontrado" "http://localhost:8080/api/genres/9999" 404

# Caso 3: Borrado exitoso
GENRE_ID=$(create_genre "Crear género para borrar" '{"name": "Delete Test Genre"}')
if [ ! -z "$GENRE_ID" ]; then
    run_delete_test "Borrado exitoso" "http://localhost:8080/api/genres/$GENRE_ID" 204
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para DELETE /api/genres/{id} pasaron exitosamente. ✅
"
fi