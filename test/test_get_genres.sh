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
        echo "FALLÓ - Se esperaba $EXPECTED_STATUS"
        echo "Comando: $COMMAND"
        echo "Resultado: $HTTP_STATUS"
        echo "----------------------------------------------------"
        FAILED=1
    fi
}

echo -e "\n===== INICIANDO PRUEBAS PARA GET /api/genres ====="

# Caso 1: Paginación correcta
run_get_test "Paginación correcta (page=1, rows=5)" "http://localhost:8080/api/genres?page=1&rows=5" 200

# Caso 2: Número de página inválido (no es un número)
run_get_test "Número de página inválido (page=abc)" "http://localhost:8080/api/genres?page=abc&rows=5" 400

# Caso 3: Número de página inválido (menor a 1)
run_get_test "Número de página inválido (page=0)" "http://localhost:8080/api/genres?page=0&rows=5" 400

# Caso 4: Filas por página inválidas (no es un número)
run_get_test "Filas por página inválidas (rows=abc)" "http://localhost:8080/api/genres?page=1&rows=abc" 400

# Caso 5: Filas por página inválidas (menor a 1)
run_get_test "Filas por página inválidas (rows=0)" "http://localhost:8080/api/genres?page=1&rows=0" 400

if [[ "$FAILED" -eq 0 ]]; then
    echo "----------------------------------------------------" >&2
    echo -e "Todas las pruebas para GET /api/genres pasaron exitosamente. ✅ \n"
fi