#Levantar el servidor y la base de datos
up:
	@echo "Levantando la base de datos..."
	cd base_de_datos && docker compose up --build -d 

	@echo "Levantando el servidor Go..."
	go run .  > logs.txt 2>&1 &

	@echo "Generando codigo SQLC..."
	cd base_de_datos && sqlc generate

test:
	@echo "Ejecutando pruebas Hurl..."
	hurl --test requests.hurl 

down:
	@echo "Deteniendo la base de datos..."
	cd base_de_datos && docker-compose stop

	@echo "Deteniendo servidor Go..."
	cd servidor && kill -9 $(shell lsof -t -i :8080)

aux: 
	make down
	rm views/*_templ.go
	templ generate
	make up