document.addEventListener('DOMContentLoaded', () => {
    // ESTADO DE LA PÁGINA:
    // Entidad en pantalla, por defecto películas
    let currentEntity = 'movies';
    // Acción en pantalla, por ejemplo, creando o actualizando un objeto
    let editState = { isEditing: false, itemId: null };
    // Temporizador de la notificación
    let notificationTimer;

    // Metadatos de las Entidades manejadas en la página (determinados por currentEntity)
    const entityConfigs = {
        movies: {
            plural: 'Películas',
            singular: 'Película',
            apiEndpoint: '/api/movies',
            displayField: 'title',
            fields: [
                { name: 'title', placeholder: 'Título', type: 'text', required: true },
                { name: 'synopsis', placeholder: 'Sinopsis', type: 'textarea', required: true },
                { name: 'released_at', placeholder: 'Fecha de Lanzamiento', type: 'date', required: true },
                { name: 'duration_minutes', placeholder: 'Duración (minutos)', type: 'number', required: true },
                { name: 'poster_url', placeholder: 'URL del Póster', type: 'text', required: false },
            ]
        },
        genres: {
            plural: 'Géneros',
            singular: 'Género',
            apiEndpoint: '/api/genres',
            displayField: 'name',
            fields: [
                { name: 'name', placeholder: 'Nombre del Género', type: 'text', required: true },
            ]
        },
        celebrities: {
            plural: 'Celebridades',
            singular: 'Celebridad',
            apiEndpoint: '/api/celebrities',
            displayField: 'name',
            fields: [
                { name: 'name', placeholder: 'Nombre de la Celebridad', type: 'text', required: true },
                { name: 'birth_date', placeholder: 'Fecha de Nacimiento', type: 'date', required: true },
            ]
        }
    };

    // ESTRUCTURAS DINÁIMCAS DE LA PÁGINA
    // Div que permtie elegir entre diferentes entidades: películas, géneros, celebridades ...
    const entitySelector = document.getElementById('entity-selector');
    // Formulario dinámico que muestra los campos correspondientes a cada entidad
    const entityForm = document.getElementById('entity-form');
    // Título asociado al formulario dinámico: Crear Película, Crear Celebridad ...
    const formTitle = document.getElementById('form-title');
    // Campos asociados al formulario dinámico: título, sinopsis, nombre ...
    const formFields = document.getElementById('form-fields');
    // Boton dinámico asociado al formulario dinámico: añadir pelicula ...
    const formSubmitButton = document.getElementById('form-submit-button');
    // Título de la lista dinámica
    const listTitle = document.getElementById('list-title');
    // Lísta de entidades dinámica
    const entityList = document.getElementById('entity-list');
    // Div asociado a las notificaciones
    const notificationBanner = document.getElementById('notification-banner');

    // ------------------------------------------------------------------------------------------- //
    //                                          LOGICA                                             //
    // ------------------------------------------------------------------------------------------- //

    // Muestra una notificacion, de error o no, durante 5000 ms en pantalla
    const showNotification = (message, isError = false) => {
        if (!notificationBanner)
            return;
        clearTimeout(notificationTimer);
        notificationBanner.textContent = message;
        notificationBanner.className = 'notification';
        if (isError)
            notificationBanner.classList.add('error');
        notificationBanner.classList.add('show');
        notificationTimer = setTimeout(() => notificationBanner.classList.remove('show'), 5000);
    };

    // Se para en los metadatos de la entidad en pantalla y agrega un campo al formulario dinámico por cada
    // elemento del arreglo fields, determinando los datos de ID, placeholder y required. Se renderiza como
    // html
    const renderForm = (entityName) => {
        const config = entityConfigs[entityName];
        formFields.innerHTML = '';
        config.fields.forEach(field => {
            const element = document.createElement(field.type === 'textarea' ? 'textarea' : 'input');
            element.id = `field-${field.name}`;
            // Si es de tipo input, se debe definir un type.
            if (field.type !== 'textarea')
                element.type = field.type;
            element.placeholder = field.placeholder;
            element.required = field.required;
            formFields.appendChild(element);
        });
    };

    // Resetea el formulario al modo crear: elimina todos los campos, vuelve a poner los titulos de creacion,
    // el estado deja de ser de edicion
    const resetFormToCreateMode = () => {
        const config = entityConfigs[currentEntity];
        entityForm.reset();
        editState = { isEditing: false, itemId: null };
        formTitle.textContent = `Crear ${config.singular}`;
        formSubmitButton.textContent = `Añadir ${config.singular}`;
    };

    // Renderiza la card de gestión. Para ello necesita de la nueva entidad seleccionada y lo que hace es
    // definir la nueva entidad activa, agregar el atributo "active" al boton correspondiente, cambiar el
    // título de la lista según el nuevo nombre, resetear el formulario al modo crear (por defecto),
    // renderizar el formulario y mostrar los items.
    const renderUIForEntity = (entityName) => {
        currentEntity = entityName;
        const config = entityConfigs[entityName];
        document.querySelectorAll('.entity-btn').forEach(btn => {
            btn.classList.toggle('active', btn.dataset.entity === entityName);
        });
        listTitle.textContent = `Lista de ${config.plural}`;
        resetFormToCreateMode();
        renderForm(entityName);
        fetchAndDisplayItems(entityName).catch(error => {
            console.error('Error fatal al hacer un display de items:', error);
            const entityList = document.getElementById('li');
            if (entityList) {
                entityList.innerHTML = '<p>Error crítico al cargar los items. Por favor, refresca.</p>';
            }
        });
    };

    // Se hace un fetch de los datos a la API y se van agregando los elementos a la lista de
    // forma dinamica. Por cada elemento se define un Span (permite decorar el texto), el
    // campo de display (ej. si es un pelicula, el título) y el contenedor con los botones de
    // borrado y edicion. IMPORTANTE: aunque no se muestre, cada elemento guarda su ID dentro
    // del item.
    const fetchAndDisplayItems = async (entityName) => {
        const config = entityConfigs[entityName];
        entityList.innerHTML = '<li>Cargando...</li>';
        try {
            const response = await fetch(`${config.apiEndpoint}?page=1&rows=100`);
            if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
            const items = await response.json();
            entityList.innerHTML = '';
            if (items && items.length > 0) {
                items.forEach(item => {
                    const listItem = document.createElement('li');
                    const textSpan = document.createElement('span');
                    textSpan.textContent = item[config.displayField] || 'Sin título';
                    listItem.appendChild(textSpan);
                    listItem.dataset.id = item.id;
                    const buttonWrapper = document.createElement('div');
                    buttonWrapper.classList.add('action-buttons');
                    const editButton = document.createElement('button');
                    editButton.textContent = 'Editar';
                    editButton.classList.add('edit-btn');
                    buttonWrapper.appendChild(editButton);
                    const deleteButton = document.createElement('button');
                    deleteButton.textContent = 'Eliminar';
                    deleteButton.classList.add('delete-btn');
                    buttonWrapper.appendChild(deleteButton);
                    listItem.appendChild(buttonWrapper);
                    entityList.appendChild(listItem);
                });
            } else {
                entityList.innerHTML = `<li>No hay ${config.plural.toLowerCase()} para mostrar.</li>`;
            }
        } catch (error) {
            console.error(`Error al cargar ${config.plural}:`, error);
            showNotification(`Error al cargar ${config.plural.toLowerCase()}.`, true);
        }
    };

    // ------------------------------------------------------------------------------------------- //
    //                                         EVENTOS                                             //
    // ------------------------------------------------------------------------------------------- //

    // Escucha los clics. Si se da que el clic ocurrió sobre un botón de entidad, se hace un
    // fetch del nombre de la misma y se renderiza la UI (si la entidad cambió).
    entitySelector.addEventListener('click', (event) => {
        const target = event.target;
        if (target.classList.contains('entity-btn')) {
            const entityName = target.dataset.entity;
            if (entityName && entityName !== currentEntity) {
                renderUIForEntity(entityName);
            }
        }
    });

    // Escucha los botones de submit (el del formulario, en una creación o actualizacion). Crea un JSON
    // con todos los campos del formulario. Determina si se está queriendo hacer un PUT o POST dependiendo de
    // si el formulario era de actualizacion o creación. Para construir la URL se usan los metadatos de las
    // entidades. Si ocurre un error con el PUT o POST, se invoca el showNotification con el flag de error
    // en true. De lo contrario, se envia con el flag en falso y se hace un display de los elementos.
    entityForm.addEventListener('submit', async (event) => {
        event.preventDefault();
        const config = entityConfigs[currentEntity];
        const data = {};
        config.fields.forEach(field => {
            const input = document.getElementById(`field-${field.name}`);
            let value = input.value;
            if (input.type === 'number') value = parseInt(value, 10);
            else if (input.type === 'date') value = new Date(value).toISOString();
            data[field.name] = value || null;
        });
        const isEditing = editState.isEditing;
        const url = isEditing ? `${config.apiEndpoint}/${editState.itemId}` : config.apiEndpoint;
        const method = isEditing ? 'PUT' : 'POST';
        try {
            const response = await fetch(url, { method, headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(data) });
            if (!response.ok) {
                let errorDetails = '';
                const contentType = response.headers.get('content-type');
                if (contentType && contentType.includes('application/json')) {
                    errorDetails = JSON.stringify(await response.json());
                } else {
                    errorDetails = await response.text();
                }
                throw new Error(`HTTP error! status: ${response.status}, body: ${errorDetails}`);
            }
            const action = isEditing ? 'Actualización' : 'Creación';
            showNotification(`${action} de ${config.singular} exitosa.`);
            resetFormToCreateMode();
            await fetchAndDisplayItems(currentEntity);
        } catch (error) {
            const action = isEditing ? 'actualizar' : 'crear';
            console.error(`Error al ${action} ${config.singular}:`, error);
            showNotification(`Error al ${action} ${config.singular.toLowerCase()}.`, true);
        }
    });

    // Se escuchan los clics. Si el clic ocurrio dentro de un item de una lista, entonces se
    // verifica si ocurrio dentro del boton de edicion o borrado. Si es en el primero, se hace
    // un fetch del objeto usando el ID guardado en la lista y se rellena el formulario
    // dinamicamente con los datos obtenidos. Además, se hace un scroll al tope de la página.
    // Si es un DELETE, se pide confirmación del usuario. En caso afirmativo, se efectúa la
    // operación y se vuelve a hacer un fetch de los items
    entityList.addEventListener('click', async (event) => {
        const config = entityConfigs[currentEntity];
        const listItem = event.target.closest('li');
        if (!listItem) return;
        const itemId = listItem.dataset.id;
        if (event.target.classList.contains('edit-btn')) {
            try {
                const response = await fetch(`${config.apiEndpoint}/${itemId}`);
                if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
                const item = await response.json();
                config.fields.forEach(field => {
                    const input = document.getElementById(`field-${field.name}`);
                    if (!input) return;
                    const value = item[field.name];
                    if (field.type === 'date' && value) {
                        input.value = new Date(value).toISOString().split('T')[0];
                    } else {
                        input.value = value || '';
                    }
                });
                editState = { isEditing: true, itemId: itemId };
                formTitle.textContent = `Editando ${config.singular}`;
                formSubmitButton.textContent = 'Guardar Cambios';
                window.scrollTo({ top: 0, behavior: 'smooth' });
            } catch (error) {
                console.error('Error al cargar datos para editar:', error);
                showNotification('No se pudieron cargar los datos para editar.', true);
            }
        } else if (event.target.classList.contains('delete-btn')) {
            if (confirm(`¿Estás seguro de que quieres eliminar este elemento?`)) {
                try {
                    const response = await fetch(`${config.apiEndpoint}/${itemId}`, { method: 'DELETE' });
                    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
                    showNotification(`${config.singular} eliminado exitosamente.`);
                    await fetchAndDisplayItems(currentEntity);
                } catch (error) {
                    console.error(`Error al eliminar ${config.singular}:`, error);
                    showNotification(`Error al eliminar ${config.singular.toLowerCase()}.`, true);
                }
            }
        }
    });
    // La primera vez que se accede a la página se renderiza con la UI asociada a "movies"
    renderUIForEntity("movies");
});
