package controllers

import (
	"chatbot-backend/internal/models"
	"chatbot-backend/internal/services"
	"encoding/json"
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req models.ChatRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Error en la petición", http.StatusBadRequest)
		return
	}

	response := services.GetChatResponse(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
