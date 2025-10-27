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
    DATA='{"celebrity_id": '$CELEBRITY_ID', "role": "Initial Role"}'
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
===== INICIANDO PRUEBAS PARA PUT /api/movies/{movie_id}/roles/{celebrity_id} ====="

# Setup
MOVIE_ID=$(create_movie "" '{"title": "Put Movie Role", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')
CELEBRITY_ID=$(create_celebrity "" '{"name": "Put Movie Role Celebrity", "birth_date": "1990-01-01T00:00:00Z"}')
create_role "$MOVIE_ID" "$CELEBRITY_ID"

# Caso 1: Actualización exitosa
run_put_test "Actualización exitosa" "http://localhost:8080/api/movies/$MOVIE_ID/roles/$CELEBRITY_ID" '{"role": "Updated Role"}' 200

# Caso 2: Rol no encontrado
run_put_test "Rol no encontrado" "http://localhost:8080/api/movies/$MOVIE_ID/roles/999999" '{"role": "Some Role"}' 404

# Caso 3: Falla de validación (role vacío)
run_put_test "Falla de validación" "http://localhost:8080/api/movies/$MOVIE_ID/roles/$CELEBRITY_ID" '{"role": ""}' 400

# Limpieza
delete_movie "$MOVIE_ID"
delete_celebrity "$CELEBRITY_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para PUT /api/movies/{movie_id}/roles/{celebrity_id} pasaron exitosamente. ✅
"
fi
