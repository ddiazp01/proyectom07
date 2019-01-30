package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathEnvioPeticion Ruta de envío de peticiones
const PathEnvioPeticion string = "/envio"

//PathEnvioIdioma Ruta de envío de idioma
const PathEnvioIdioma string = "/envioIdioma"

//PathListadoIdioma Ruta de obtención de las peticiones
const PathListadoIdioma string = "/listadoIdioma"

//PathEnvioTraduccion Ruta de envío de traduccion
const PathEnvioTraduccion string = "/envioTraduccion"

//PathListadoPeticiones Ruta de obtención de las peticiones de hoy
const PathListadoPeticiones string = "/lista"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Manejadores Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathEnvioPeticion] = Insert
	Manejadores[PathEnvioIdioma] = InsertIdioma
	Manejadores[PathListadoIdioma] = ListIdioma
	Manejadores[PathEnvioTraduccion] = Insert
	Manejadores[PathListadoPeticiones] = List
}
