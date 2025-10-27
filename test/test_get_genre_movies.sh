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

create_genre() {
    URL="http://localhost:8080/api/genres"
    DATA="$2"
    BODY_AND_STATUS=$(curl -s -w "
%{http_code}" -X POST -H "Content-Type: application/json" -d "$DATA" "$URL")
    HTTP_STATUS=$(echo "$BODY_AND_STATUS" | tail -n1 | tr -d '
')
    BODY=$(echo "$BODY_AND_STATUS" | sed '$d')
    if [ "$HTTP_STATUS" -eq 201 ]; then
       GENRE_ID=$(echo "$BODY" | jq -r '.id')
       echo "$GENRE_ID"
    fi
}

delete_genre() {
    GENRE_ID="$1"
    URL="http://localhost:8080/api/genres/$GENRE_ID"
    curl -s -o /dev/null -X DELETE "$URL"
}

echo -e "
===== INICIANDO PRUEBAS PARA GET /api/genres/{id}/categories ====="

# Setup: Crear género para las pruebas
GENRE_ID=$(create_genre "" '{"name": "Get Genre Movies"}')

if [ -z "$GENRE_ID" ]; then
    echo "Error: No se pudo crear el género necesario para las pruebas."
    exit 1
fi

# Caso 1: Obtención exitosa (incluso si está vacía)
run_get_test "Obtención exitosa" "http://localhost:8080/api/genres/$GENRE_ID/categories?page=1&rows=5" 200

# Caso 2: ID de género no es un número
run_get_test "ID de género no es un número" "http://localhost:8080/api/genres/abc/categories?page=1&rows=5" 500

# Caso 3: Género no encontrado
run_get_test "Género no encontrado" "http://localhost:8080/api/genres/999999/categories?page=1&rows=5" 404

# Caso 4: Paginación inválida
run_get_test "Paginación inválida" "http://localhost:8080/api/genres/$GENRE_ID/categories?page=abc&rows=5" 500

# Limpieza
delete_genre "$GENRE_ID"

if [[ "$FAILED" -eq 0 ]]; then
    echo -e "Todas las pruebas para GET /api/genres/{id}/categories pasaron exitosamente. ✅
"
fi
