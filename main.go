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
	//LOGICA DE NEGOCIO
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

/*package main

import (
	"database/sql" // para sql.Open()
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"galeriadearte.com/handlers"

	_ "github.com/lib/pq"

	db "galeriadearte.com/base_de_datos/db/sqlc"
)

type ObraResponse struct {
	ID          int32  `json:"id"`
	Titulo      string `json:"titulo"`
	Artista     string `json:"artista"`
	Descripcion string `json:"descripcion"`
	Precio      string `json:"precio"`
	Vendida     string `json:"vendida"`
}

var dbQueries *db.Queries

func main() {

	//LOGICA DE NEGOCIO
	// Conexión a la base de datos
	connStr := "postgresql://milibianeuge:programacionweb@localhost:5432/db?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	dbQueries = db.New(dbConn)
	obrasHandler := handlers.ObraHandler(dbQueries)
	http.Handle("/", obrasHandler)

	// Puerto
	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	// Levantar servidor
	err3 := http.ListenAndServe(port, nil)
	if err3 != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}

}
func inicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Carpeta de archivos estáticos
	fmt.Print(w, "./html")
}

func administrar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Carpeta de archivos estáticos, envia el contenido del archivo administrar.html
	http.ServeFile(w, r, "./html/administrar.html")
}

func exposiciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Carpeta de archivos estáticos, envia el contenido del archivo exposiciones.html
	http.ServeFile(w, r, "./html/exposiciones.html")
}

func contenedorObras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Carpeta de archivos estáticos, envia el contenido del archivo contenedor.html
	http.ServeFile(w, r, "./html/contenedor.html")
}

func obrasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getObras(w, r)
	case http.MethodPost:
		createObra(w, r)
	case http.MethodPut:
		updateObra(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func obraHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid product ID",
			http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getObra(w, r, id)
	case http.MethodDelete:
		deleteObra(w, r, id)
	default:
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}

func getObras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//llamo a listar obras
	obras, err := dbQueries.ListObras(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var obrasResponse []ObraResponse
	for _, o := range obras {
		obrasResponse = append(obrasResponse, ObraResponse{
			ID:          o.ID,
			Titulo:      o.Titulo,
			Artista:     o.Artista,
			Descripcion: nullStringToString(o.Descripcion),
			Precio:      o.Precio,
			Vendida:     nullBoolToString(o.Vendida),
		})
	}
	// Convertir a JSON y enviar respuesta
	json.NewEncoder(w).Encode(obrasResponse)
}
func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func nullBoolToString(nb sql.NullBool) string {
	if nb.Valid {
		if nb.Bool {
			return "Vendida"
		} else {
			return "Disponible"
		}
	}
	return "-"
}

func createObra(w http.ResponseWriter, r *http.Request) {
	log.Println("Creando una nueva obra...")
	// Definir estructura para decodificar JSON
	type reqObra struct {
		Titulo      string `json:"titulo"`
		Descripcion string `json:"descripcion"`
		Artista     string `json:"artista"`
		Precio      string `json:"precio"`
		Vendida     string `json:"vendida"`
	}

	//creo variable de tipo Obra
	var reqobra reqObra
	// Decodificar JSON del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&reqobra)
	if err != nil {
		//http.Error(w, "Invalid inputtttt", http.StatusBadRequest)
		http.Error(w, fmt.Sprintf("%s: %v", "Error", err), http.StatusBadRequest)
		return
	}

	var nuevaObra db.CreateObraParams
	nuevaObra.Titulo = reqobra.Titulo
	nuevaObra.Artista = reqobra.Artista
	nuevaObra.Precio = reqobra.Precio

	//Solo si el usuario envia "true" en el campo Vendida, se setea como true
	//cualquier otro valor (incluido "" o "false") se setea como false
	if reqobra.Vendida == "true" {
		nuevaObra.Vendida = sql.NullBool{Bool: true, Valid: true}
	} else {
		nuevaObra.Vendida = sql.NullBool{Bool: false, Valid: true}
	}

	if reqobra.Descripcion != "" {
		nuevaObra.Descripcion = sql.NullString{String: reqobra.Descripcion, Valid: true}
	} else {
		nuevaObra.Descripcion = sql.NullString{String: "", Valid: false}
	}

	// Inserto en la base de datos
	createdObra, err := dbQueries.CreateObra(r.Context(), nuevaObra)
	if err != nil {
		http.Error(w, "Failed to create obra", http.StatusInternalServerError)
		return
	}
	log.Println("Obra creada exitosamente.")

	// Enviar respuesta con la obra creada
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdObra)
}

func getObra(w http.ResponseWriter, r *http.Request, id int) {
	obraBuscada, err := dbQueries.GetObraById(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Obra not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obraBuscada)
}

func updateObra(w http.ResponseWriter, r *http.Request) {
	log.Println("Actualizando obra...")
	//id de la obra a actualizar se encuentra dentro del Body.
	// Definir estructura para decodificar JSON
	type reqObra struct {
		Id          int32  `json:"id"`
		Titulo      string `json:"titulo"`
		Descripcion string `json:"descripcion"`
		Artista     string `json:"artista"`
		Precio      string `json:"precio"`
		Vendida     string `json:"vendida"`
	}

	//creo variable de tipo Obra
	var reqobra reqObra
	// Decodificar JSON del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&reqobra)
	if err != nil {
		//http.Error(w, "Invalid inputtttt", http.StatusBadRequest)
		http.Error(w, fmt.Sprintf("%s: %v", "Error", err), http.StatusBadRequest)
		return
	}

	var obraPorActualizar db.Obra
	//buscar la obra por id
	obraPorActualizar, err1 := dbQueries.GetObraById(r.Context(), reqobra.Id)
	if err1 != nil {
		http.Error(w, "Obra not found", http.StatusNotFound)
		return
	}

	var obraActualizada db.UpdateObraParams

	obraActualizada.ID = reqobra.Id
	if reqobra.Titulo != "" {
		obraActualizada.Titulo = reqobra.Titulo
	} else {
		obraActualizada.Titulo = obraPorActualizar.Titulo
	}
	if reqobra.Descripcion != "" {
		obraActualizada.Descripcion = sql.NullString{String: reqobra.Descripcion, Valid: true}
	} else {
		obraActualizada.Descripcion = obraPorActualizar.Descripcion
	}
	if reqobra.Artista != "" {
		obraActualizada.Artista = reqobra.Artista
	} else {
		obraActualizada.Artista = obraPorActualizar.Artista
	}
	if reqobra.Precio != "" {
		obraActualizada.Precio = reqobra.Precio
	} else {
		obraActualizada.Precio = obraPorActualizar.Precio
	}
	if reqobra.Vendida == "true" {
		obraActualizada.Vendida = sql.NullBool{Bool: true, Valid: true}
	} else {
		obraActualizada.Vendida = sql.NullBool{Bool: false, Valid: true}
	}

	err2 := dbQueries.UpdateObra(r.Context(), obraActualizada)
	if err2 != nil {
		http.Error(w, "Failed to update obra", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Obra actualizada exitosamente")
	log.Println("Obra actualizada exitosamente.")

}

func deleteObra(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("Borrando obra...")
	//borramos la obra
	err := dbQueries.DeleteObra(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete obra", http.StatusInternalServerError)
		return
	}
	// Respuesta de éxito
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Obra eliminada exitosamente")
	log.Println("Obra borrada exitosamente.")
}

func obrasDisponibles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//llamo a listar obras
	obras, err := dbQueries.ListAvailableObras(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var obrasResponse []ObraResponse
	for _, o := range obras {
		obrasResponse = append(obrasResponse, ObraResponse{
			ID:          o.ID,
			Titulo:      o.Titulo,
			Artista:     o.Artista,
			Descripcion: nullStringToString(o.Descripcion),
			Precio:      o.Precio,
			Vendida:     nullBoolToString(o.Vendida),
		})
	}
	// Convertir a JSON y enviar respuesta
	json.NewEncoder(w).Encode(obrasResponse)
}
*/
