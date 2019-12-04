package main

import (
	"net/http"
	"log"
	"encoding/json"
)


func obtenerIp(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func obtenerMensajes(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Ip: %v\n", vars["id"])
}