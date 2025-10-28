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

create_user() {
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="$APP_TEST_URL/api/users"
    DATA="$2"

    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
       USER_ID=$(echo "$BODY" | jq -r '.id')

       if [ -z "$USER_ID" ] || [ "$USER_ID" == "null" ]; then
            echo "Error: Creación OK (201), pero ID no encontrado en la respuesta JSON. Body: $BODY" >&2
            FAILED=1
       fi

       echo "Usuario creado con ID: $USER_ID" >&2
       echo "$USER_ID"
       return 0
    else
       echo "Error al crear el usuario. Status: $HTTP_STATUS, Body: $BODY" >&2
       FAILED=1
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA DELETE /api/users/{id} ====="

# Caso 1: ID no es un número
run_delete_test "ID no es un número" "$APP_TEST_URL/api/users/abc" 500

# Caso 2: Usuario no encontrado
run_delete_test "Usuario no encontrado" "$APP_TEST_URL/api/users/999999" 404

# Caso 3: Borrado exitoso
USER_ID=$(create_user "Crear usuario para borrar" '{"username": "delete_me", "password": "p", "mail": "delete@me.com"}')
if [ ! -z "$USER_ID" ]; then
    run_delete_test "Borrado exitoso" "$APP_TEST_URL/api/users/$USER_ID" 204
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para DELETE /api/users/{id} pasaron exitosamente. ✅
"
fi
