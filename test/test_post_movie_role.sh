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

create_celebrity() {
    URL="http://localhost:8080/api/celebrities"
    DATA="$2"
    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')
    if [ "$HTTP_STATUS" -eq 201 ]; then
       CELEBRITY_ID=$(echo "$BODY" | jq -r '.id')
       echo "$CELEBRITY_ID"
    fi
}

delete_movie() {
    MOVIE_ID="$1"
    URL="http://localhost:8080/api/movies/$MOVIE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

delete_celebrity() {
    CELEBRITY_ID="$1"
    URL="http://localhost:8080/api/celebrities/$CELEBRITY_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA POST /api/movies/{id}/roles ===="

# Setup: Crear película y celebridad para las pruebas
MOVIE_ID=$(create_movie "" '{"title": "Post Movie Role", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
CELEBRITY_ID=$(create_celebrity "" '{"name": "Post Movie Role Celebrity", "birth_date": "1990-01-01T00:00:00Z"}')

if [ -z "$MOVIE_ID" ] || [ -z "$CELEBRITY_ID" ]; then
    echo "Error: No se pudieron crear la película o la celebridad necesarias para las pruebas de roles."
    exit 1
fi

# Caso 1: Rol exitoso
run_post_test "Rol exitoso" "http://localhost:8080/api/movies/$MOVIE_ID/roles" '{"celebrity_id": '$CELEBRITY_ID', "role": "Actor"}' 201

# Caso 2: Rol duplicado
run_post_test "Rol duplicado" "http://localhost:8080/api/movies/$MOVIE_ID/roles" '{"celebrity_id": '$CELEBRITY_ID', "role": "Actor"}' 409

# Caso 3: Película no encontrada
run_post_test "Película no encontrada" "http://localhost:8080/api/movies/99999/roles" '{"celebrity_id": '$CELEBRITY_ID', "role": "Director"}' 404

# Caso 4: Celebridad no encontrada
run_post_test "Celebridad no encontrada" "http://localhost:8080/api/movies/$MOVIE_ID/roles" '{"celebrity_id": 99999, "role": "Director"}' 404

# Caso 5: Falla de validación (role vacío)
run_post_test "Falla de validación (role vacío)" "http://localhost:8080/api/movies/$MOVIE_ID/roles" '{"celebrity_id": '$CELEBRITY_ID', "role": ""}' 400

# Limpieza
delete_movie "$MOVIE_ID"
delete_celebrity "$CELEBRITY_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para POST /api/movies/{id}/roles pasaron exitosamente. ✅
"
fi
