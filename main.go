package main

import (
	"fmt"
	"log"
	"net/http"
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	// Endpoint que recibirá los mensajes
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Mensaje enviado exitosamente!"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método no permitido"))
	}
}

func main() {
	http.HandleFunc("/api/sendMessage", sendMessage)
	fmt.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
