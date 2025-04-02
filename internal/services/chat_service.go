package services

import "chatbot-backend/internal/models"

func GetChatResponse(req models.ChatRequest) models.ChatResponse {
	// Aquí podríamos hacer lógica más avanzada
	reply := "¡Hola! Soy tu chatbot, aún en construcción."
	return models.ChatResponse{Reply: reply}
}
