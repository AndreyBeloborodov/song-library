package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"song-libary/models"
	"song-libary/service"
	"strconv"
)

type SongHandler struct {
	Service *service.SongService
}

func NewSongHandler(service *service.SongService) *SongHandler {
	return &SongHandler{Service: service}
}

// AddSongHandler добавляет новую песню
// @Summary Add a new song
// @Description Add a new song to the music library
// @Tags songs
// @Accept json
// @Produce json
// @Param request body models.AddSongRequest true "Song details"
// @Success 201 {object} models.DefaultResponse
// @Failure 400 {object} models.DefaultResponse
// @Failure 500 {object} models.DefaultResponse
// @Router /songs/add [post]
func (h *SongHandler) AddSongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Received request to add a new song")

	if r.Method != http.MethodPost {
		log.Printf("[ERROR] Method not allowed: %s", r.Method)
		response := models.DefaultResponse{
			Message: "Method not allowed",
			Status:  http.StatusMethodNotAllowed,
		}
		h.writeJSONResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	var request models.AddSongRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("[ERROR] Failed to decode request body: %v", err)
		response := models.DefaultResponse{
			Message: "Invalid request body",
			Status:  http.StatusBadRequest,
		}
		h.writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	log.Printf("[DEBUG] Request data: %+v", request)

	_, err := h.Service.AddSong(request.Group, request.Song, request.Text, request.ReleaseDate, request.Link)
	if err != nil {
		log.Printf("[ERROR] Failed to save song to database: %v", err)
		response := models.DefaultResponse{
			Message: "Failed to save song to database",
			Status:  http.StatusInternalServerError,
		}
		h.writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := models.DefaultResponse{
		Message: "Song added successfully",
		Status:  http.StatusCreated,
	}

	log.Println("[INFO] Song added successfully")
	h.writeJSONResponse(w, http.StatusCreated, response)
}

// DeleteSongHandler удаляет песню по названию
// @Summary Delete a song by name
// @Description Delete a song from the music library by song name and group
// @Tags songs
// @Accept json
// @Produce json
// @Param song_name query string true "Song Name"
// @Param group query string true "Group Name"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.DefaultResponse
// @Failure 404 {object} models.DefaultResponse
// @Failure 500 {object} models.DefaultResponse
// @Router /songs/delete [delete]
func (h *SongHandler) DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Received request to delete song")

	if r.Method != http.MethodDelete {
		log.Printf("[ERROR] Method not allowed: %s", r.Method)
		response := models.DefaultResponse{
			Message: "Method not allowed",
			Status:  http.StatusMethodNotAllowed,
		}
		h.writeJSONResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Читаем параметр `song_name` из запроса
	songName := r.URL.Query().Get("song_name")
	group := r.URL.Query().Get("group")
	if songName == "" || group == "" {
		log.Printf("[ERROR] Missing query parameter: song_name, group")
		response := models.DefaultResponse{
			Message: "Missing query parameter: song_name, group",
			Status:  http.StatusBadRequest,
		}
		h.writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	log.Printf("[DEBUG] Song name to delete: %s", songName)

	// Вызываем сервис для удаления песни
	err := h.Service.DeleteSongByNameAndGroup(songName, group)
	if err != nil {
		if errors.Is(err, service.ErrSongNotFound) {
			log.Printf("[INFO] Song not found: %s", songName)
			response := models.DefaultResponse{
				Message: "Song not found",
				Status:  http.StatusNotFound,
			}
			h.writeJSONResponse(w, http.StatusNotFound, response)
			return
		}
		log.Printf("[ERROR] Failed to delete song: %v", err)
		response := models.DefaultResponse{
			Message: "Failed to delete song",
			Status:  http.StatusInternalServerError,
		}
		h.writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Формируем успешный ответ
	response := models.DefaultResponse{
		Message: "Song deleted successfully",
		Status:  http.StatusOK,
	}
	h.writeJSONResponse(w, http.StatusOK, response)
}

// UpdateSongHandler обновляет данные песни
// @Summary Update song details
// @Description Update the song details like group, song name, and text
// @Tags songs
// @Accept json
// @Produce json
// @Param request body models.UpdateSongRequest true "Updated song details"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.DefaultResponse
// @Failure 404 {object} models.DefaultResponse
// @Failure 500 {object} models.DefaultResponse
// @Router /songs/update [put]
func (h *SongHandler) UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Received request to update song")

	if r.Method != http.MethodPut {
		log.Printf("[ERROR] Method not allowed: %s", r.Method)
		response := models.DefaultResponse{
			Message: "Method not allowed",
			Status:  http.StatusMethodNotAllowed,
		}
		h.writeJSONResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	var request models.UpdateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("[ERROR] Failed to decode request body: %v", err)
		response := models.DefaultResponse{
			Message: "Invalid request body",
			Status:  http.StatusBadRequest,
		}
		h.writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	log.Printf("[DEBUG] Update request: %+v", request)

	err := h.Service.UpdateSong(request)
	if err != nil {
		if errors.Is(err, service.ErrSongNotFound) {
			response := models.DefaultResponse{
				Message: "Song not found for update",
				Status:  http.StatusNotFound,
			}
			h.writeJSONResponse(w, http.StatusNotFound, response)
			return
		}
		response := models.DefaultResponse{
			Message: "Failed to update song",
			Status:  http.StatusInternalServerError,
		}
		h.writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := models.DefaultResponse{
		Message: "Song updated successfully",
		Status:  http.StatusOK,
	}
	h.writeJSONResponse(w, http.StatusOK, response)
}

// GetSongsHandler обрабатывает запрос на получение песен с фильтрацией и пагинацией
// @Summary Get songs with filtering and pagination
// @Description Retrieve songs from the music library with optional filters and pagination
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Group Name"
// @Param song query string false "Song Name"
// @Param text query string false "Text of the song"
// @Param release_date query string false "Release Date"
// @Param limit query int false "Limit of songs per page" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} models.Song
// @Failure 400 {object} models.DefaultResponse
// @Failure 500 {object} models.DefaultResponse
// @Router /songs [get]
func (h *SongHandler) GetSongsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Received request to fetch songs")

	if r.Method != http.MethodGet {
		log.Printf("[ERROR] Method not allowed: %s", r.Method)
		response := models.DefaultResponse{
			Message: "Method not allowed",
			Status:  http.StatusMethodNotAllowed,
		}
		h.writeJSONResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Читаем параметры фильтрации и пагинации
	params := models.FilterParams{
		Group:       r.URL.Query().Get("group"),
		SongName:    r.URL.Query().Get("song"),
		Text:        r.URL.Query().Get("text"),
		ReleaseDate: r.URL.Query().Get("release_date"),
	}

	// Параметры пагинации
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // Значение по умолчанию
	}
	params.Limit = limit

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0 // Значение по умолчанию
	}
	params.Offset = offset

	log.Printf("[DEBUG] Filter and pagination params: %+v", params)

	// Вызываем сервис для получения песен
	songs, err := h.Service.GetSongs(params)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch songs: %v", err)
		response := models.DefaultResponse{
			Message: "Failed to fetch songs",
			Status:  http.StatusInternalServerError,
		}
		h.writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Отправляем результат
	h.writeJSONResponse(w, http.StatusOK, songs)
}

// GetSongTextHandler обрабатывает запрос на получение текста песни с пагинацией
// @Summary Get song text with pagination
// @Description Retrieve song text with pagination (verses)
// @Tags songs
// @Accept json
// @Produce json
// @Param song_name query string true "Song Name"
// @Param group query string true "Group Name"
// @Param limit query int false "Limit of verses per page" default(5)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {object} models.SongTextResponse
// @Failure 400 {object} models.DefaultResponse
// @Failure 404 {object} models.DefaultResponse
// @Failure 500 {object} models.DefaultResponse
// @Router /songs/text [get]
func (h *SongHandler) GetSongTextHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Received request to fetch song text with pagination")

	if r.Method != http.MethodGet {
		log.Printf("[ERROR] Method not allowed: %s", r.Method)
		response := models.DefaultResponse{
			Message: "Method not allowed",
			Status:  http.StatusMethodNotAllowed,
		}
		h.writeJSONResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Получаем параметры запроса
	songName := r.URL.Query().Get("song_name")
	group := r.URL.Query().Get("group")
	if songName == "" || group == "" {
		log.Printf("[ERROR] Missing query parameter: song_name, group")
		response := models.DefaultResponse{
			Message: "Missing query parameter: song_name, group",
			Status:  http.StatusBadRequest,
		}
		h.writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 5 // Значение по умолчанию
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0 // Значение по умолчанию
	}

	// Вызываем сервис для получения текста песни
	response, err := h.Service.GetSongText(songName, group, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrSongNotFound) {
			log.Printf("[INFO] Song not found: %s", songName)
			errorResponse := models.DefaultResponse{
				Message: "Song not found",
				Status:  http.StatusNotFound,
			}
			h.writeJSONResponse(w, http.StatusNotFound, errorResponse)
			return
		}

		log.Printf("[ERROR] Failed to fetch song text: %v", err)
		errorResponse := models.DefaultResponse{
			Message: "Failed to fetch song text",
			Status:  http.StatusInternalServerError,
		}
		h.writeJSONResponse(w, http.StatusInternalServerError, errorResponse)
		return
	}

	// Отправляем успешный ответ
	h.writeJSONResponse(w, http.StatusOK, response)
}

// InfoHandler обрабатывает запрос на получение информации о песне
// @Summary Get song details by group and song name
// @Description Retrieve details of a song, including release date, text, and link
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string true "Group Name"
// @Param song_name query string true "Song Name"
// @Success 200 {object} models.SongDetail
// @Failure 400 {object} models.DefaultResponse
// @Failure 500 {object} models.DefaultResponse
// @Router /songs/info [get]
func (h *SongHandler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Received request to get song info")

	if r.Method != http.MethodGet {
		log.Printf("[ERROR] Method not allowed: %s", r.Method)
		response := models.DefaultResponse{
			Message: "Method not allowed",
			Status:  http.StatusMethodNotAllowed,
		}
		h.writeJSONResponse(w, http.StatusMethodNotAllowed, response)
		return
	}

	// Получаем параметры из запроса
	group := r.URL.Query().Get("group")
	songName := r.URL.Query().Get("song_name")
	if group == "" || songName == "" {
		log.Printf("[ERROR] Missing required query parameters: group or song_name")
		response := models.DefaultResponse{
			Message: "Missing required query parameters: group or song_name",
			Status:  http.StatusBadRequest,
		}
		h.writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Вызываем сервис для получения информации о песне
	songDetail, err := h.Service.GetSongInfo(group, songName)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch song info: %v", err)
		response := models.DefaultResponse{
			Message: "Failed to fetch song info",
			Status:  http.StatusInternalServerError,
		}
		h.writeJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	h.writeJSONResponse(w, http.StatusOK, songDetail)
}

// writeJSONResponse отправляет JSON-ответ с заданным статусом
func (h *SongHandler) writeJSONResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("[ERROR] Failed to encode response: %v", err)
	}
}
