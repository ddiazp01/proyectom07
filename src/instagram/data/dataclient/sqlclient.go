package dataclient

import (
	"database/sql"
	"fmt"
	"instagram/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarUser test
func InsertarUser(objeto *model.User) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("nombre: ", objeto.Nombre)

	defer db.Close()
	insert, err := db.Query("INSERT INTO User(Nombre, Usuario, Email, Password) VALUES (?, ?, ?, ?)", objeto.Nombre, objeto.Usuario, objeto.Email, objeto.Password)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}
