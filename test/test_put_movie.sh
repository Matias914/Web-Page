#!/bin/bash

# Verificar si jq está instalado
if ! command -v jq &> /dev/null
then
    echo "Error: jq no está instalado. Por favor, instálalo para ejecutar las pruebas." >&2
    exit 1
fi

FAILED=0

run_put_test() {

    URL="$2"
    DATA="$3"
    EXPECTED_STATUS="$4"
    
    COMMAND="curl -s -o /dev/null -w '%{http_code}' -X PUT -H 'Content-Type: application/json' -d '$DATA' '$URL'"
    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X PUT -H "Content-Type: application/json" -d "$DATA" "$URL")
    

    if [[ ! "$HTTP_STATUS" -eq "$EXPECTED_STATUS" ]]; then
        echo -e "\n----------------------------------------------------"
        echo "Test: $1"
        echo "Comando: $COMMAND"
        echo  "FALLÓ - Se esperaba $EXPECTED_STATUS ❌"
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
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "\n===== INICIANDO PRUEBAS PARA PUT /api/movies/{id} ====="

# Crear películas para las pruebas
MOVIE_ID_1=$(create_movie "Crear película 1 para actualizar" '{"title": "Put Test Movie 1", "synopsis": "Synopsis", "released_at": "2025-01-01T00:00:00Z", "duration_minutes": 110}')
MOVIE_ID_2=$(create_movie "Crear película 2 para conflicto" '{"title": "Put Test Movie 2", "synopsis": "Synopsis", "released_at": "2025-01-02T00:00:00Z", "duration_minutes": 120}')

# Caso 1: ID no es un número
run_put_test "ID no es un número" "http://localhost:8080/api/movies/abc" '{}' 400

# Caso 2: Película no encontrada
run_put_test "Película no encontrada" "http://localhost:8080/api/movies/999999" '{"id": 999999, "title": "Not Found Test", "synopsis": "Valid synopsis", "released_at": "2025-01-01T00:00:00Z", "duration_minutes": 100}' 404

# Caso 3: JSON Inválido
run_put_test "JSON Inválido" "http://localhost:8080/api/movies/$MOVIE_ID_1" '{"title": "Invalid JSON",}' 400

# Case 4: Falla de validación (duration_minutes=0)
DATA_CASE_4='{"id": '"$MOVIE_ID_1"', "title": "Valid Title", "synopsis": "Valid Synopsis", "released_at": "2025-01-01T00:00:00Z", "duration_minutes": 0}'
run_put_test "Falla de validación" "http://localhost:8080/api/movies/$MOVIE_ID_1" "$DATA_CASE_4" 400

# Caso 5: Conflicto de duplicados
CONFLICT_DATA='{"id": '"$MOVIE_ID_1"', "title": "Put Test Movie 2", "synopsis": "Synopsis", "released_at": "2025-01-02T00:00:00Z", "duration_minutes": 120}'
run_put_test "Conflicto de duplicados" "http://localhost:8080/api/movies/$MOVIE_ID_1" "$CONFLICT_DATA" 409

# Limpieza
delete_movie "$MOVIE_ID_1"
delete_movie "$MOVIE_ID_2"

if [[ "$FAILED" -eq 0 ]]; then
    echo "----------------------------------------------------" >&2
    echo -e "Todas las pruebas para PUT /api/movies/{id} pasaron exitosamente. ✅ \n"
fi