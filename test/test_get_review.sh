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
    DATA='{"movie_id": '$MOVIE_ID', "comment": "A great movie"}'
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
===== INICIANDO PRUEBAS PARA GET /api/users/{user_id}/reviews/{movie_id} ====="

# Setup
USER_ID=$(create_user "" '{"username": "get_review_user", "password": "p", "mail": "getreview@test.com"}')
MOVIE_ID=$(create_movie "" '{"title": "Get Review Movie", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
create_review "$USER_ID" "$MOVIE_ID"

# Caso 1: Obtención exitosa
run_get_test "Obtención exitosa" "http://localhost:8080/api/users/$USER_ID/reviews/$MOVIE_ID" 200

# Caso 2: Usuario no encontrado
run_get_test "Usuario no encontrado" "http://localhost:8080/api/users/999999/reviews/$MOVIE_ID" 404

# Caso 3: Película no encontrada
run_get_test "Película no encontrada" "http://localhost:8080/api/users/$USER_ID/reviews/999999" 404

# Caso 4: ID de usuario inválido
run_get_test "ID de usuario inválido" "http://localhost:8080/api/users/abc/reviews/$MOVIE_ID" 500

# Limpieza
delete_user "$USER_ID"
delete_movie "$MOVIE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para GET /api/users/{user_id}/reviews/{movie_id} pasaron exitosamente. ✅
"
fi
