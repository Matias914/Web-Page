document.addEventListener('DOMContentLoaded', () => {

    const catalogGrid = document.getElementById('catalog-grid');

    // Hace fetch de la API REST para buscar una cierta cantidad de películas junto a sus datos
    const fetchMovies = async () => {
        try {
            const response = await fetch('/api/movies?page=1&rows=50');
            if (!response.ok) {
                console.error(`Error HTTP! status: ${response.status}`);
                return [];
            }
            return await response.json();
        } catch (error) {
            console.error('Error de red al buscar películas:', error);
            return [];
        }
    };

    // Recibe una lista de películas en formato JSON y las convierte al formato HTML.
    const displayMovies = (movies) => {
        catalogGrid.innerHTML = '';
        if (!movies || movies.length === 0) {
            catalogGrid.innerHTML = '<p>No hay películas en el catálogo.</p>';
            return;
        }
        const movieCards = movies.map(({ title, synopsis, poster_url }) => {
            return `
                <div class="movie-card">
                    <img src="${poster_url || 'https://via.placeholder.com/220x330.png?text=No+Poster'}" alt="Póster de ${title}">
                    <div class="movie-card-content">
                        <h3>${title}</h3>
                        <p>${synopsis}</p>
                    </div>
                </div>
            `;
        });
        catalogGrid.innerHTML = movieCards.join('');
    };

    // Inicializa el catálogo con los datos de las peliculas. Mostrando un mensaje de espera.
    const initializeCatalog = async () => {
        if (catalogGrid) {
            catalogGrid.innerHTML = '<p>Cargando películas...</p>';
            const movies = await fetchMovies();
            displayMovies(movies);
        }
    };

    initializeCatalog().catch(error => {
        console.error('Error fatal al inicializar el catálogo:', error);
        const catalogGrid = document.getElementById('catalog-grid');
        if (catalogGrid) {
            catalogGrid.innerHTML = '<p>Error crítico al cargar la página. Por favor, refresca.</p>';
        }
    });
});