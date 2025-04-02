package routes

import (
	"chatbot-backend/internal/controllers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/chat", controllers.ChatHandler)
}
