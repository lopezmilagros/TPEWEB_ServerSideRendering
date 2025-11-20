#Levantar el servidor y la base de datos
# rm views/*_templ.go
export PATH := $(PATH):/home/eugenia/go/bin


up:
	@echo "Levantando la base de datos..."
	cd base_de_datos && docker compose up --build -d 

	@echo "Borrando templ viejos y regenerando..."
	
	rm views/*_templ.go
	templ generate

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

