package enviarMensajes

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      
        log.Println(r.RequestURI," Actualmente ip: ",ip,", id: ",id)
       
        next.ServeHTTP(w, r)
    })
}