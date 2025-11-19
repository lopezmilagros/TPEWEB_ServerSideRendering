package models

import (
	"time"
)

type Obra struct {
	ID           int32
	Titulo       string
	Descripcion  string
	Artista      string
	FechaIngreso time.Time
	Precio       string
	Vendida      string
}
