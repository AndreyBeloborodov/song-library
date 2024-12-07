package models

// AddSongResponse представляет ответ сервера после добавления песни
type AddSongResponse struct {
	Message string `json:"message"` // Сообщение о результате операции
	Status  int    `json:"status"`  // HTTP-статус операции
}
