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

create_genre() {
    URL="http://localhost:8080/api/genres"
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

delete_movie() {
    MOVIE_ID="$1"
    URL="http://localhost:8080/api/movies/$MOVIE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

delete_genre() {
    GENRE_ID="$1"
    URL="http://localhost:8080/api/genres/$GENRE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA POST /api/movies/{id}/categories ===="

# Setup: Crear película y género para las pruebas
MOVIE_ID=$(create_movie "" '{"title": "Post Movie Category", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
GENRE_ID=$(create_genre "" '{"name": "Action"}')
GENRE_ID_2=$(create_genre "" '{"name": "Comedy"}')

if [ -z "$MOVIE_ID" ] || [ -z "$GENRE_ID" ] || [ -z "$GENRE_ID_2" ]; then
    echo "Error: No se pudieron crear la película o los géneros necesarios para las pruebas de categorías."
    exit 1
fi

# Caso 1: Categoría exitosa
run_post_test "Categoría exitosa" "http://localhost:8080/api/movies/$MOVIE_ID/categories" '{"genre_id": '$GENRE_ID'}' 204

# Caso 2: Categoría duplicada
run_post_test "Categoría duplicada" "http://localhost:8080/api/movies/$MOVIE_ID/categories" '{"genre_id": '$GENRE_ID'}' 409

# Caso 3: Película no encontrada
run_post_test "Película no encontrada" "http://localhost:8080/api/movies/99999/categories" '{"genre_id": '$GENRE_ID'}' 404

# Caso 4: Género no encontrado
run_post_test "Género no encontrado" "http://localhost:8080/api/movies/$MOVIE_ID/categories" '{"genre_id": 99999}' 404

# Caso 5: Falla de validación (genre_id vacío)
run_post_test "Falla de validación (genre_id vacío)" "http://localhost:8080/api/movies/$MOVIE_ID/categories" '{}' 400

# Limpieza
delete_movie "$MOVIE_ID"
delete_genre "$GENRE_ID"
delete_genre "$GENRE_ID_2"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para POST /api/movies/{id}/categories pasaron exitosamente. ✅
"
fi
