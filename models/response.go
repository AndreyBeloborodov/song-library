package models

// DefaultResponse представляет тандартный ответ сервера
type DefaultResponse struct {
	Message string `json:"message"` // Сообщение о результате операции
	Status  int    `json:"status"`  // HTTP-статус операции
}

// SongTextResponse представляет ответ с текстом песни
type SongTextResponse struct {
	SongName string   `json:"song_name"` // Название песни
	Group    string   `json:"group"`     // Название группы
	Verses   []string `json:"verses"`    // Список куплетов
	Limit    int      `json:"limit"`     // Количество куплетов на страницу
	Offset   int      `json:"offset"`    // Смещение
	Total    int      `json:"total"`     // Общее количество куплетов
}

// SongDetail представляет информацию о песне
type SongDetail struct {
	ReleaseDate string `json:"release_date"` // Дата релиза
	Text        string `json:"text"`         // Текст песни
	Link        string `json:"link"`         // Ссылка на песню (например, на YouTube)
}
