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

create_user() {
    URL="$APP_TEST_URL/api/users"
    DATA="$2"
    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')
    if [ "$HTTP_STATUS" -eq 201 ]; then
       USER_ID=$(echo "$BODY" | jq -r '.id')
       echo "$USER_ID"
    fi
}

delete_user() {
    USER_ID="$1"
    URL="$APP_TEST_URL/api/users/$USER_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA GET /api/users/{id}/reviews ====="

# Setup: Crear usuario para las pruebas
USER_ID=$(create_user "" '{"username": "get_reviews_user", "password": "p", "mail": "getreviews@test.com"}')

if [ -z "$USER_ID" ]; then
    echo "Error: No se pudo crear el usuario necesario para las pruebas."
    exit 1
fi

# Caso 1: Obtención exitosa (incluso si está vacía)
run_get_test "Obtención exitosa" "$APP_TEST_URL/api/users/$USER_ID/reviews?page=1&rows=5" 200

# Caso 2: ID de usuario no es un número
run_get_test "ID de usuario no es un número" "$APP_TEST_URL/api/users/abc/reviews?page=1&rows=5" 500

# Caso 3: Usuario no encontrado
run_get_test "Usuario no encontrado" "$APP_TEST_URL/api/users/999999/reviews?page=1&rows=5" 404

# Caso 4: Paginación inválida
run_get_test "Paginación inválida" "$APP_TEST_URL/api/users/$USER_ID/reviews?page=abc&rows=5" 500

# Limpieza
delete_user "$USER_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para GET /api/users/{id}/reviews pasaron exitosamente. ✅
"
fi
