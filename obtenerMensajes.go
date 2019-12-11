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
    vars,ok := r.URL.Query()["ip"]
    
    if ok && len(vars[0]) > 0{
    	ip = string(vars[0])
    	enviandoPrimer = false
    }

    w.WriteHeader(http.StatusOK)
    log.Println("Ip entrante: ",vars[0])
}

func ObtenerId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    
    if vars["id"] != ""{
    	id = vars["id"]
    	enviandoPrimer = false
    }

    w.WriteHeader(http.StatusOK)
    log.Println("Id entrante: ",vars["id"])
}

func ObtenerMensajes(w http.ResponseWriter, r *http.Request) {
	var mensaje map[string]string

	err := json.NewDecoder(r.Body).Decode(&mensaje) 

	if err != nil{
		log.Println("ERROR EN LA RESPUESTA: ",err)
		w.WriteHeader(500)
		return
	}

	log.Println("Este es el mensaje: ",mensaje["title"],": ",mensaje["body"])

    w.WriteHeader(http.StatusOK)
}