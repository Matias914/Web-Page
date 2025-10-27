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

create_role() {
    MOVIE_ID="$1"
    CELEBRITY_ID="$2"
    URL="http://localhost:8080/api/movies/$MOVIE_ID/roles"
    DATA='{"celebrity_id": '$CELEBRITY_ID', "role": "Role to be deleted"}'
    curl -s -o /dev/null -X POST -H "Content-Type: application/json" -d "$DATA" "$URL"
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
===== INICIANDO PRUEBAS PARA DELETE /api/movies/{movie_id}/roles/{celebrity_id} ====="

# Setup
MOVIE_ID=$(create_movie "" '{"title": "Delete Movie Role", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
CELEBRITY_ID=$(create_celebrity "" '{"name": "Delete Movie Role Celebrity", "birth_date": "1990-01-01T00:00:00Z"}')
CELEBRITY_ID_NO_ROLE=$(create_celebrity "" '{"name": "Delete Movie Role Celebrity 2", "birth_date": "1991-01-01T00:00:00Z"}')
create_role "$MOVIE_ID" "$CELEBRITY_ID"

# Caso 1: Borrado exitoso
run_delete_test "Borrado exitoso" "http://localhost:8080/api/movies/$MOVIE_ID/roles/$CELEBRITY_ID" 204

# Caso 2: Rol no encontrado
run_delete_test "Rol no encontrado" "http://localhost:8080/api/movies/$MOVIE_ID/roles/$CELEBRITY_ID_NO_ROLE" 404

# Caso 3: Película no encontrada
run_delete_test "Película no encontrada" "http://localhost:8080/api/movies/999999/roles/$CELEBRITY_ID" 404

# Limpieza
delete_movie "$MOVIE_ID"
delete_celebrity "$CELEBRITY_ID"
delete_celebrity "$CELEBRITY_ID_NO_ROLE"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para DELETE /api/movies/{movie_id}/roles/{celebrity_id} pasaron exitosamente. ✅
"
fi
