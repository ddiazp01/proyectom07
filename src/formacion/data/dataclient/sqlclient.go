package dataclient

import (
	"database/sql"
	"fmt"
	"formacion/data/model"
	"time"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarPeticion test
func InsertarPeticion(objeto *model.Peticion) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@1234@tcp(localhost:3306)/Formacion")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO Peticion(palabra, fecha) VALUES (?, utc_timestamp())", objeto.Palabra)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ListarRegistros test
func ListarRegistros(objeto *model.Filtro) []model.RPeticion {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@1234@tcp(localhost:3306)/Formacion?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT * FROM Peticion WHERE (fecha <= '" + objeto.Fecha.Format(time.RFC3339) + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT * FROM Peticion WHERE (fecha >= ?)", objeto.Fecha.Format(time.RFC3339))

	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	resultado := make([]model.RPeticion, 0)
	for query.Next() {
		var fila = model.RPeticion{}

		err = query.Scan(&fila.ID, &fila.Palabra, &fila.Fecha)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}
