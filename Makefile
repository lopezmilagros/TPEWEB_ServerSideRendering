#Levantar el servidor y la base de datos

up:
	@echo "Levantando la base de datos..."
	cd base_de_datos && docker-compose up --build -d 	

	@echo "Generando archivos.templ"
	templ generate


	@echo "Levantando el servidor Go..."
	go run .  > logs.txt 2>&1 &

	@echo "Generando codigo SQLC..."
	cd base_de_datos && sqlc generate

down:
	@echo "Deteniendo la base de datos..."
	cd base_de_datos && docker-compose stop

	@echo "Deteniendo servidor Go..."
	cd servidor && kill -9 $(shell lsof -t -i :8080)

