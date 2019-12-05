package main

import (
	"net/http"
    "log"
	"github.com/gorilla/mux"
	em "enviarMensajes"
)


func main(){
	r := mux.NewRouter()
    r.HandleFunc("/id/{id}", em.ObtenerId).Methods("GET")
    r.HandleFunc("/ip", em.ObtenerIp).Methods("GET")
    r.HandleFunc("/ms", em.ObtenerMensajes).Methods("POST")

    r.Use(em.LoggingMiddleware)

    go em.Enviar()

    log.Fatal(http.ListenAndServe(":8080", r))


}