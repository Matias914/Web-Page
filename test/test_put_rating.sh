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
    URL="http://localhost:8080/api/users"
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
    URL="http://localhost:8080/api/movies"
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

create_rating() {
    USER_ID="$1"
    MOVIE_ID="$2"
    URL="http://localhost:8080/api/users/$USER_ID/ratings"
    DATA='{"movie_id": '$MOVIE_ID', "rating": 8}'
    curl -s -o /dev/null -X POST -H "Content-Type: application/json" -d "$DATA" "$URL"
}

delete_user() {
    USER_ID="$1"
    URL="http://localhost:8080/api/users/$USER_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

delete_movie() {
    MOVIE_ID="$1"
    URL="http://localhost:8080/api/movies/$MOVIE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA PUT /api/users/{user_id}/ratings/{movie_id} ====="

# Setup
USER_ID=$(create_user "" '{"username": "put_rating_user", "password": "p", "mail": "putrating@test.com"}')
MOVIE_ID=$(create_movie "" '{"title": "Put Rating Movie", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
create_rating "$USER_ID" "$MOVIE_ID"

# Caso 1: Actualización exitosa
run_put_test "Actualización exitosa" "http://localhost:8080/api/users/$USER_ID/ratings/$MOVIE_ID" '{"rating": 5}' 200

# Caso 2: Rating no encontrado
run_put_test "Rating no encontrado" "http://localhost:8080/api/users/$USER_ID/ratings/9999" '{"rating": 5}' 404

# Caso 3: Falla de validación (rating > 10)
run_put_test "Falla de validación" "http://localhost:8080/api/users/$USER_ID/ratings/$MOVIE_ID" '{"rating": 11}' 400

# Limpieza
delete_user "$USER_ID"
delete_movie "$MOVIE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para PUT /api/users/{user_id}/ratings/{movie_id} pasaron exitosamente. ✅
"
fi
