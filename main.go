package main


import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


var id string = ""

func main(){
	r := mux.NewRouter()
    r.HandleFunc("/id/{id}", obtenerIp).Methods("GET")
    r.HandleFunc("/Ms", obtenerMensaje).Methods("POST")

    r.Use(loggingMiddleware)

    http.Handle("/", r)
}