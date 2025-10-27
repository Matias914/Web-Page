#!/bin/bash

# Salir si un comando falla
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
chmod +x "$SCRIPT_DIR/test_get_celebrities.sh"
chmod +x "$SCRIPT_DIR/test_post_celebrity.sh"
chmod +x "$SCRIPT_DIR/test_get_celebrity.sh"
chmod +x "$SCRIPT_DIR/test_put_celebrity.sh"
chmod +x "$SCRIPT_DIR/test_delete_celebrity.sh"
chmod +x "$SCRIPT_DIR/test_get_celebrity_roles.sh"
chmod +x "$SCRIPT_DIR/test_get_users.sh"
chmod +x "$SCRIPT_DIR/test_post_user.sh"
chmod +x "$SCRIPT_DIR/test_get_user.sh"
chmod +x "$SCRIPT_DIR/test_put_user.sh"
chmod +x "$SCRIPT_DIR/test_delete_user.sh"
chmod +x "$SCRIPT_DIR/test_post_rating.sh"
chmod +x "$SCRIPT_DIR/test_get_user_ratings.sh"
chmod +x "$SCRIPT_DIR/test_get_rating.sh"
chmod +x "$SCRIPT_DIR/test_put_rating.sh"
chmod +x "$SCRIPT_DIR/test_delete_rating.sh"
chmod +x "$SCRIPT_DIR/test_post_review.sh"
chmod +x "$SCRIPT_DIR/test_delete_review.sh"
chmod +x "$SCRIPT_DIR/test_get_movie_ratings.sh"
chmod +x "$SCRIPT_DIR/test_get_movie_reviews.sh"
chmod +x "$SCRIPT_DIR/test_get_movie_roles.sh"
chmod +x "$SCRIPT_DIR/test_post_movie_role.sh"
chmod +x "$SCRIPT_DIR/test_get_movie_role.sh"
chmod +x "$SCRIPT_DIR/test_put_movie_role.sh"
chmod +x "$SCRIPT_DIR/test_delete_movie_role.sh"
chmod +x "$SCRIPT_DIR/test_post_movie_category.sh"
chmod +x "$SCRIPT_DIR/test_get_movie_categories.sh"
chmod +x "$SCRIPT_DIR/test_get_movie_category.sh"
chmod +x "$SCRIPT_DIR/test_delete_movie_category.sh"
chmod +x "$SCRIPT_DIR/test_get_genre_movies.sh"

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

run_script "$SCRIPT_DIR/test_get_celebrities.sh"
run_script "$SCRIPT_DIR/test_post_celebrity.sh"
run_script "$SCRIPT_DIR/test_get_celebrity.sh"
run_script "$SCRIPT_DIR/test_put_celebrity.sh"
run_script "$SCRIPT_DIR/test_delete_celebrity.sh"
run_script "$SCRIPT_DIR/test_get_celebrity_roles.sh"

run_script "$SCRIPT_DIR/test_get_users.sh"
run_script "$SCRIPT_DIR/test_post_user.sh"
run_script "$SCRIPT_DIR/test_get_user.sh"
run_script "$SCRIPT_DIR/test_put_user.sh"
run_script "$SCRIPT_DIR/test_delete_user.sh"

run_script "$SCRIPT_DIR/test_post_rating.sh"
run_script "$SCRIPT_DIR/test_get_user_ratings.sh"
run_script "$SCRIPT_DIR/test_get_rating.sh"
run_script "$SCRIPT_DIR/test_put_rating.sh"
run_script "$SCRIPT_DIR/test_delete_rating.sh"

run_script "$SCRIPT_DIR/test_post_review.sh"
run_script "$SCRIPT_DIR/test_get_user_reviews.sh"
run_script "$SCRIPT_DIR/test_get_review.sh"
run_script "$SCRIPT_DIR/test_put_review.sh"
run_script "$SCRIPT_DIR/test_delete_review.sh"

run_script "$SCRIPT_DIR/test_get_movie_ratings.sh"
run_script "$SCRIPT_DIR/test_get_movie_reviews.sh"
run_script "$SCRIPT_DIR/test_get_movie_roles.sh"
run_script "$SCRIPT_DIR/test_post_movie_role.sh"
run_script "$SCRIPT_DIR/test_get_movie_role.sh"
run_script "$SCRIPT_DIR/test_put_movie_role.sh"
run_script "$SCRIPT_DIR/test_delete_movie_role.sh"

run_script "$SCRIPT_DIR/test_post_movie_category.sh"
run_script "$SCRIPT_DIR/test_get_movie_categories.sh"
run_script "$SCRIPT_DIR/test_get_movie_category.sh"
run_script "$SCRIPT_DIR/test_delete_movie_category.sh"

run_script "$SCRIPT_DIR/test_get_genre_movies.sh"