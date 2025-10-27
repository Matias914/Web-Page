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

delete_movie() {
    MOVIE_ID="$1"
    URL="http://localhost:8080/api/movies/$MOVIE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA GET /api/movies/{id}/reviews ====="

# Setup: Crear película para las pruebas
MOVIE_ID=$(create_movie "" '{"title": "Get Movie Reviews", "synopsis": "s", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 100}')

if [ -z "$MOVIE_ID" ]; then
    echo "Error: No se pudo crear la película necesaria para las pruebas."
    exit 1
fi

# Caso 1: Obtención exitosa (incluso si está vacía)
run_get_test "Obtención exitosa" "http://localhost:8080/api/movies/$MOVIE_ID/reviews?page=1&rows=5" 200

# Caso 2: ID de película no es un número
run_get_test "ID de película no es un número" "http://localhost:8080/api/movies/abc/reviews?page=1&rows=5" 500

# Caso 3: Película no encontrada
run_get_test "Película no encontrada" "http://localhost:8080/api/movies/999999/reviews?page=1&rows=5" 404

# Caso 4: Paginación inválida
run_get_test "Paginación inválida" "http://localhost:8080/api/movies/$MOVIE_ID/reviews?page=abc&rows=5" 500

# Limpieza
delete_movie "$MOVIE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para GET /api/movies/{id}/reviews pasaron exitosamente. ✅
"
fi
