package models

// AddSongRequest представляет тело запроса для добавления новой песни
type AddSongRequest struct {
	Group string `json:"group"` // Название группы
	Song  string `json:"song"`  // Название песни
	Text  string `json:"text"`  // Текст песни
}
