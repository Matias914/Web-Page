# 🍿 Plataforma de Gestión y Catálogo de Películas — Entrega TP3

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-API_Docs-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)

> 📢 **Nota de la Entrega:** Esta rama contiene el código fuente correspondiente al **Trabajo Práctico 3 (TP3)** de la currícula de la **UNICEN**. El alcance de este hito se centra exclusivamente en el motor del backend: implementación de la arquitectura en capas, persistencia de datos, tracking automatizado de esquemas relacionales con Atlas, generación de consultas type-safe mediante SQLC y la definición de contratos de la API REST con Swagger.

---

## 🏗️ Arquitectura y Estructura del Proyecto

El proyecto sigue la convención estándar de layouts en Go, implementando un patrón desacoplado en capas (`Transporte ➔ Servicio ➔ Almacenamiento`) para garantizar el aislamiento de la lógica de negocio y una inyección de dependencias limpia.

```
.
├── cmd/
│   └── web/main.go             # Punto de entrada de la aplicación y cableado de dependencias
├── internal/                   # Código privado del dominio y sus capas
│   ├── api/                    # Controladores REST y manejadores de payloads HTTP
│   ├── config/                 # Estructuras de validación de variables de entorno
│   ├── middleware/             # Hooks de middleware (logging, recuperación, seguridad)
│   ├── service/                # Capa de Lógica de Negocio (reglas de dominio y validaciones)
│   ├── transport/              # Ciclo de vida del servidor HTTP y ruteo central
│   └── storage/                # Capa de Acceso a Datos (DAL)
│       └── postgres/
│           ├── schema/         # Esquema declarativo SQL (Única fuente de verdad)
│           ├── queries/        # Plantillas de consultas SQL optimizadas para sqlc
│           ├── migrations/     # Logs de migraciones versionadas planificadas por Atlas
│           └── sqlc/           # Código Go seguro y tipado, auto-generado desde las queries
├── web/                        # Componentes de Presentación
│   ├── static/styles/          # Hojas de estilo base para la interfaz
│   └── templates/              # Estructuras base de plantillas HTML de Go (Index, 404, 500)
├── test/                       # Pruebas de Integración Automatizadas
│   ├── Dockerfile              # Contenedor aislado para los tests (Bash + Curl + JQ)
│   └── run_all_tests.sh        # Motor de pruebas de caja negra sobre los endpoints REST
├── Makefile                    # Centro de comandos para la automatización de tareas
├── sqlc.yaml                   # Archivo de configuración para el compilador SQLC
└── docker-compose.yml          # Topología de orquestación (Servicio de App + DB PostgreSQL)
```

---

## 🛠️ Stack Tecnológico Core

* **Lenguaje y Runtime:** [Go (Golang)](https://go.dev/) (v1.21+) utilizando estructuras nativas compiladas.
* **Base de Datos:** Instancia contenedorizada de [PostgreSQL](https://www.postgresql.org/).
* **Generación de Código:** [SQLC](https://sqlc.dev/) para compilar sentencias SQL puras en funciones nativas de Go.
* **Framework de Migraciones:** [Atlas](https://atlasgo.io/) para la inspección declarativa del esquema y manejo de drifts.
* **Recarga en Caliente:** [Air](https://github.com/air-verse/air) para el reinicio automático del backend en desarrollo.
* **Documentación de API:** [Swagger/OpenAPI](https://github.com/swaggo/swag) proporcionando un entorno interactivo de pruebas.

---

## 🚀 Configuración del Entorno

### Requisitos Previos

Asegúrate de contar con las siguientes herramientas instaladas en tu sistema host:
* [Docker & Docker Compose](https://docs.docker.com/engine/install/)
* [Go Toolchain](https://go.dev/dl/) (v1.21+)
* [GNU Make](https://www.gnu.org/software/make/)

Instala las utilidades de desarrollo Go necesarias ejecutando:
```bash
# Generador de acceso a datos de SQLC
go install [github.com/sqlc-dev/sqlc/cmd/sqlc@latest](https://github.com/sqlc-dev/sqlc/cmd/sqlc@latest)

# Demonio de live-reloading
go install [github.com/air-verse/air@latest](https://github.com/air-verse/air@latest)

# Planificador de migraciones Atlas
go install ariga.io/atlas/cmd/atlas@latest

# Generador de documentación Swagger
go install [github.com/swaggo/swag/cmd/swag@latest](https://github.com/swaggo/swag/cmd/swag@latest)
```
> *Nota: Asegúrate de tener los binarios locales en tu PATH:* `export PATH=$PATH:$HOME/go/bin`

### Instalación y Lanzamiento del Sistema

1. **Clonar el proyecto y navegar al directorio:**
   ```bash
   git clone [https://github.com/Matias914/Web-Page.git](https://github.com/Matias914/Web-Page.git)
   cd Web-Page
   ```

2. **Establecer las variables de entorno:**
   ```bash
   cp .env.example .env
   ```

3. **Orquestar el encendido del sistema:**
   Compila los assets, descarga las imágenes correspondientes, ejecuta la estructura de la base de datos y levanta el servidor de producción usando el comando unificado:
   ```bash
   make prod
   ```
   El motor del backend estará disponible inmediatamente en `http://localhost:8080`.

4. **Apagar los contenedores de producción:**
   ```bash
   make prod-down
   ```

---

## 🧪 Testing y Evaluación Interactiva de la API

### Suite de Tests de Integración

Este hito incluye una matriz de pruebas automatizadas de caja negra. Al ejecutar los comandos de test, se despliega una base de datos de pruebas aislada, se inicializa un contenedor limpio de la aplicación y se impactan los endpoints REST mediante operaciones estructuradas de `curl`, cuyas respuestas son validadas con aserciones mediante el parseo de `jq`.
```bash
make test
```
*Los entornos de prueba corren completamente separados de la estructura de desarrollo, evitando cualquier corrupción de datos persistentes.*

### Sandbox de API Interactivo (Swagger)

Los endpoints, modelos de requests y códigos de respuesta de las entregas pueden ser evaluados interactivamente visitando el dashboard de Swagger integrado en la ruta: `http://localhost:8080/swagger/`.

![Swagger Live Automation Sandbox](docs/swagger.gif)

---

## 🔄 Flujos de Trabajo de Ingeniería (Alcance TP3)

### Recarga en Caliente (Hot Reload)
Al ejecutar el entorno de desarrollo mediante `make dev`, el sistema inicializa **Air** para escuchar los cambios en archivos `.go` o `.sql`. Cualquier guardado disparará una recompilación automática instantánea.

### Evolución Declarativa del Esquema
Los cambios en la base de datos se manejan de manera declarativa con **Atlas** para evitar inconsistencias entre entornos:
1. **Modificación del esquema:** Actualiza los cambios directamente en `internal/storage/postgres/schema/schema.sql`.
2. **Calcular diferencias (Diff):** Genera los archivos de migración versionados con `make migrate-diff NAME=nombre_del_cambio`.
3. **Aplicar cambios estructurales:** Ejecuta `make migrate-up` para sincronizarlos con el motor de almacenamiento del contenedor activo.

---

## 📋 Referencia de Comandos Disponibles (Makefile)

El archivo `Makefile` actúa como el punto de entrada absoluto para el control del entorno. Ejecuta `make help` para listar todas las acciones disponibles:

```bash
$ make help
Uso: make [comando]

Comandos Principales:
  dev           - Levanta la base de datos, corre las migraciones e inicializa el servidor local con hot-reload (Air).
  prod          - Compila los binarios y enciende la red completa de servicios vía Docker.
  prod-down     - Detiene y remueve los recursos y contenedores de producción.
  server        - Ejecuta el demonio de la aplicación de forma nativa.

Gestión de Base de Datos:
  db-up         - Inicializa el servicio contenedorizado de PostgreSQL.
  db-down       - Detiene de forma segura los contenedores de la base de datos.
  db-nuke       - Elimina las estructuras de datos, limpia volúmenes y purga streams activos.

Migraciones con Atlas:
  migrate-diff  - Calcula el delta del esquema basándose en el archivo declarativo (Requiere parámetro NAME).
  migrate-up    - Impacta los scripts históricos pendientes en la base de datos objetivo.

Generación y Desarrollo:
  sqlc-gen      - Ejecuta el parser de sqlc para escribir los bindings estructurales de Go desde el código SQL.
  build         - Compila los binarios optimizados para la arquitectura nativa.
  test          - Levanta el entorno y ejecuta la suite de pruebas de caja negra contenedorizadas.
  swagger       - Recompila los contratos de la interfaz OpenAPI y actualiza las vistas.
```

---

## 👥 Desarrolladores

* **Matias Ortiz** - *Estudiante de Ingeniería de Sistemas (UNICEN)*
* **Nicolas Leon** - *Estudiante de Ingeniería de Sistemas (UNICEN)*

**Universidad Nacional del Centro de la Provincia de Buenos Aires (UNICEN)** Facultad de Ciencias Exactas — Tandil, Buenos Aires, Argentina.
