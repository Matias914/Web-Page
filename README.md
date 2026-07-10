# 🍿 Plataforma Web de Cine — Entrega TP5 (Component-Driven Architecture)

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Templ](https://img.shields.io/badge/Templ-Type__Safe__UI-00CDA6?style=for-the-badge&logo=go&logoColor=white)

> 📢 **Nota de la Entrega:** Esta rama contiene el código fuente consolidado correspondiente al **Trabajo Práctico 5 (TP5)** de la UNICEN. El núcleo de este hito radica en la migración total hacia una **Arquitectura Basada en Componentes de Servidor**. Reemplazamos las plantillas HTML débiles clásicas y los scripts pesados del lado del cliente por componentes fuertemente tipados con **Templ**, compilando la interfaz de usuario directamente dentro del binario de Go para máxima eficiencia y seguridad en tiempo de compilación.

![Demostración de la Aplicación](docs/web-page.gif)

---

## 🏗️ Estructura del Proyecto

La capa de frontend ahora vive encapsulada dentro de la infraestructura de transporte del servidor, permitiendo un desarrollo acoplado al tipo de datos nativo de Go.

```
.
├── cmd/
│   └── web/main.go             # Punto de entrada de la aplicación y cableado de servicios
├── internal/                   # Código privado del ecosistema backend y vistas compiladas
│   ├── config/                 # Validación de variables de entorno del sistema
│   ├── middleware/             # Interceptores de peticiones HTTP
│   ├── service/                # Capa de Lógica de Negocio y validaciones relacionales
│   ├── storage/                # Motor de Datos (PostgreSQL, SQLC Binding y Atlas Migrations)
│   └── transport/              # Capa de Comunicación e Interfaz de Usuario
│       ├── handlers.go         # Controladores base y lógica de ruteo genérica
│       ├── handlers_celebrities.go # Endpoints y lógica de negocio para celebridades
│       ├── handlers_genres.go      # Endpoints y renderizado para géneros cinematográficos
│       ├── handlers_movies.go      # Endpoints y despacho lógico de carteleras
│       ├── router.go           # Declaración centralizada de rutas de la aplicación
│       └── views/              # 🎨 Componentes de Interfaz de Usuario (Templ Engine)
│           ├── 404.templ       # Componente de vista para recursos no encontrados
│           ├── 500.templ       # Componente de vista para fallos críticos de backend
│           ├── layout.templ    # Armazón HTML global (Navbar, Head, footer común)
│           ├── management_form.templ # Formularios modulares para administración (ABM)
│           ├── movies_card.templ     # Tarjeta de película reutilizable en grillas
│           ├── movies_catalog.templ  # Contenedor dinámico de grillas de películas
│           ├── page_catalog.templ    # Vista principal de la cartelera pública
│           ├── page_control.templ    # Vista de la consola de administración central
│           ├── page_index.templ      # Vista de la página de inicio de la plataforma
│           └── *_templ.go      # Código Go nativo auto-generado por el compilador Templ
├── web/                        # Activos del Frontend del Servidor
│   └── static/                 # Recursos Estáticos No Compilables
│       └── styles/             # Hojas de estilo modulares (CSS puro)
├── .air.toml                   # Orquestación de live-reloading continuo
├── Dockerfile                  # Instrucciones de empaquetado del binario contenedorizado
├── Makefile                    # Catálogo unificado de comandos de desarrollo y producción
├── atlas.hcl                   # Declaración de contexto para el motor de Atlas
├── sqlc.yaml                   # Configuración del compilador SQL para Go
└── docker-compose.yml          # Topología de servicios en contenedores (App Engine + PostgreSQL)
```

---

## 🛠️ Stack Tecnológico Core

* **Backend & Compilación de UI:** [Go (Golang)](https://go.dev/) (v1.21+).
* **Motor de Vistas:** [Templ](https://templ.guide/), procesamiento de maquetado interactivo fuertemente tipado que elimina errores sintácticos de HTML en tiempo de ejecución.
* **Persistencia Relacional:** [PostgreSQL](https://www.postgresql.org/) con mapeo de queries seguras vía [SQLC](https://sqlc.dev/).
* **Evolución de Base de Datos:** [Atlas](https://atlasgo.io/) para la gestión integral de cambios de esquema mediante migración declarativa.
* **Refresco en Caliente:** [Air](https://github.com/air-verse/air), configurado para vigilar tanto lógica de negocio (`.go`) como vistas estáticas y componentes (`.templ`).

---

## 🚀 Despliegue Local del Entorno

### Requisitos Previos

Instala las herramientas fundamentales del sistema (*Docker, Go v1.21+, GNU Make*) y añade los siguientes compiladores de desarrollo a las utilidades globales de tu Go Toolchain:

```bash
go install [github.com/sqlc-dev/sqlc/cmd/sqlc@latest](https://github.com/sqlc-dev/sqlc/cmd/sqlc@latest)
go install [github.com/air-verse/air@latest](https://github.com/air-verse/air@latest)
go install ariga.io/atlas/cmd/atlas@latest
go install [github.com/go-templ/templ/cmd/templ@latest](https://github.com/go-templ/templ/cmd/templ@latest)
```
> *Asegúrate de mapear los binarios instalados en tu entorno:* `export PATH=$PATH:$HOME/go/bin`

### Configuración e Inicio Rápido

1. **Clonar e inicializar variables secretas:**
   ```bash
   git clone [https://github.com/Matias914/Web-Page.git](https://github.com/Matias914/Web-Page.git)
   cd Web-Page
   cp .env.example .env
   ```

2. **Ejecutar el pipeline de producción:**
   Este comando compila las vistas estructuradas de Templ, genera el mapeo relacional de base de datos, monta las imágenes en Docker y levanta el entorno unificado:
   ```bash
   make prod
   ```
   La aplicación full-stack estará disponible de inmediato en: `http://localhost:8080`.

3. **Apagar los servicios activos:**
   ```bash
   make prod-down
   ```

---

## 🔄 Flujo de Trabajo en Desarrollo (Ecosistema Templ)

### Ciclo de Recarga Automatizado (Hot-Reload)
Cuando ejecutas `make dev`, el binario de **Air** no solo vigila los cambios de tus funciones de Go, sino que además ejecuta por detrás el transpiler de **Templ**. Al guardar un archivo `.templ`, la UI se recompila en un archivo `.go` equivalente y el servidor web se reinicia de manera transparente en milisegundos.

### Evolución de Datos con Atlas
1. Realiza los cambios estructurales que desees sobre el archivo maestro en `internal/storage/postgres/schema/schema.sql`.
2. Genera la migración incremental: `make migrate-diff NAME=descripcion_del_cambio`.
3. Sincroniza la base de datos local del contenedor corriendo: `make migrate-up`.

---

## 📊 Diseño de la Base de Datos

El motor relacional persiste las estructuras de películas, géneros cinematográficos, usuarios y las interacciones de reseñas/roles de las celebridades:

![Esquema de Base de Datos Relacional](docs/db.png)

---

## 📋 Comandos del Makefile (Centro de Control)

Ejecuta `make help` en tu terminal para auditar la suite de tareas automatizadas integradas:

```bash
$ make help
Uso: make [comando]

Comandos Principales:
  dev           - Inicia DB, aplica migraciones de Atlas y corre el servidor en desarrollo con hot-reload (Air).
  prod          - Construye de cero y levanta toda la infraestructura (app y base de datos) en Docker.
  prod-down     - Detiene los servicios de producción y remueve recursos huérfanos.
  server        - Corre el servidor web local con soporte de recarga en caliente vía Air.

Comandos de Base de Datos (Docker):
  db-up         - Levanta de manera aislada el contenedor de PostgreSQL.
  db-down       - Detiene el contenedor de almacenamiento.
  db-nuke       - Destruye contenedores y limpia volúmenes físicos de datos persistentes.

Comandos de Migraciones (Atlas):
  migrate-diff  - Compara el esquema declarativo y genera un nuevo archivo incremental (Requires NAME).
  migrate-up    - Impacta todas las migraciones estructuradas pendientes.
  migrate-set   - Fuerza el puntero de migración a una versión histórica específica (Requires VERSION).

Comandos de Desarrollo y Compilación:
  sqlc-gen      - Compila sentencias SQL en código Go fuertemente tipado.
  templ-gen     - Compila los componentes de vistas de interfaz (*.templ) a estructuras nativas de Go.
  build         - Compila de forma unificada el binario ejecutable optimizado.
  run           - Compila y ejecuta la aplicación nativamente en el host.
  tidy          - Limpia y actualiza las dependencias registradas en el go.mod.
  clean         - Remueve los directorios y binarios resultantes de compilaciones previas.
  docker-clean  - Limpieza controlada de imágenes y contenedores vinculados al proyecto.
  docker-nuke   - Purga absoluta del caché de Docker y volúmenes inactivos.
```

---

## 👥 Desarrolladores

* **Ortiz Matias** - *Estudiante de Ingeniería de Sistemas (UNICEN)*
* **Leon Nicolas** - *Estudiante de Ingeniería de Sistemas (UNICEN)*

**Universidad Nacional del Centro de la Provincia de Buenos Aires (UNICEN)** Facultad de Ciencias Exactas — Tandil, Buenos Aires, Argentina.  
*Septiembre, 2025.*

![Banner del Proyecto](docs/banner.png)
