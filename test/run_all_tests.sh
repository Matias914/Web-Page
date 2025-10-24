#!/bin/bash

# Salir inmediatamente si un comando falla
set -e

# Dar permisos de ejecución a los scripts de prueba
chmod +x ./test_get_movies.sh
chmod +x ./test_post_movie.sh
chmod +x ./test_get_movie.sh
chmod +x ./test_put_movie.sh
chmod +x ./test_delete_movie.sh

# Función para ejecutar un script de prueba
run_script() {
    echo "======================================================"
    echo "EJECUTANDO: $1"
    echo "======================================================"
    ./$1
    echo ""
}

# Ejecutar todos los scripts de prueba
run_script "./test_get_movies.sh"
run_script "./test_post_movie.sh"
run_script "./test_get_movie.sh"
run_script "./test_put_movie.sh"
run_script "./test_delete_movie.sh"