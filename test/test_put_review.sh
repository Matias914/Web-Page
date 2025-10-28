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

create_movie() {
    URL="$APP_TEST_URL/api/movies"
    DATA="$2"
    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')
    if [ "$HTTP_STATUS" -eq 201 ]; then
       MOVIE_ID=$(echo "$BODY" | jq -r '.id')
       echo "$MOVIE_ID"
    fi
}

create_review() {
    USER_ID="$1"
    MOVIE_ID="$2"
    URL="$APP_TEST_URL/api/users/$USER_ID/reviews"
    DATA='{"movie_id": '$MOVIE_ID', "comment": "Initial comment"}'
    curl -s -o /dev/null -X POST -H "Content-Type: application/json" -d "$DATA" "$URL"
}

delete_user() {
    USER_ID="$1"
    URL="$APP_TEST_URL/api/users/$USER_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

delete_movie() {
    MOVIE_ID="$1"
    URL="$APP_TEST_URL/api/movies/$MOVIE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA PUT /api/users/{user_id}/reviews/{movie_id} ====="

# Setup
USER_ID=$(create_user "" '{"username": "put_review_user", "password": "p", "mail": "putreview@test.com"}')
MOVIE_ID=$(create_movie "" '{"title": "Put Review Movie", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
create_review "$USER_ID" "$MOVIE_ID"

# Caso 1: Actualización exitosa
run_put_test "Actualización exitosa" "$APP_TEST_URL/api/users/$USER_ID/reviews/$MOVIE_ID" '{"comment": "Updated comment"}' 200

# Caso 2: Review no encontrada
run_put_test "Review no encontrada" "$APP_TEST_URL/api/users/$USER_ID/reviews/999999" '{"comment": "some comment"}' 404

# Caso 3: Falla de validación (comment vacío)
run_put_test "Falla de validación" "$APP_TEST_URL/api/users/$USER_ID/reviews/$MOVIE_ID" '{"comment": ""}' 400

# Limpieza
delete_user "$USER_ID"
delete_movie "$MOVIE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para PUT /api/users/{user_id}/reviews/{movie_id} pasaron exitosamente. ✅
"
fi
