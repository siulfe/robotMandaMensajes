package enviarMensajes

import (
    "math/rand"
    "time"
    "net/http"
    "log"
    "encoding/json"
    "fmt"
    "bytes"
)

type mensajes struct{
    body string
    title string
}

var enviandoPrimer bool

func Enviar() {
    
    for {

        time.Sleep(1 * time.Second)

        if(id == "" || ip == ""){
            continue
        }

        mensaje := obtenerMensaje()
        
        body, _ := json.Marshal(mensaje)
       
        resp,err := http.Post(ip,"application/json",bytes.NewBuffer(body))

        if err != nil{
            continue
        }

        defer resp.Body.Close()

        if !enviandoPrimer{
            enviandoPrimer = true
            log.Println("Enviando a la IP: ",ip,", con id: ",id, " Codigo de la respuesta: ", resp.StatusCode)
        }

    }

}    

func obtenerMensaje() mensajes{
    mensaje := &mensajes{
        body: "",
        title: "",
    }

    param := rand.Intn(10)

    mensaje.body = obtenerBody(param)

    mensaje.title = obtenerTitle(param)

    return *mensaje
}


func obtenerBody(indice int) string{

    switch indice{
        case 0: return "Se le ha mandado un mensaje al Email"
        case 1: return "Una transacción se ha realizado con exito, por parte de un miembro de la empresa"
        case 2: return "La factura número: "+obtenerNumeroFactura()+", ha sido rechazada."
        case 3: return "El equipo de TuFactoring le desea una feliz navidad"
        case 4: return "Usar TuFactoring mejora las ganancias de las empresas y de los inversionistas"
        case 5: return "TuFactoring posee la mayor cantidad de tecnologias de ultima generación en le mercado"
        case 6: return "La factura número: " + obtenerNumeroFactura() + ", se encuentra actualmente en procesos de pago, le felicitamos por usar TuFactoring"
        case 7: return "Factura número: " + obtenerNumeroFactura() + ", Ha sido la factura con más ofertas en la historia de TuFactoring"
        case 8: return "Factura n+umero: " + obtenerNumeroFactura() + ", fue la primera factura con un valor de más de un billon registrada en TuFactoring"
        case 9: return "Navidad"
    }

    return "default"
}

func obtenerTitle(indice int) string{

    switch indice{
        case 0: return "Información del Sistema"
        case 1: return "Otros Usuarios"
        case 2: return "Información del Sistema"
        case 3: return "TuFactoring le Recuerda"
        case 4: return "TuFactoring le Recuerda"
        case 5: return "TuFactoring le Recuerda"
        case 6: return "Mensajes del Sistema"
        case 7: return "Mensajes del Sistema"
        case 8: return "Mensajes del Sistema"
        case 9: return "Un Mensaje que dice Navidad"
    }

    return "default"
}


func obtenerNumeroFactura() string{
    return fmt.Sprintf("%f.0",(rand.Float64() * 9000) + 1000) + "-" + fmt.Sprintf("%f.0",(rand.Float64() * 90000000) + 10000000) 
}