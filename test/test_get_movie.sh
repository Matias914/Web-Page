#!/bin/bash

# Verificar si jq está instalado
if ! command -v jq &> /dev/null
then
    echo "Error: jq no está instalado. Por favor, instálalo para ejecutar las pruebas." >&2
    exit 1
fi

FAILED=0

# Función para hacer una solicitud GET y mostrar los resultados
run_get_test() {
    URL="$2"
    EXPECTED_STATUS="$3"
    COMMAND="curl -s -o /dev/null -w '%{http_code}' '$URL'"
    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

    if [[ ! "$HTTP_STATUS" -eq "$EXPECTED_STATUS" ]]; then
        echo "Test: $1" >&2
        echo "FALLÓ - Se esperaba $EXPECTED_STATUS ❌" >&2
        echo "Comando: $COMMAND" >&2
        echo "Resultado: $HTTP_STATUS" >&2
        echo "----------------------------------------------------" >&2
        FAILED=1
    fi
}

create_movie() {
    TEST_NAME="$1"
    URL="http://localhost:8080/api/movies"
    DATA_TEMPLATE="$2"

    # Generate a unique title
    UNIQUE_TITLE="$(echo "$DATA_TEMPLATE" | jq -r '.title')-$(date +%s)"
    DATA=$(echo "$DATA_TEMPLATE" | jq --arg title "$UNIQUE_TITLE" '.title = $title')

    # 1. Intentar POST (Lógica de curl, tail, sed es correcta)
    BODY_AND_STATUS=$(curl -s -w "\n%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '\n')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
        # Manejo de Éxito (201)
        MOVIE_ID=$(echo "$BODY" | jq -r '.id')
        if [ -z "$MOVIE_ID" ] || [ "$MOVIE_ID" == "null" ]; then
             echo "Error: Creación OK (201), pero ID no encontrado. Body: $BODY" >&2
             exit 1
        fi
        echo "Película creada con ID: $MOVIE_ID (Título: $UNIQUE_TITLE)" >&2
        echo "$MOVIE_ID"
        FAILED=1

    elif [ "$HTTP_STATUS" -eq 409 ]; then
        echo "Error: Película con título '$UNIQUE_TITLE' ya existe (Status 409). El test no puede continuar con datos duplicados." >&2
        FAILED=1

    else
        echo "Error al crear la película. Status: $HTTP_STATUS, Body: $BODY" >&2
        FAILED=1
    fi
}

# Función para eliminar una película
delete_movie() {
    MOVIE_ID="$1"
    echo "----------------------------------------------------" >&2
    echo "Acción: Eliminando película con ID $MOVIE_ID" >&2

    if [ -z "$MOVIE_ID" ] || [ "$MOVIE_ID" == "null" ]; then
        echo "Advertencia: No se puede eliminar la película. ID de película vacío o nulo." >&2
        return 1
    fi

    URL="http://localhost:8080/api/movies/$MOVIE_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Película eliminada exitosamente" >&2
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS" >&2
        # exit 1 # No salimos para permitir la limpieza de otras películas
    fi
}

echo -e "\n===== INICIANDO PRUEBAS PARA GET /api/movies/{id} ====="

# Caso 1: ID válido
MOVIE_ID=$(create_movie "Crear película para borrar" '{"title": "Get Test Movie", "synopsis": "Synopsis", "released_at": "2023-01-01T00:00:00Z", "duration_minutes": 130}')
run_get_test "ID válido" "http://localhost:8080/api/movies/$MOVIE_ID" 200

# Caso 2: ID no es un número
run_get_test "ID no es un número" "http://localhost:8080/api/movies/abc" 400

# Caso 3: Película no encontrada (ID muy grande)
run_get_test "Película no encontrada" "http://localhost:8080/api/movies/999999" 404

delete_movie "$MOVIE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo "----------------------------------------------------" >&2
    echo -e "\nTodas las pruebas para GET /api/movies/{id} pasaron exitosamente. ✅ \n"
fi