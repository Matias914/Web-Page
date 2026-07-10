# 🍿 Plataforma Web de Cine — Entrega TP6 (UI/UX & Formularios Modulares)

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Templ](https://img.shields.io/badge/Templ-Type__Safe__UI-00CDA6?style=for-the-badge&logo=go&logoColor=white)

> 📢 **Nota de la Entrega:** Esta rama contiene el código fuente consolidado correspondiente al **Trabajo Práctico 6 (TP6)** de la UNICEN. El objetivo de este hito es la refactorización de la capa de presentación para mejorar la usabilidad (UX). Se implementó un sistema de notificaciones en tiempo real manejado desde el servidor y se desacoplaron los formularios de administración en componentes especializados (Películas, Celebridades y Géneros) utilizando el motor Templ.

![Demostración de la Aplicación](docs/web-page.gif)

---

## 🏗️ Estructura del Proyecto

La capa de frontend en servidor continúa expandiéndose, ahora con soporte para interacciones granulares y retroalimentación visual al usuario.

```
.
├── cmd/
│   └── web/main.go             # Punto de entrada de la aplicación y cableado de dependencias
├── internal/                   # Código privado del ecosistema backend
│   ├── config/                 # Validación de variables de entorno
│   ├── middleware/             # Interceptores de peticiones HTTP
│   ├── service/                # Capa de Lógica de Negocio 
│   ├── storage/                # Motor de Datos (PostgreSQL, SQLC Binding y Atlas)
│   └── transport/              # Capa de Comunicación HTTP
│       ├── handlers_*.go       # Controladores segmentados por dominio de datos
│       ├── handlers_notification.go # Lógica de despacho de alertas y mensajes flash
│       ├── router.go           # Enrutador principal
│       └── views/              # 🎨 Componentes de Interfaz de Usuario (Templ Engine)
│           ├── edit_form_celebrities.templ # Formulario específico para Actores/Directores
│           ├── edit_form_genres.templ      # Formulario específico para Géneros
│           ├── edit_form_movies.templ      # Formulario específico para Películas
│           ├── notification.templ          # Sistema de alertas dinámicas (Toast/Flash)
│           ├── layout.templ                # Armazón HTML global
│           ├── page_*.templ                # Vistas principales de página
│           └── *_templ.go                  # Código Go autogenerado
├── web/                        
│   └── static/styles/          # Hojas de estilo modulares (CSS puro)
├── .air.toml                   # Configuración de recarga en caliente (Air)
├── Dockerfile                  # Instrucciones para contenedor de producción
├── Makefile                    # Catálogo de comandos de automatización
├── atlas.hcl                   # Declaración del motor de migraciones
└── docker-compose.yml          # Topología multi-contenedor (App + BD)
```

---

## 🛠️ Stack Tecnológico

* **Backend:** [Go (Golang)](https://go.dev/) (v1.21+).
* **Motor de Interfaz:** [Templ](https://templ.guide/), procesamiento de vistas web tipadas en Go.
* **Base de Datos:** [PostgreSQL](https://www.postgresql.org/) interactuando mediante sentencias tipadas de [SQLC](https://sqlc.dev/).
* **Migraciones:** [Atlas](https://atlasgo.io/) para control de versiones estructurales.
* **Desarrollo:** [Air](https://github.com/air-verse/air) para *hot-reload* continuo.

---

## 🚀 Despliegue Local

### Requisitos Previos

Asegúrate de contar con Docker, Docker Compose, GNU Make y Go v1.21+. Instala las utilidades globales ejecutando:

```bash
go install [github.com/sqlc-dev/sqlc/cmd/sqlc@latest](https://github.com/sqlc-dev/sqlc/cmd/sqlc@latest)
go install [github.com/air-verse/air@latest](https://github.com/air-verse/air@latest)
go install ariga.io/atlas/cmd/atlas@latest
go install [github.com/a-h/templ/cmd/templ@latest](https://github.com/a-h/templ/cmd/templ@latest)
```

### Configuración Rápida

1. **Clonar e inicializar el entorno:**
   ```bash
   git clone [https://github.com/Matias914/Web-Page.git](https://github.com/Matias914/Web-Page.git)
   cd Web-Page
   cp .env.example .env
   ```

2. **Despliegue Full-Stack Automático:**
   ```bash
   make prod
   ```
   La aplicación compilará todos sus componentes, ejecutará las migraciones e iniciará en: `http://localhost:8080`.

3. **Detener servicios:**
   ```bash
   make prod-down
   ```

---

## 🔄 Flujo de Trabajo

### Sistema de Componentes (Templ) y Hot-Reload
Mediante la ejecución de `make dev`, **Air** supervisará de forma conjunta los archivos `.go` y `.templ`. Al modificar la estructura visual de una notificación o un formulario, el compilador generará automáticamente el árbol tipado de Go y reiniciará el servidor de forma instantánea.

### Administración de Base de Datos
1. Modifica la estructura en `internal/storage/postgres/schema/schema.sql`.
2. Genera el diferencial: `make migrate-diff NAME=nombre_del_cambio`.
3. Aplica los cambios: `make migrate-up`.

---

## 📋 Panel de Comandos (Makefile)

Utiliza `make help` para auditar la suite de tareas integradas:

```bash
$ make help
Uso: make [comando]

Comandos Principales:
  dev           - Inicia BD, aplica migraciones y corre el servidor en desarrollo (hot-reload).
  prod          - Construye y levanta toda la infraestructura en contenedores Docker.
  prod-down     - Detiene los servicios en ejecución.
  server        - Corre localmente el servidor web aislando dependencias de BD.

Comandos de Base de Datos:
  db-up         - Levanta únicamente PostgreSQL en Docker.
  db-down       - Apaga el contenedor de BD.
  db-nuke       - Destruye contenedores y purga volúmenes persistentes.

Sistemas de Migraciones (Atlas):
  migrate-diff  - Genera un archivo incremental desde el esquema (requiere NAME).
  migrate-up    - Impacta migraciones pendientes a la BD activa.
  migrate-set   - Revierte el historial a una versión previa (requiere VERSION).

Desarrollo y Generación de Código:
  sqlc-gen      - Recrea las estructuras Go basadas en queries SQL.
  templ-gen     - Compila los archivos visuales .templ a código nativo Go.
  build         - Compila la aplicación en un binario.
  clean         - Limpia directorios de *build*.
```

---

## 👥 Desarrolladores

* **Ortiz Matias** - *Estudiante de Ingeniería de Sistemas (UNICEN)*
* **Leon Nicolas** - *Estudiante de Ingeniería de Sistemas (UNICEN)*

**Universidad Nacional del Centro de la Provincia de Buenos Aires (UNICEN)** Facultad de Ciencias Exactas — Tandil, Buenos Aires, Argentina.  
*Septiembre, 2025.*

![Banner del Proyecto](docs/banner.png)
