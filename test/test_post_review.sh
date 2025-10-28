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
===== INICIANDO PRUEBAS PARA POST /api/users/{id}/reviews ===="

# Setup: Crear usuario y película para las pruebas
USER_ID=$(create_user "" '{"username": "review_user", "password": "p", "mail": "review@test.com"}')
MOVIE_ID=$(create_movie "" '{"title": "Review Movie", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')

if [ -z "$USER_ID" ] || [ -z "$MOVIE_ID" ]; then
    echo "Error: No se pudieron crear el usuario o la película necesarios para las pruebas de reviews."
    exit 1
fi

# Caso 1: Review exitosa
run_post_test "Review exitosa" "$APP_TEST_URL/api/users/$USER_ID/reviews" '{"movie_id": '$MOVIE_ID', "comment": "This is a great movie!"}' 201

# Caso 2: Review duplicada
run_post_test "Review duplicada" "$APP_TEST_URL/api/users/$USER_ID/reviews" '{"movie_id": '$MOVIE_ID', "comment": "Another comment."}' 409

# Caso 3: Película no encontrada
run_post_test "Película no encontrada" "$APP_TEST_URL/api/users/$USER_ID/reviews" '{"movie_id": 99999, "comment": "A comment."}' 404

# Caso 4: Falla de validación (comment vacío)
run_post_test "Falla de validación (comment vacío)" "$APP_TEST_URL/api/users/$USER_ID/reviews" '{"movie_id": '$MOVIE_ID', "comment": ""}' 400

# Limpieza
delete_user "$USER_ID"
delete_movie "$MOVIE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para POST /api/users/{id}/reviews pasaron exitosamente. ✅
"
fi
