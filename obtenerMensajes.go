package enviarMensajes

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)
var id string
var ip string

func ObtenerIp(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    
    if vars["ip"] != ""{
    	ip = vars["ip"]
    }

    w.WriteHeader(http.StatusOK)
    log.Println("Ip entrante: ",vars["ip"])
}

func ObtenerId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    
    if vars["id"] != ""{
    	id = vars["id"]
    }

    w.WriteHeader(http.StatusOK)
    log.Println("Id entrante: ",vars["id"])
}

func ObtenerMensajes(w http.ResponseWriter, r *http.Request) {

	var mensaje string 

	err := json.NewDecoder(r.Body).Decode(&mensaje) 

	if err != nil{
		w.WriteHeader(500)
		return
	}

	log.Println(mensaje)

    w.WriteHeader(http.StatusOK)
}