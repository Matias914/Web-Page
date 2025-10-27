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

create_review() {
    USER_ID="$1"
    MOVIE_ID="$2"
    URL="http://localhost:8080/api/users/$USER_ID/reviews"
    DATA='{"movie_id": '$MOVIE_ID', "comment": "A review to be deleted"}'
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
===== INICIANDO PRUEBAS PARA DELETE /api/users/{user_id}/reviews/{movie_id} ====="

# Setup
USER_ID=$(create_user "" '{"username": "delete_review_user", "password": "p", "mail": "deletereview@test.com"}')
MOVIE_ID=$(create_movie "" '{"title": "Delete Review Movie", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
MOVIE_ID_NO_REVIEW=$(create_movie "" '{"title": "Delete Review Movie 2", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
create_review "$USER_ID" "$MOVIE_ID"

# Caso 1: Borrado exitoso
run_delete_test "Borrado exitoso" "http://localhost:8080/api/users/$USER_ID/reviews/$MOVIE_ID" 204

# Caso 2: Review no encontrada
run_delete_test "Review no encontrada" "http://localhost:8080/api/users/$USER_ID/reviews/$MOVIE_ID_NO_REVIEW" 404

# Caso 3: Usuario no encontrado
run_delete_test "Usuario no encontrado" "http://localhost:8080/api/users/999999/reviews/$MOVIE_ID" 404

# Limpieza
delete_user "$USER_ID"
delete_movie "$MOVIE_ID"
delete_movie "$MOVIE_ID_NO_REVIEW"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para DELETE /api/users/{user_id}/reviews/{movie_id} pasaron exitosamente. ✅
"
fi
