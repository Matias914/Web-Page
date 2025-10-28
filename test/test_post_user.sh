#!/bin/bash

FAILED=0

# Función para hacer una solicitud POST y mostrar los resultados
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

# Función para eliminar un usuario
delete_user() {
    USER_ID="$1"
    echo "----------------------------------------------------"
    echo "Acción: Eliminando usuario con ID $USER_ID"
    URL="$APP_TEST_URL/api/users/$USER_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Usuario eliminado exitosamente"
        echo "----------------------------------------------------" >&2
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA POST /api/users ===="

# Caso 1: JSON Inválido
run_post_test "JSON Inválido" "$APP_TEST_URL/api/users" '{"username": "Test User",,}' 500

# Caso 2: Campo requerido faltante (username)
run_post_test "Campo requerido faltante (username)" "$APP_TEST_URL/api/users" '{"password": "pass", "mail": "a@a.com"}' 400

# Caso 3: Crear un usuario para probar duplicados
USER_DATA='{"username": "duplicate_user", "password": "password123", "mail": "duplicate@test.com"}'
USER_ID=$(create_user "Crear usuario para prueba de duplicados" "$USER_DATA")

# Caso 4: Usuario duplicado (username)
run_post_test "Usuario duplicado (username)" "$APP_TEST_URL/api/users" '{"username": "duplicate_user", "password": "newpass", "mail": "new@test.com"}' 409

# Caso 5: Usuario duplicado (mail)
run_post_test "Usuario duplicado (mail)" "$APP_TEST_URL/api/users" '{"username": "new_user", "password": "newpass", "mail": "duplicate@test.com"}' 409

# Limpieza: Eliminar el usuario creado
if [ ! -z "$USER_ID" ]; then
    delete_user "$USER_ID"
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para POST /api/users pasaron exitosamente. ✅
"
fi
