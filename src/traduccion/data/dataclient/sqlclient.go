package dataclient

import (
	"database/sql"
	"fmt"
	"traduccion/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarPeticion test
func InsertarPeticion(objeto *model.Peticion) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traduccion")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO Peticion(palabra) VALUES (?)", objeto.Palabra)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ListarRegistros test
func ListarRegistros() []model.RPeticion {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traduccion?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT * FROM Peticion)"
	fmt.Println(comando)
	query, err := db.Query("SELECT * FROM Peticion )")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	resultado := make([]model.RPeticion, 0)
	for query.Next() {
		var fila = model.RPeticion{}
		err = query.Scan(&fila.ID, &fila.Palabra)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}

//InsertarIdioma test
func InsertarIdioma(objeto *model.Idioma) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traduccion")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO Idioma(nombre) VALUES (?)", objeto.Nombre)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ListarIdiomas test
func ListarIdiomas() []model.RIdioma {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traduccion?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT * FROM Idioma"
	fmt.Println(comando)
	query, err := db.Query("SELECT * FROM Idioma")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	resultado := make([]model.RIdioma, 0)
	for query.Next() {
		var fila = model.RIdioma{}
		err = query.Scan(&fila.ID, &fila.Nombre)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}

//Traduccion test
func Traduccion(objeto *model.Traductor) []model.RTraductor {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Traduccion?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT t.ID, p.PalabraTraduccion, Traduccion, i.Idioma, FROM Palabra p INNER JOIN Traduccion t ON  p.ID = t.Palabra_ID = i.ID WHERE p.Palabra = 'palabra' AND Idioma_ID = 'ID'"
	fmt.Println(comando)
	query, err := db.Query("SELECT t.ID, p.PalabraTraduccion, Traduccion, i.Idioma, FROM Palabra p INNER JOIN Traduccion t ON  p.ID = t.Palabra_ID = i.ID WHERE p.Palabra = 'palabra' AND Idioma_ID = 'ID'")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	resultado := make([]model.RTraductor, 0)
	for query.Next() {
		var fila = model.RTraductor{}
		err = query.Scan(&fila.ID, &fila.PalabraTraduccion)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, fila)
	}
	return resultado
}
