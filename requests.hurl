# requests.hurl

# 1. Obtener todos las obras
GET http://localhost:8080/listarObras
HTTP 200


# 2. Crear una nueva obra
POST http://localhost:8080/obras
Content-Type: application/x-www-form-urlencoded

"titulo=Guernica&descripcion=Obra+famosa+de+Picasso&artista=Picasso&precio=100.0&vendida="

HTTP 303

[Asserts]
header "Location" contains "/listarObras?"

[Captures]
obra_id: header "Location" regex "obra_id=([0-9]+)"


# 4. Eliminar la obra recién creada
POST http://localhost:8080/eliminar
Content-Type: application/x-www-form-urlencoded

"id={obra_id}"

HTTP 303

[Asserts]
header "Location" == "/listarObras"

# 3. Chequear que NO está más en la lista
GET http://localhost:8080/listarObras
HTTP 200

[Asserts]
body not contains "data-id=\"{{obra_id}}\""

