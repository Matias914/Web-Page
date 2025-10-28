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

create_genre() {
    URL="$APP_TEST_URL/api/genres"
    DATA="$2"
    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')
    if [ "$HTTP_STATUS" -eq 201 ]; then
       GENRE_ID=$(echo "$BODY" | jq -r '.id')
       echo "$GENRE_ID"
    fi
}

create_category() {
    MOVIE_ID="$1"
    GENRE_ID="$2"
    URL="$APP_TEST_URL/api/movies/$MOVIE_ID/categories"
    DATA='{"genre_id": '$GENRE_ID'}'
    curl -s -o /dev/null -X POST -H "Content-Type: application/json" -d "$DATA" "$URL"
}

delete_movie() {
    MOVIE_ID="$1"
    URL="$APP_TEST_URL/api/movies/$MOVIE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

delete_genre() {
    GENRE_ID="$1"
    URL="$APP_TEST_URL/api/genres/$GENRE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA DELETE /api/movies/{movie_id}/categories/{genre_id} ====="

# Setup
MOVIE_ID=$(create_movie "" '{"title": "Delete Movie Category", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
GENRE_ID=$(create_genre "" '{"name": "Delete Category Genre"}')
GENRE_ID_NO_ASSOC=$(create_genre "" '{"name": "Delete No Assoc Genre"}')
create_category "$MOVIE_ID" "$GENRE_ID"

# Caso 1: Borrado exitoso
run_delete_test "Borrado exitoso" "$APP_TEST_URL/api/movies/$MOVIE_ID/categories/$GENRE_ID" 204

# Caso 2: Categoría no encontrada (no asociada)
run_delete_test "Categoría no encontrada" "$APP_TEST_URL/api/movies/$MOVIE_ID/categories/$GENRE_ID_NO_ASSOC" 404

# Caso 3: Película no encontrada
run_delete_test "Película no encontrada" "$APP_TEST_URL/api/movies/999999/categories/$GENRE_ID" 404

# Limpieza
delete_movie "$MOVIE_ID"
delete_genre "$GENRE_ID"
delete_genre "$GENRE_ID_NO_ASSOC"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para DELETE /api/movies/{movie_id}/categories/{genre_id} pasaron exitosamente. ✅
"
fi
