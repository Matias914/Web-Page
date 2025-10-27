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
        echo "FALLÓ - Se esperaba $EXPECTED_STATUS"
        echo "Resultado: $HTTP_STATUS"
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
            echo "Error: Creación OK (201), pero ID no encontrado en la respuesta JSON. Body: $BODY" >&2
            FAILED=1
       fi

       echo "Película creada con ID: $MOVIE_ID" >&2
       echo "$MOVIE_ID"
       return 0
    else
       echo "Error al crear la película. Status: $HTTP_STATUS, Body: $BODY" >&2
       FAILED=1
    fi
}

# Función para eliminar una película
delete_movie() {
    MOVIE_ID="$1"
    echo "----------------------------------------------------"
    echo "Acción: Eliminando película con ID $MOVIE_ID"
    URL="http://localhost:8080/api/movies/$MOVIE_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Película eliminada exitosamente"
        echo "----------------------------------------------------" >&2
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "\n===== INICIANDO PRUEBAS PARA POST /api/movies ===="

# Caso 1: JSON Inválido
# run_post_test "JSON Inválido" "http://localhost:8080/api/movies" '{"title": "Test Movie",,}' 400

# Caso 2: Campo requerido faltante (title)
run_post_test "Campo requerido faltante (title)" "http://localhost:8080/api/movies" '{"synopsis": "Post Test Synopsis", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 120}' 400

# Caso 3: Falla de validación (duration_minutes=0)
run_post_test "Falla de validación (duration_minutes=0)" "http://localhost:8080/api/movies" '{"title": "Post Test Movie", "synopsis": "Test Synopsis", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 0}' 400

# Caso 4: Crear una película para probar duplicados
MOVIE_DATA='{"title": "Duplicate Test Movie", "synopsis": "Synopsis", "released_at": "2024-01-01T00:00:00Z", "duration_minutes": 100}'
MOVIE_ID=$(create_movie "Crear película para prueba de duplicados" "$MOVIE_DATA")

# Caso 5: Película duplicada
run_post_test "Película duplicada" "http://localhost:8080/api/movies" "$MOVIE_DATA" 409

# Limpieza: Eliminar la película creada
if [ ! -z "$MOVIE_ID" ]; then
    delete_movie "$MOVIE_ID"
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para POST /api/movies pasaron exitosamente. ✅ \n"
fi
