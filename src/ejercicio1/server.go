package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Peticion struct {
	ID      int
	Palabra string
	Fecha   time.Time
}

type Data struct {
	Quote string
}

func main() {

	puerto := 8080

	fmt.Println("Iniciando servidor", puerto)

	http.HandleFunc("/", iniciohandle)
	http.HandleFunc("/js/", jshandle)
	http.HandleFunc("/ejercicio", ejercicio)
	http.ListenAndServe(":"+strconv.Itoa(puerto), nil)

}

func iniciohandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibiendo petecion desde " + r.URL.EscapedPath())

	if r.URL.EscapedPath() != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return

	}
	http.ServeFile(w, r, "pages/index.html")

}

func jshandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibiendo petecion desde " + r.URL.EscapedPath())

	/*if r.URL.EscapedPath() != "/js/" {
		http.NotFound(w, r)
		return
	}
	*/
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return

	}
	http.ServeFile(w, r, "js/base.js")

}

func ejercicio(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibiendo petecion desde " + r.URL.EscapedPath())

	if r.URL.Path != "/ejercicio" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return

	}

	defer r.Body.Close()
	byte, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var peticion Peticion
		json.Unmarshal(byte, &peticion)

		if peticion.Palabra != "" {

			peticion.Palabra = strings.ToUpper(peticion.Palabra)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, e)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		respuesta, _ := json.Marshal(peticion)
		fmt.Fprint(w, string(respuesta))

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, e)
	}

}
