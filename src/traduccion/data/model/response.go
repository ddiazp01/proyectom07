package model

//RPeticion struct
type RPeticion struct {
	ID      int
	Palabra string
}

// RIdioma struct
type RIdioma struct {
	ID     int
	Nombre string
}

// RTraductor struct
type RTraductor struct {
	ID                int
	PalabraTraduccion string
	Idioma_ID         int
	Peticion_ID       int
}
