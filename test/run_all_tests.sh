#!/bin/bash

# Salir inmediatamente si un comando falla
set -e

# Obtener el directorio donde se encuentra este script
SCRIPT_DIR=$(dirname "$0")

# Dar permisos de ejecución a los scripts de prueba
chmod +x "$SCRIPT_DIR/test_get_movies.sh"
chmod +x "$SCRIPT_DIR/test_post_movie.sh"
chmod +x "$SCRIPT_DIR/test_get_movie.sh"
chmod +x "$SCRIPT_DIR/test_put_movie.sh"
chmod +x "$SCRIPT_DIR/test_delete_movie.sh"
chmod +x "$SCRIPT_DIR/test_get_genres.sh"
chmod +x "$SCRIPT_DIR/test_post_genre.sh"
chmod +x "$SCRIPT_DIR/test_get_genre.sh"
chmod +x "$SCRIPT_DIR/test_put_genre.sh"
chmod +x "$SCRIPT_DIR/test_delete_genre.sh"

# Función para ejecutar un script de prueba
run_script() {
    echo "======================================================"
    echo "EJECUTANDO: $1"
    echo "======================================================"
    "$1"
    echo ""
}

# Ejecutar todos los scripts de prueba
run_script "$SCRIPT_DIR/test_get_movies.sh"
run_script "$SCRIPT_DIR/test_post_movie.sh"
run_script "$SCRIPT_DIR/test_get_movie.sh"
run_script "$SCRIPT_DIR/test_put_movie.sh"
run_script "$SCRIPT_DIR/test_delete_movie.sh"

run_script "$SCRIPT_DIR/test_get_genres.sh"
run_script "$SCRIPT_DIR/test_post_genre.sh"
run_script "$SCRIPT_DIR/test_get_genre.sh"
run_script "$SCRIPT_DIR/test_put_genre.sh"
run_script "$SCRIPT_DIR/test_delete_genre.sh"