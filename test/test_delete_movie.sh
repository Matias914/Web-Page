#!/bin/bash

FAILED=0

run_delete_test() {

    TEST_NAME="$1"
    URL="$2"
    EXPECTED_STATUS="$3"
    COMMAND="curl -s -o /dev/null -w '%{http_code}' -X DELETE '$URL'"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [[ "$HTTP_STATUS" -ne "$EXPECTED_STATUS" ]]; then
        echo "----------------------------------------------------"
        echo "Test: $TEST_NAME"
        echo "FALLÓ - Se esperaba $EXPECTED_STATUS ❌"
        echo "Comando: $COMMAND" # Ahora se imprime el comando correctamente
        echo "Resultado: $HTTP_STATUS"
        FAILED=1
    fi
}

create_movie() {
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="$APP_TEST_URL/api/movies"
    DATA="$2"

    BODY_AND_STATUS=$(curl -s -w "\n%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '\n')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
       MOVIE_ID=$(echo "$BODY" | jq -r '.id')

       if [ -z "$MOVIE_ID" ] || [ "$MOVIE_ID" == "null" ]; then
            echo "Error: Creación OK (201), pero ID no encontrado en la respuesta JSON. Body: $BODY" >&2
            FAILED=1
       fi

       echo "Película creada con ID: $MOVIE_ID" >&2
       echo "$MOVIE_ID"
       echo "----------------------------------------------------" >&2
       return 0
    else
       echo "Error al crear la película. Status: $HTTP_STATUS, Body: $BODY" >&2
       echo "----------------------------------------------------" >&2
       FAILED=1
    fi
}

echo -e "\n===== INICIANDO PRUEBAS PARA DELETE /api/movies/{id} ====="

# Caso 1: ID no es un número
run_delete_test "ID no es un número" "$APP_TEST_URL/api/movies/abc" 500

# Caso 2: Película no encontrada
run_delete_test "Película no encontrada" "$APP_TEST_URL/api/movies/999999" 404

# Caso 3: Borrado exitoso
MOVIE_ID=$(create_movie "Crear película para borrar" '{"title": "Delete Test Movie 3", "synopsis": "Synopsis", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 130}')
if [ ! -z "$MOVIE_ID" ]; then
    run_delete_test "Borrado exitoso" "$APP_TEST_URL/api/movies/$MOVIE_ID" 204
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para DELETE /api/movies/{id} pasaron exitosamente. ✅ \n"
fi
