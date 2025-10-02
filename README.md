# Web-Page 🍿

Este proyecto es una aplicación web desarrollada como parte de la currícula de la **Universidad Nacional del Centro de la Provincia de Buenos Aires (UNICEN)**. La aplicación se centra en el dominio de la cinematografía, permitiendo a los usuarios interactuar con una base de datos de películas, géneros, actores y más.

---
## Tecnologías Utilizadas

El proyecto está construido sobre un stack de tecnologías moderno, enfocado en la eficiencia, la automatización y las buenas prácticas de desarrollo.

* [Go (Golang)](https://go.dev/)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
* [GNU Make](https://www.gnu.org/software/make/)
* [SQL Compiler (sqlc)](https://sqlc.dev/)
* [Atlas](https://atlasgo.io/)
* [Air](https://github.com/cosmtrek/air)

---
## Estructura del Proyecto

El repositorio está organizado siguiendo convenciones estándar para facilitar la mantenibilidad y escalabilidad.

```
.
├── cmd/                        # Punto de entrada de la aplicación (main.go)
│   └── web/
│       └── main.go
├── internal/                   # Código privado del proyecto (no importable por otros)
│   ├── config/                 # Lógica de configuración del entorno
│   │   └── config.go
│   ├── handler/                # Servidor HTTP, handlers y renderer de plantillas
│   │   ├── application.go
│   │   ├── handlers.go
│   │   ├── renderer.go
│   │   └── router.go
│   ├── middleware/             # Middleware del proyecto
│   │   └── middleware.go       
│   └── storage/                # Capa de Datos
│       └── postgres/           # Lógica de la base de datos PostgreSQL
│           ├── migrations/     # Archivos de migración generados por Atlas
│           ├── queries/        # Consultas SQL para sqlc
│           │   └── movies.sql
│           ├── schema/         # Esquemas de la base de datos
│           │   └── schema.sql
│           ├── sqlc/           # Archivos generados por sqlc
│           │   ├── db.go
│           │   ├── models.go
│           │   └── movies.go
│           └── postgre.go
├── web/                        # Archivos estáticos y plantillas HTML
│   ├── static/
│   │   └── style.css
│   └── templates/
│          ├── index.html
│          ├── 404.html
│          └── 500.html
├── .air.toml                   # Configuración para Air (hot-reload)
├── .dockerignore               # Archivos a ignorar por Docker
├── .gitignore                  # Archivos a ignorar por Git
├── atlas.hcl                   # Configuración para Atlas (migraciones)
├── docker-compose.yml          # Definición de servicios Docker (app y db)
├── Dockerfile                  # Instrucciones para construir la imagen de la app
├── go.mod                      # Dependencias del proyecto Go
├── sqlc.yml                    # Configuración para sqlc
├── .env                        # Variables de entorno
└── Makefile                    # Centro de comandos para automatizar tareas
```

---
## Build

Sigue estos pasos para construir y ejecutar el proyecto en tu entorno local.

### Requisitos Previos

Antes de empezar, asegúrate de tener instaladas las siguientes herramientas en tu sistema.

#### 1. Herramientas del Sistema

* **Git:** Para clonar el repositorio.
* **Docker & Docker Compose:** Para ejecutar la base de datos y la aplicación en contenedores.
* **Go:** El lenguaje de programación (versión 1.21 o superior).
* **Make:** Para ejecutar los comandos automatizados del proyecto.

#### 2. Herramientas de Línea de Comandos de Go

Estas son herramientas de desarrollo que nos ayudan a automatizar tareas. Se instalan fácilmente con `go install`:

* **sqlc** (Generador de código para la base de datos):
    ```bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```
* **Air** (Recarga en caliente para desarrollo local):
    ```bash
    go install github.com/air-verse/air@latest
    ```
* **Atlas** (Herramienta de migraciones de base de datos):
    ```bash
    go install ariga.io/atlas/cmd/atlas@latest
    ```

### Instalación y Ejecución

1.  **Clona el repositorio:**
    ```bash
    git clone https://github.com/Matias914/Web-Page.git
    cd Web-Page
    ```

2.  **Crea tu archivo de entorno:**
    Copia el archivo de ejemplo `.env.example` a un nuevo archivo llamado `.env`. Este archivo es ignorado por Git y contiene tus secretos locales.
    ```bash
    cp .env.example .env
    ```

3.  **Inicia el entorno de producción:**
    Este único comando utiliza el `Makefile` para orquestar todo: levanta la base de datos y la aplicación en contenedores docker.
    ```bash
    make dev
    ```

4.  **¡Listo!**
    La aplicación estará corriendo y accesible en `http://localhost:8080`.
---
## Flujo de Trabajo

### Recarga en Caliente
Gracias a **Air**, cualquier cambio que guardes en un archivo `.go` o `.sql` disparará automáticamente la regeneración de código, la recompilación y el reinicio del servidor. Verás los cambios reflejados en segundos.

### Migraciones de Base de Datos
La evolución del esquema de la base de datos se gestiona con **Atlas**. El flujo de trabajo es el siguiente:
1.  **Modifica el esquema:** Realiza cambios en el archivo `internal/storage/postgres/schema/schema.sql`.
2.  **Genera una nueva migración:** Ejecuta `make migrate-new NAME=nombre_descriptivo_del_cambio`.
3.  **Aplica la migración:** Ejecuta `make migrate-up` para aplicar los cambios a tu base de datos.

---
## Comandos Disponibles

El `Makefile` es el centro de control del proyecto. Ejecuta `make help` para ver una lista completa y actualizada de todos los comandos disponibles.

```bash
$ make help
Uso: make [comando]

Comandos disponibles:
clean                Limpieza segura de artefactos de build locales.
clean-docker         Limpieza completa del PROYECTO ACTUAL en Docker.
db-down              Detiene los contenedores de la base de datos.
db-up                Inicia el contenedor de la base de datos.
dev                  Levanta la BD y el servidor para desarrollo.
help                 Muestra esta ayuda.
migrate-new          Crea un nuevo archivo de migración con Atlas.
migrate-up           Aplica todas las migraciones pendientes con Atlas.
server               Inicia el servidor con Air (hot-reload).
sqlc                 Genera el código Go a partir de las consultas SQL.
test                 Ejecuta las pruebas unitarias.
tidy                 Limpia y verifica las dependencias de Go.
```

---
## Desarrolladores

* **Ortiz Matias** - *Estudiante*
* **Leon Nicolas** - *Estudiante*

**Universidad Nacional del Centro de la Provincia de Buenos Aires (UNICEN)**
Facultad de Ciencias Exactas - Tandil, Buenos Aires.
Septiembre, 2025.