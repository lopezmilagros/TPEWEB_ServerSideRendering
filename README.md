# Descarga

Clonar el repositorio de git desde la terminal:

1. Crear una carpeta local en donde se guardará el repositorio

2. Abrir una nueva terminal y Navegar hasta la carpeta creada con el comando: 'cd/ < ruta >'

3. Ejecute el siguiente comando:

```bash
git clone https://github.com/lopezmilagros/TPEWEB_ServerSideRendering.git
```

---

# TPE Parte 6: Interfaces Reactivas con HTMX

Implementación de interfaces reactivas utilizando HTMX sobre un backend en Go. Se modificó el flujo tradicional de petición-redirección para utilizar intercambio de fragmentos HTML, logrando que la creación, eliminación y actualizacion de registros se refleje instantáneamente en la interfaz sin refrescar el navegador.

## Estructura

```

base_de_datos/
├── db/               # Definiciones de queries y modelos generados
│   ├── queries/      # Consultas SQL
│   ├── schema/       # Esquema de la base de datos
│   └── sqlc/         # Código Go generado automáticamente
models/
├── obras.go                    # Entidades del negocio: Obra
views/                          # Capa de presentacion: .templ: plantillas originales, .go codigo generdo
├── entity_form.templ           # Formulario
├── entity_form_templ.go
├── entity_list.templ           # Lista de obras
├── entity_list_templ.go
├── obras.templ                 # Inicio
├── obras_templ.go
├── layout.templ                # Renderizacion
├── layout_templ.go
├── styles.css
handlers/
├── obras.go                    # Manejo de rutas. Funciones para cada ruta.
servidor/html
├── styles.css                  # Estilos de los componentes
├── imagenes/
go.mod
go.sum
main.go               # Código principal del servidor y conexion con la base de datos.
Makefile              # Archivo para levantar, testear y bajar la pagina

```

---

## Requisitos

Las instrucciones están escritas para Debian/Ubuntu. En caso de usar otro sistema operativo, buscá los comandos específicos para tu sistema.

### Docker para linux:

Tener instalado Docker desde la pagina: https://docs.docker.com/compose/install/
O en la terminal, instalarlo con el comando:

```bash
sudo apt update && sudo apt install docker.io docker-compose -y
```

Esta instalacion funciona para Ubuntu y Debian

### SQLC para linux:

Abrir la terminal y ejecutar:

```bash
sudo snap install sqlc
```

### Make

Abrir la terminal y ejecutar:

```bash
sudo apt install make
```

### Templ

Abrir la terminal y ejecutar:

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

## Uso

### 1. Ejecutar el servidor y la base de datos

Desde la carpeta `TPEWEB_ServerSideRendering/` correr:

```bash
make up
```

---

### 2. Acceder a la aplicación

Abrir en el navegador: http://localhost:8080

### 3. Navegar en la aplicación web:

#### Inicio:

En el inicio se encuentra desplegada el catálogo de todas las obras de la galeria.

#### Obras:

Acá se encuentra el mismo listado de obras que en el inicio.

#### Exposición:

Se listan las obras disponibles de la galería.

#### Añadir:

Se renderiza el formulario para agregar obras, luego de agregar una obra se renderizara la nueva entidad automaticamente en la lista de obras.

#### Actualizar:

Aca se encuentra la lista de obras con su id, para actualizar una obra se debe indicar el id en el formulario y completar los campos que quieras actualizar. El resultado se vera atomaticamente en la pagina sin ser recargada.

- Siempre que se listan las obras, está la opcion de eliminarla. Al eliminarla desaparece automaticamente de la ista de obras sin tener que recargar la pagina.

### 4. Bajar la aplicacion

```bash
make down
```

---

## Integrantes

- Milagros Lopez
- Maria Eugenia Arriaga
- Bianca Rizzalli
