package handlers

import (
	"encoding/json"
	"net/http"
	"song-libary/models"
	"song-libary/service"
)

type SongHandler struct {
	Service *service.SongService
}

// NewSongHandler создает новый хендлер
func NewSongHandler(service *service.SongService) *SongHandler {
	return &SongHandler{Service: service}
}

// AddSongHandler обрабатывает запрос на добавление песни
func (h *SongHandler) AddSongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request models.AddSongRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Вызываем метод сервиса
	_, err := h.Service.AddSong(request.Group, request.Song, request.Text)
	if err != nil {
		http.Error(w, "Failed to save song to database", http.StatusInternalServerError)
		return
	}

	// Формируем успешный ответ
	response := models.AddSongResponse{
		Message: "Song added successfully",
		Status:  http.StatusCreated,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
