package model

//Peticion struct
type Peticion struct {
	ID      int
	Palabra string
}

//Idioma struct
type Idioma struct {
	Nombre string
}

//Traductor struct
type Traductor struct {
	ID                int
	PalabraTraduccion string
}
