package model

import "time"

//RUser struct
type RUser struct {
	ID       int
	Nombre   string
	Usuario  string
	Email    string
	Password string
}

//RFoto struct
type RFoto struct {
	Nomnre       string
	Texto        string
	URL          string
	Fecha        time.Time
	ComentarioID int
}

//RComentario struct
type RComentario struct {
	Texto      string
	Comentario string
}
