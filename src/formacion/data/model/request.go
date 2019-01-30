package model

import "time"

//Peticion struct
type Peticion struct {
	Palabra string
}

//Filtro struct
type Filtro struct {
	Fecha time.Time
}
