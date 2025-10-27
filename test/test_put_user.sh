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

create_user() {
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="http://localhost:8080/api/users"
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
    URL="http://localhost:8080/api/users/$USER_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Usuario eliminado exitosamente"
        echo "----------------------------------------------------" >&2
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA PUT /api/users/{id} ====="

# Caso 1: ID no es un número
run_put_test "ID no es un número" "http://localhost:8080/api/users/abc" '{}' 500

# Caso 2: Usuario no encontrado
run_put_test "Usuario no encontrado" "http://localhost:8080/api/users/999999" '{"username": "not_found", "mail": "not@found.com"}' 404

# Caso 3: JSON Inválido
USER_ID_INV_JSON=$(create_user "Crear usuario para JSON inválido" '{"username": "inv_json_user", "password": "p", "mail": "ij@test.com"}')
if [ ! -z "$USER_ID_INV_JSON" ]; then
    run_put_test "JSON Inválido" "http://localhost:8080/api/users/$USER_ID_INV_JSON" '{"username": "updated",,}' 500
    delete_user "$USER_ID_INV_JSON"
fi

# Caso 4: Conflicto de duplicados
USER_ID_CONFLICT_1=$(create_user "Crear usuario 1 para conflicto" '{"username": "conflict1", "password": "p", "mail": "conflict1@test.com"}')
USER_ID_CONFLICT_2=$(create_user "Crear usuario 2 para conflicto" '{"username": "conflict2", "password": "p", "mail": "conflict2@test.com"}')
if [ ! -z "$USER_ID_CONFLICT_1" ] && [ ! -z "$USER_ID_CONFLICT_2" ]; then
    run_put_test "Conflicto de duplicados" "http://localhost:8080/api/users/$USER_ID_CONFLICT_2" '{"username": "conflict1", "mail": "c2@test.com"}' 409
    delete_user "$USER_ID_CONFLICT_1"
    delete_user "$USER_ID_CONFLICT_2"
fi

# Caso 5: Actualización exitosa
USER_ID_SUCCESS=$(create_user "Crear usuario para actualizar" '{"username": "update_me", "password": "p", "mail": "update@me.com"}')
if [ ! -z "$USER_ID_SUCCESS" ]; then
    run_put_test "Actualización exitosa" "http://localhost:8080/api/users/$USER_ID_SUCCESS" '{"username": "updated_user", "mail": "updated@me.com"}' 200
    delete_user "$USER_ID_SUCCESS"
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para PUT /api/users/{id} pasaron exitosamente. ✅
"
fi
