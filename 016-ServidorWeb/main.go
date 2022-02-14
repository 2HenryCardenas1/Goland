package main

import "net/http"

func main() {
	routes()
	http.ListenAndServe(":8080", nil) //parametros que pide => el puerto y http handler
}

func home(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./index.html")
}

func routes() {
	//Rutas
	http.HandleFunc("/", home)
}
