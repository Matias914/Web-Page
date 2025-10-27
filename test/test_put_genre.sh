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

create_genre() {
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="http://localhost:8080/api/genres"
    DATA="$2"

    BODY_AND_STATUS=$(curl -s -w "\n%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '\n')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
       GENRE_ID=$(echo "$BODY" | jq -r '.id')

       if [ -z "$GENRE_ID" ] || [ "$GENRE_ID" == "null" ]; then
            echo "Error: Creación OK (201), pero ID no encontrado en la respuesta JSON. Body: $BODY" >&2
            FAILED=1
       fi

       echo "Género creado con ID: $GENRE_ID" >&2
       echo "$GENRE_ID"
       return 0
    else
       echo "Error al crear el género. Status: $HTTP_STATUS, Body: $BODY" >&2
       FAILED=1
    fi
}

# Función para eliminar un género
delete_genre() {
    GENRE_ID="$1"
    echo "----------------------------------------------------"
    echo "Acción: Eliminando género con ID $GENRE_ID"
    URL="http://localhost:8080/api/genres/$GENRE_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Género eliminado exitosamente"
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA PUT /api/genres/{id} ====="

# Crear géneros para las pruebas
GENRE_ID_1=$(create_genre "Crear género 1 para actualizar" '{"name": "Put Test Genre 1"}')
GENRE_ID_2=$(create_genre "Crear género 2 para conflicto" '{"name": "Put Test Genre 2"}')

# Caso 1: ID no es un número
run_put_test "ID no es un número" "http://localhost:8080/api/genres/abc" '{}' 400

# Caso 2: Género no encontrado
run_put_test "Género no encontrado" "http://localhost:8080/api/genres/999999" '{"name": "some name"}' 404

# Caso 3: JSON Inválido
run_put_test "JSON Inválido" "http://localhost:8080/api/genres/$GENRE_ID_1" '{"name": ""}' 400

# Caso 4: Conflicto de duplicados
CONFLICT_DATA='{"name": "put test genre 1"}'
run_put_test "Conflicto de duplicados" "http://localhost:8080/api/genres/$GENRE_ID_2" "$CONFLICT_DATA" 409

# Limpieza
delete_genre "$GENRE_ID_1"
delete_genre "$GENRE_ID_2"

if [[ "$FAILED" -eq 0 ]]; then
    echo "----------------------------------------------------" >&2
    echo -e "Todas las pruebas para PUT /api/genres/{id} pasaron exitosamente. ✅ \n"
fi