package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// Estructura para la solicitud SOAP
type SoapRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName  xml.Name `xml:"Body"`
		SayHello struct {
			Name string `xml:"name"`
		} `xml:"SayHello"`
	} `xml:"Body"`
}

// Estructura para la respuesta SOAP
type SoapResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName     xml.Name `xml:"Body"`
		SayHelloRes struct {
			Return string `xml:"return"`
		} `xml:"SayHelloResponse"`
	} `xml:"Body"`
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	var soapRequest SoapRequest
	decoder := xml.NewDecoder(r.Body)
	err := decoder.Decode(&soapRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Procesar la solicitud SOAP y construir la respuesta
	response := SoapResponse{}
	response.Body.SayHelloRes.Return = "¡Hola " + soapRequest.Body.SayHello.Name + "!"

	// Codificar la respuesta como XML
	responseXml, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar la respuesta SOAP
	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml.Header)) // Agregar cabecera XML estándar
	w.Write(responseXml)
}

func main() {
	http.HandleFunc("/hello", sayHelloHandler)

	fmt.Println("Servicio SOAP en http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
