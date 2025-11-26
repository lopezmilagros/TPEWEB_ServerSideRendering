package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	log.Println("Entró a ServeHTTP con ruta:", r.URL.Path, "método:", r.Method)

	switch {
	case r.URL.Path == "/" && r.Method == http.MethodGet:
		h.inicio(w, r)
	case r.URL.Path == "/obras" && r.Method == http.MethodPost:
		h.CrearObra(w, r)
	case r.URL.Path == "/agregar" && r.Method == http.MethodGet:
		h.formularioAgregarObra(w, r)
	case r.URL.Path == "/actualizar" && r.Method == http.MethodGet:
		h.formularioActualizar(w, r)
	case r.URL.Path == "/exposiciones" && r.Method == http.MethodGet:
		{ //lista obras disponibles
			h.listarObrasDisponibles(w, r)
		}
	case r.URL.Path == "/listarObras" && r.Method == http.MethodGet:
		{ //lista todas las obras
			h.listarObras(w, r)
		}
	/*case r.URL.Path == "/eliminar" && r.Method == http.MethodPost:
		h.deleteObra(w, r)*/
	case r.URL.Path == "/update" && r.Method == http.MethodPost:
		h.updateObra(w, r)
	case strings.HasPrefix(r.URL.Path, "/obras/") && r.Method == http.MethodDelete:
    h.deleteObraHTMX(w, r)
	
	default:
		http.NotFound(w, r)
	}
}

func (h *ObraHandlerType) formularioAgregarObra(w http.ResponseWriter, r *http.Request) {
	if err := views.ObraForm().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *ObraHandlerType) formularioActualizar(w http.ResponseWriter, r *http.Request) {

	// Después obtiene las obras de la base
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

	// Y ahora sí renderiza la lista simple y formulario
	if err := views.ActualizarObraPage(obrasResponse).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ObraHandlerType) listarObrasDisponibles(w http.ResponseWriter, r *http.Request) {
	obras, err := h.queries.ListAvailableObras(r.Context())
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

	if err := views.ObrasDisponibles(obrasResponse).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	if err := views.ObraList(obrasResponse).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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

	nuevaObra.Vendida = sql.NullBool{
		Bool:  vendidaStr != "",
		Valid: true,
	}

	created, err := h.queries.CreateObra(r.Context(), nuevaObra)
	if err != nil {
		http.Error(w, "Error creando obra", http.StatusInternalServerError)
		return
	}

	log.Println("Obra creada correctamente a través del formulario", created)

	//volvemos a consultar la lista completa
	obras, err := h.queries.ListObras(r.Context())
    if err != nil {
        http.Error(w, "Error listando obras", http.StatusInternalServerError)
        return
    }

	//Convertimos a modelos adecuado y luego solo devolvemos el fragmeto de la lista, no el layout
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

    if err := views.ObraListContent(obrasResponse).Render(r.Context(), w); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
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
		}
		return "Disponible"
	}
	return "-"
}

func (h *ObraHandlerType) deleteObraHTMX(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    idStr := parts[len(parts)-1]

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    err = h.queries.DeleteObra(r.Context(), int32(id))
    if err != nil {
        http.Error(w, "Error eliminando obra", 500)
        return
    }

    w.WriteHeader(http.StatusOK) // respuesta vacía
}

/*func (h *ObraHandlerType) deleteObra(w http.ResponseWriter, r *http.Request) {
	log.Println("Borrando obra...")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parseando formulario", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("id")
	if idStr == "" {
		http.Error(w, "ID de obra no recibido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	//borramos la obra
	err2 := h.queries.DeleteObra(r.Context(), int32(id))
	if err2 != nil {
		http.Error(w, "Failed to delete obra", http.StatusInternalServerError)
		return
	}
	log.Println("Obra borrada correctamente")

	// Redirigimos a la lista de obras
	http.Redirect(w, r, "/listarObras", http.StatusSeeOther)
}*/

func (h *ObraHandlerType) updateObra(w http.ResponseWriter, r *http.Request) {
	log.Println("Actualizando obra...")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parseando formulario", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("id")
	if idStr == "" {
		http.Error(w, "ID de obra no recibido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	titulo := r.FormValue("titulo")
	artista := r.FormValue("artista")
	descripcion := r.FormValue("descripcion")
	precio := r.FormValue("precio")
	vendidaStr := r.FormValue("vendida")

	// Obtenemos obra existente
	obraActual, err := h.queries.GetObraById(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Obra no encontrada", http.StatusNotFound)
		return
	}

	// Creamos estructura para actualización

	var obraUpdate db.UpdateObraParams
	obraUpdate.ID = int32(id)

	// Título
	if titulo != "" {
		obraUpdate.Titulo = titulo
	} else {
		obraUpdate.Titulo = obraActual.Titulo
	}

	// Descripción
	if descripcion != "" {
		obraUpdate.Descripcion = sql.NullString{String: descripcion, Valid: true}
	} else {
		obraUpdate.Descripcion = obraActual.Descripcion
	}

	// Artista
	if artista != "" {
		obraUpdate.Artista = artista
	} else {
		obraUpdate.Artista = obraActual.Artista
	}

	// Precio
	if precio != "" {
		obraUpdate.Precio = precio
	} else {
		obraUpdate.Precio = obraActual.Precio
	}

	// Vendida
	switch vendidaStr {
	case "true":
		obraUpdate.Vendida = sql.NullBool{Bool: true, Valid: true}
	case "false":
		obraUpdate.Vendida = sql.NullBool{Bool: false, Valid: true}
	default:
		obraUpdate.Vendida = obraActual.Vendida
	}

	// Actualizar en DB
	if err := h.queries.UpdateObra(r.Context(), obraUpdate); err != nil {
		http.Error(w, "No se pudo actualizar la obra", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/listarObras", http.StatusSeeOther)
	log.Println("Obra actualizada exitosamente.")
}
