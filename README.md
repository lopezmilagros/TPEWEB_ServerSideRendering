# Descarga

Clonar el repositorio de git desde la terminal:

1. Crear una carpeta local en donde se guardará el repositorio

2. .Abrir una nueva terminal y Navegar hasta la carpeta creada con el comando: 'cd/ < ruta >'

3. Ejecute el siguiente comando:
```bash
git clone https://github.com/arriagaeugenia/Tp_programacion_web
```

---
# TPE Parte 3: Logica de negocio

Este módulo implementa la **logica de negocio** de la aplicación conectando la base de datos con el servidor.

---

## Estructura

```

base_de_datos/
├── db/               # Definiciones de queries y modelos generados
│   ├── queries/      # Consultas SQL
│   ├── schema/       # Esquema de la base de datos
│   └── sqlc/         # Código Go generado automáticamente
servidor/
├── html/             # Archivos estáticos 
│   ├── index.html
│   └── Imagenes/
└── main.go           # Código principal del servidor HTTP
go.mod            
go.sum
Makefile              # Archivo para levantar, testear y bajar la pagina
request.hurl          # Herramienta Hurl para testear los endpoints de la pagina.

```

---
## Requisitos
### Docker para linux:
Tener instalado Docker desde la pagina: https://docs.docker.com/compose/install/
O en la terminal, instalarlo con el comando: 

```bash
sudo apt update && sudo apt install docker.io docker-compose -y
```
Esta instalacion funciona para Ubuntu y Debian

### SQLC para linux:
```bash
sudo snap install sqlc
```

### Hurl para linux:
    Abrir la terminal y ejecutar:

```bash
$ INSTALL_DIR=/tmp
$ VERSION=7.0.0
$ curl --silent --location https://github.com/Orange-OpenSource/hurl/releases/download/$VERSION/hurl-$VERSION-x86_64-unknown-linux-gnu.tar.gz | tar xvz -C $INSTALL_DIR
$ export PATH=$INSTALL_DIR/hurl-$VERSION-x86_64-unknown-linux-gnu/bin:$PATH
```

## Uso


### 1. Ejecutar el servidor y la base de datos
Desde la carpeta `TP_PROGRAMACION_WEB/` correr: 

```bash
make up
```
---

### 2. Testear funcionamiento de endpoints

```bash
make test
```
---
Luego ir al archivo logs.txt para verificar el correcto funcionamiento del test. Deberias ver algo como esto:
```
Servidor escuchando en http://localhost:8080
Sirviendo archivos desde: ./html
2025/10/28 17:23:36 Creando una nueva obra...
2025/10/28 17:23:36 Obra creada exitosamente.
2025/10/28 17:23:36 Actualizando obra...
2025/10/28 17:23:36 Obra actualizada exitosamente.
2025/10/28 17:23:36 Borrando obra...
2025/10/28 17:23:36 Obra borrada exitosamente.
```

### 3. Bajar el servidor y la base de datos

```bash
make down
```
---

# TPE Parte 4: La capa de presentacion
### 1. Ejecutar el servidor y la base de datos
Desde la carpeta `TP_PROGRAMACION_WEB/` correr: 

```bash
make up
```
---

### 2. Acceder a la aplicación
Abrir en el navegador:  http://localhost:8080

### 3. Navegar en la aplicación web:

#### -Inicio: 
En el inicio se encuentra desplegada el catálogo de todas las obras de la galeria, en cada una de ellas se encuentra un botón de eliminar.

#### -Obras:
Acá se encuentra el mismo listado de obras que en el inicio.

#### -Añadir:
Acá se encuentra el formulario para agregar obras, luego de agregar una obra te redirige a la sección Obras de la aplicación actualizada.

#### -Exposición:
Acá se encuentran listadas las obras disponibles de la galería.

### 4. Bajar la aplicacion
```bash
make down
```
---

# TPE Parte 5:
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

## Integrantes
- Milagros Lopez
- Maria Eugenia Arriaga
- Bianca Rizzalli

