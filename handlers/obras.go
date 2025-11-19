package handlers

import (
	"database/sql"
	"log"
	"net/http"

	db "galeriadearte.com/base_de_datos/db/sqlc"
	"galeriadearte.com/models"
	"galeriadearte.com/views"
)

type ObraHandlerType struct {
	queries *db.Queries
}

func ObraHandler(q *db.Queries) *ObraHandlerType {
	return &ObraHandlerType{queries: q}
}

// Router principal para este handler
func (h *ObraHandlerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("➡️ Entró a ServeHTTP con ruta:", r.URL.Path, "método:", r.Method)

	switch {
	case r.URL.Path == "/" && r.Method == http.MethodGet:
		h.inicio(w, r)
	case r.URL.Path == "/administrar" && r.Method == http.MethodPost:
		h.CrearObra(w, r)
	case r.URL.Path == "/exposiciones" && r.Method == http.MethodGet:
		{ //lista obras disponibles
			h.inicio(w, r)
		}
	case r.URL.Path == "/listarObras" && r.Method == http.MethodGet:
		{ //lista todas las obras
			h.listarObras(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}

// GET /
func (h *ObraHandlerType) inicio(w http.ResponseWriter, r *http.Request) {
	obras, err := h.queries.ListObras(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var obrasResponse []models.Obra
	for _, obra := range obras {
		obrasResponse = append(obrasResponse, models.Obra{
			ID:          obra.ID,
			Titulo:      obra.Titulo,
			Artista:     obra.Artista,
			Descripcion: nullStringToString(obra.Descripcion),
			Precio:      obra.Precio,
			Vendida:     nullBoolToString(obra.Vendida),
		})
	}

	// Render de la página principal con templ
	if err := views.ObraPage(obrasResponse).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ObraHandlerType) listarObras(w http.ResponseWriter, r *http.Request) {
	obras, err := h.queries.ListObras(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var obrasResponse []models.Obra
	for _, obra := range obras {
		obrasResponse = append(obrasResponse, models.Obra{
			ID:          obra.ID,
			Titulo:      obra.Titulo,
			Artista:     obra.Artista,
			Descripcion: nullStringToString(obra.Descripcion),
			Precio:      obra.Precio,
			Vendida:     nullBoolToString(obra.Vendida),
		})
	}

	// Render de la página principal con templ
	if err := views.ObraList(obrasResponse).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// POST /obras  (manejo clásico de formulario + PRG)
func (h *ObraHandlerType) CrearObra(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parseando formulario", http.StatusBadRequest)
		return
	}

	titulo := r.FormValue("titulo")
	artista := r.FormValue("artista")
	descripcion := r.FormValue("descripcion")
	precio := r.FormValue("precio")
	vendidaStr := r.FormValue("vendida")

	var nuevaObra db.CreateObraParams
	nuevaObra.Titulo = titulo
	nuevaObra.Artista = artista
	nuevaObra.Precio = precio

	// Descripción opcional
	if descripcion != "" {
		nuevaObra.Descripcion = sql.NullString{String: descripcion, Valid: true}
	} else {
		nuevaObra.Descripcion = sql.NullString{String: "", Valid: false}
	}

	// Vendida: "true" / "false" desde el <select>
	if vendidaStr == "true" {
		nuevaObra.Vendida = sql.NullBool{Bool: true, Valid: true}
	} else {
		nuevaObra.Vendida = sql.NullBool{Bool: false, Valid: true}
	}

	_, err := h.queries.CreateObra(r.Context(), nuevaObra)
	if err != nil {
		http.Error(w, "Error creando obra", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Obra creada correctamente a través del formulario")

	// PRG: redirigir para evitar re-envío del form
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Helpers para convertir NullString / NullBool a string
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
		}
		return "Disponible"
	}
	return "-"
}

/*func (h *ObraHandlerType) obrasDisponibles(w http.ResponseWriter, r *http.Request) {
	obras, err := h.queries.ListObrasDisponibles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var obrasResponse []models.Obra
	for _, obra := range obras {
		obrasResponse = append(obrasResponse, models.Obra{
			ID:          obra.ID,
			Titulo:      obra.Titulo,
			Artista:     obra.Artista,
			Descripcion: nullStringToString(obra.Descripcion),
			Precio:      obra.Precio,
			Vendida:     nullBoolToString(obra.Vendida),
		})
	}

	// Render de la página principal con templ
	if err := views.ObraList(obrasResponse).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}*/
