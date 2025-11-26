package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "galeriadearte.com/base_de_datos/db/sqlc"
	"galeriadearte.com/handlers"

	_ "github.com/lib/pq"
)

var dbQueries *db.Queries

func main() {

	// Conexión a la base de datos
	connStr := "postgresql://milibianeuge:programacionweb@localhost:5432/db?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	dbQueries = db.New(dbConn)
	obrasHandler := handlers.ObraHandler(dbQueries)

	// Puerto
	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	// Servir archivos estáticos
	fileServer := http.FileServer(http.Dir("servidor/html"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Handler principal (templ)
	http.Handle("/", obrasHandler)

	// Levantar servidor
	err3 := http.ListenAndServe(port, nil)
	if err3 != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}

}


