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
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="http://localhost:8080/api/movies"
    DATA="$2"

    BODY_AND_STATUS=$(curl -s -w "\n%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '\n')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
       MOVIE_ID=$(echo "$BODY" | jq -r '.id')

       if [ -z "$MOVIE_ID" ] || [ "$MOVIE_ID" == "null" ]; then
            echo "Error FATAL: Creación OK (201), pero ID no encontrado en la respuesta JSON. Body: $BODY" >&2
            exit 1
       fi

       echo "Película creada con ID: $MOVIE_ID" >&2
       echo "$MOVIE_ID"
       return 0
    else
       echo "Error FATAL al crear la película. Status: $HTTP_STATUS, Body: $BODY" >&2
       exit 1
    fi
}

echo -e "\n===== INICIANDO PRUEBAS PARA GET /api/movies/{id} ====="

# Caso 1: ID válido
MOVIE_ID=$(create_movie "Crear película para borrar" '{"title": "Delete Test Movie", "synopsis": "Synopsis", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 130}')
run_get_test "ID válido" "http://localhost:8080/api/movies/$MOVIE_ID" 200

# Caso 2: ID no es un número
run_get_test "ID no es un número" "http://localhost:8080/api/movies/abc" 400

# Caso 3: Película no encontrada (ID muy grande)
run_get_test "Película no encontrada" "http://localhost:8080/api/movies/999999" 404

delete_movie "$MOVIE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "\nTodas las pruebas para GET /api/movies/{id} pasaron exitosamente. ✅ \n"
fi