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

create_genre() {
    TEST_NAME="$1"
    URL="http://localhost:8080/api/genres"
    DATA_TEMPLATE="$2"

    # Generate a unique name
    UNIQUE_NAME="$(echo "$DATA_TEMPLATE" | jq -r '.name')-$(date +%s)"
    DATA=$(echo "$DATA_TEMPLATE" | jq --arg name "$UNIQUE_NAME" '.name = $name')

    # 1. Intentar POST (Lógica de curl, tail, sed es correcta)
    BODY_AND_STATUS=$(curl -s -w "%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')

    if [ "$HTTP_STATUS" -eq 201 ]; then
        # Manejo de Éxito (201)
        GENRE_ID=$(echo "$BODY" | jq -r '.id')
        if [ -z "$GENRE_ID" ] || [ "$GENRE_ID" == "null" ]; then
             echo "Error: Creación OK (201), pero ID no encontrado. Body: $BODY" >&2
             exit 1
        fi
        echo -e "\nGénero creado con ID: $GENRE_ID (Nombre: $UNIQUE_NAME)" >&2
        echo "$GENRE_ID"
        FAILED=1

    elif [ "$HTTP_STATUS" -eq 409 ]; then
        echo "Error: Género con nombre '$UNIQUE_NAME' ya existe (Status 409). El test no puede continuar con datos duplicados." >&2
        FAILED=1

    else
        echo "Error al crear el género. Status: $HTTP_STATUS, Body: $BODY" >&2
        FAILED=1
    fi
}

# Función para eliminar un género
delete_genre() {
    GENRE_ID="$1"
    echo "----------------------------------------------------" >&2
    echo "Acción: Eliminando género con ID $GENRE_ID" >&2

    if [ -z "$GENRE_ID" ] || [ "$GENRE_ID" == "null" ]; then
        echo "Advertencia: No se puede eliminar el género. ID de género vacío o nulo." >&2
        return 1
    fi

    URL="http://localhost:8080/api/genres/$GENRE_ID"

    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$URL")

    if [ "$HTTP_STATUS" -eq 204 ]; then
        echo "Género eliminado exitosamente" >&2
        echo "----------------------------------------------------" >&2
    else
        echo "FALLÓ la eliminación - Se esperaba 204, se obtuvo $HTTP_STATUS" >&2
        echo "----------------------------------------------------" >&2
        # exit 1 # No salimos para permitir la limpieza de otros géneros
    fi
}

echo -e "\n===== INICIANDO PRUEBAS PARA GET /api/genres/{id} ====="

# Caso 1: ID válido
GENRE_ID=$(create_genre "Crear género para obtener" '{"name": "Get Test Genre"}')
run_get_test "ID válido" "http://localhost:8080/api/genres/$GENRE_ID" 200

# Caso 2: ID no es un número
run_get_test "ID no es un número" "http://localhost:8080/api/genres/abc" 500

# Caso 3: Género no encontrado (ID muy grande)
run_get_test "Género no encontrado" "http://localhost:8080/api/genres/999999" 404

delete_genre "$GENRE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para GET /api/genres/{id} pasaron exitosamente. ✅ \n"
fi