package main

import (
	"chatbot-backend/internal/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
