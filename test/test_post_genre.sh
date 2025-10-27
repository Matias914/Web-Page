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

create_genre() {
    echo "----------------------------------------------------" >&2
    echo "Acción: $1" >&2
    URL="http://localhost:8080/api/genres"
    DATA="$2"

    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
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
        echo "----------------------------------------------------" >&2
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS"
    fi
}

echo -e "
===== INICIANDO PRUEBAS PARA POST /api/genres ===="

# Caso 1: JSON Inválido
run_post_test "JSON Inválido" "http://localhost:8080/api/genres" '{"name": "Test Genre",,}' 400

# Caso 2: Campo requerido faltante (name)
run_post_test "Campo requerido faltante (name)" "http://localhost:8080/api/genres" '{}' 400

# Caso 3: Crear un género para probar duplicados
GENRE_DATA='{"name": "Duplicate Test Genre"}'
GENRE_ID=$(create_genre "Crear género para prueba de duplicados" "$GENRE_DATA")

# Caso 4: Género duplicado
run_post_test "Género duplicado" "http://localhost:8080/api/genres" "$GENRE_DATA" 409

# Limpieza: Eliminar el género creado
if [ ! -z "$GENRE_ID" ]; then
    delete_genre "$GENRE_ID"
fi

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para POST /api/genres pasaron exitosamente. ✅
"
fi