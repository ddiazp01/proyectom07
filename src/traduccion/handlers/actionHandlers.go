package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	client "traduccion/data/dataclient"
	"traduccion/data/model"
)

//Insert Función que inserta una petición en la base de datos local
func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioPeticion {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var peticion model.Peticion
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &peticion)

		if peticion.Palabra != "" {
			peticion.Palabra = strings.ToUpper(peticion.Palabra)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición está vacía")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(peticion)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarPeticion(&peticion)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//List Función que devuelve las peticiones de la base de datos dado un filtro
func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListadoPeticiones {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	lista := client.ListarRegistros()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	respuesta, _ := json.Marshal(&lista)
	fmt.Fprint(w, string(respuesta))
}

//InsertIdioma Función que inserta una petición en la base de datos local
func InsertIdioma(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathEnvioIdioma {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	if e == nil {
		var idioma model.Idioma
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &idioma)

		if idioma.Nombre != "" {
			idioma.Nombre = strings.ToUpper(idioma.Nombre)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "La petición está vacía")
			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "application/json")

		respuesta, _ := json.Marshal(idioma)
		fmt.Fprint(w, string(respuesta))

		go client.InsertarIdioma(&idioma)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}
}

//ListIdioma Función que devuelve las peticiones de la base de datos dado un filtro
func ListIdioma(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListadoIdioma {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	lista := client.ListarIdiomas()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	respuesta, _ := json.Marshal(&lista)
	fmt.Fprint(w, string(respuesta))
}
