# 🍿 Plataforma Web de Cine — Entrega TP4

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-API_Docs-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)

> 📢 **Nota de la Entrega:** Esta rama contiene el código fuente consolidado correspondiente al **Trabajo Práctico 4 (TP4)** de la UNICEN. El alcance de este hito completa la suite full-stack del proyecto, integrando la **Capa de Presentación**: Server-Side Rendering (SSR) mediante plantillas dinámicas de Go, procesamiento de assets estáticos independientes, manejo modularizado de UI con JavaScript Vanilla y vistas completas para el Catálogo de Películas y el Panel de Administración.

![Demostración de la Aplicación](docs/web-page.gif)

---

## 🏗️ Arquitectura de la Solución y Vistas de UI

El proyecto evoluciona hacia un monolito robusto bien estructurado, manteniendo el desacoplamiento backend en capas e integrando de forma nativa la resolución y despacho de componentes frontend a través de la capa de transporte.

```
.
├── cmd/
│   └── web/main.go             # Punto de entrada de la aplicación e inyección de dependencias
├── internal/                   # Código privado del dominio backend
│   ├── api/                    # Controladores REST, enrutamiento API y payloads de datos
│   ├── config/                 # Parser de configuraciones de entorno (.env)
│   ├── middleware/             # Componentes de interceptación HTTP (Loggers, Auth, CORS)
│   ├── service/                # Capa de Lógica de Negocio y validaciones de dominio
│   ├── storage/                # Capa de Persistencia (PostgreSQL + SQLC + Atlas Migrations)
│   └── transport/              # Capa HTTP: Gestión del ciclo del servidor y Motor de Renderizado (SSR)
├── web/                        # 🎨 Capa de Presentación (Frontend)
│   ├── static/                 # Recursos Estáticos Dedicados
│   │   ├── js/                 # Control de comportamiento asíncrono y fetch de datos
│   │   │   ├── catalog.js      # Lógica interactiva del catálogo de películas
│   │   │   └── management.js   # Orquestación del panel de administración (ABM/CRUD)
│   │   └── styles/             # Arquitectura CSS de componentes
│   │       ├── catalog.css     # Estilos de grillas de películas y filtros
│   │       ├── error.css       # Layouts adaptados para códigos de error HTTP
│   │       ├── index.css       # Estilos globales y variables de UI
│   │       └── management.css  # Diseño de tablas y formularios de gestión
│   └── templates/              # Plantillas dinámicas de Go (HTML estructurado por servidor)
│       ├── 404.html            # Vista de Recurso No Encontrado
│       ├── 500.html            # Vista de Error Crítico del Servidor
│       ├── catalog.html        # Layout del catálogo y cartelera pública
│       ├── index.html          # Portal de bienvenida principal
│       └── management.html     # Consola de administración y ABM de datos cinemáticos
├── test/                       # Suite de Pruebas de Integración Automatizadas (Caja Negra)
├── Makefile                    # Automatización global de compilación, generación y despliegue
└── docker-compose.yml          # Orquestación multi-contenedor (App Node + PostgreSQL DB)
```

---

## 🛠️ Stack Tecnológico Completo

* **Backend Core:** [Go (Golang)](https://go.dev/) (v1.21+) utilizando ruteo nativo y renderizado HTML concurrente.
* **Frontend Engine:** Server-Side Rendering (SSR) con `html/template` de Go + JavaScript Vanilla moderno (Fetch API de consumo interno).
* **Persistencia Relacional:** [PostgreSQL](https://www.postgresql.org/) con compilación type-safe vía [SQLC](https://sqlc.dev/).
* **Control de Esquemas:** Migraciones declarativas versionadas administradas por [Atlas](https://atlasgo.io/).
* **Developer Experience:** Recarga en caliente en tiempo real con [Air](https://github.com/air-verse/air).
* **Diseño e Interfaz de API:** Documentación interactiva y contratos mediante [Swagger](https://swagger.io/).

---

## 🚀 Despliegue Local del Entorno

### Requisitos Previos

Asegúrate de contar con las siguientes herramientas instaladas de forma local:
* [Docker & Docker Compose](https://docs.docker.com/engine/install/)
* [Go Toolchain](https://go.dev/dl/) (v1.21 o superior)
* [GNU Make](https://www.gnu.org/software/make/)

Inicializa los utilitarios globales del ecosistema de desarrollo ejecutando:
```bash
go install [github.com/sqlc-dev/sqlc/cmd/sqlc@latest](https://github.com/sqlc-dev/sqlc/cmd/sqlc@latest)
go install [github.com/air-verse/air@latest](https://github.com/air-verse/air@latest)
go install ariga.io/atlas/cmd/atlas@latest
go install [github.com/swaggo/swag/cmd/swag@latest](https://github.com/swaggo/swag/cmd/swag@latest)
```
> *Nota: Recuerda verificar la exportación de tus binarios ejecutando:* `export PATH=$PATH:$HOME/go/bin`

### Lanzamiento en un Solo Comando

1. **Navegar e Inicializar Variables:**
   ```bash
   cd Web-Page
   cp .env.example .env
   ```

2. **Compilar y Levantar el Ecosistema Full-Stack:**
   El archivo `Makefile` se encarga de empaquetar el código, levantar la base de datos PostgreSQL, aplicar las migraciones y renderizar las vistas web dinámicas automáticamente:
   ```bash
   make prod
   ```

3. **¡Acceso Listo!**
   * **Plataforma Web (Catálogo/Admin):** Accede mediante tu navegador a `http://localhost:8080`
   * **Consola Interactiva de API:** Disponible directamente en `http://localhost:8080/swagger/`

4. **Detener Servicios:**
   ```bash
   make prod-down
   ```

---

## 🧪 Pruebas Automatizadas y Sandbox

### Pruebas de Integración
El proyecto dispone de un flujo de validación automatizado que monta una base de datos temporal e impacta los endpoints del sistema usando scripts estructurados en Bash, `curl` y validadores lógicos basados en `jq`:
```bash
make test
```
*Este proceso está diseñado para ejecutarse en contenedores aislados y paralelos, garantizando que el entorno local de desarrollo no sufra alteraciones.*

---

## 🔄 Dinámica de Trabajo y Flujos de Desarrollo

### Recarga Automática en Caliente
Al programar bajo el modo de desarrollo (`make dev`), el demonio **Air** monitoreará de manera activa tanto tus archivos del backend `.go` y estructuras de datos `.sql`, como los componentes HTML y archivos dinámicos JS de la carpeta `web/`. Cualquier cambio se verá reflejado en el navegador en cuestión de segundos.

### Ciclo de Vida de Base de Datos
La persistencia de la base de datos se maneja en base a código usando el flujo declarativo provisto por **Atlas**:
1. **Modificación:** Edita el archivo maestro de esquemas: `internal/storage/postgres/schema/schema.sql`.
2. **Cálculo de Diferencial:** Corre `make migrate-diff NAME=nombre_descriptivo` para estructurar la nueva migración.
3. **Aplicación Física:** Ejecuta `make migrate-up` para consolidar los cambios estructurales sobre el contenedor activo.

---

## 📊 Arquitectura de la Base de Datos

La estructura relacional mapea eficientemente entidades de películas, géneros cinematográficos, usuarios, valoraciones y roles de actores dentro de la base de datos:

![Modelo de Arquitectura Relacional](docs/db.png)

---

## 📋 Comandos del Panel de Control (Makefile)

El `Makefile` centraliza el control operativo del proyecto full-stack. Ejecuta `make help` para auditar los comandos:

```bash
$ make help
Uso: make [comando]

Comandos Principales:
  dev           - Inicializa DB, ejecuta migraciones y levanta el servidor web con Hot-Reload (Air).
  prod          - Construye binarios y despliega la solución completa (Web + Base de Datos) en Docker.
  prod-down     - Apaga los contenedores de producción y limpia recursos huérfanos.
  server        - Corre el servidor web en caliente de manera aislada.

Gestión de Almacenamiento (Docker):
  db-up         - Levanta de forma aislada el contenedor de PostgreSQL.
  db-down       - Apaga el contenedor activo de la base de datos.
  db-nuke       - Elimina volúmenes persistentes y purga por completo los datos de almacenamiento.

Sistemas de Migraciones (Atlas):
  migrate-diff  - Compara e implementa un nuevo paso de migración incremental (Requiere NAME).
  migrate-up    - Sincroniza y ejecuta todas las migraciones históricas pendientes en la base de datos.

Compilación y Desarrollo:
  sqlc-gen      - Compila los queries SQL en código Go puro, fuertemente tipado y nativo.
  build         - Compila la aplicación web completa en un binario de producción.
  test          - Ejecuta la suite completa de pruebas de integración en entornos aislados de Docker.
  swagger       - Actualiza los esquemas OpenAPI y regenera el dashboard interactivo.
```

---

## 👥 Desarrolladores

* **Ortiz Matias** - *Estudiante de Ingeniería de Sistemas (UNICEN)*
* **Leon Nicolas** - *Estudiante de Ingeniería de Sistemas (UNICEN)*

**Universidad Nacional del Centro de la Provincia de Buenos Aires (UNICEN)**
Facultad de Ciencias Exactas — Tandil, Buenos Aires, Argentina.
