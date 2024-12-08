package models

// AddSongRequest представляет тело запроса для добавления новой песни
type AddSongRequest struct {
	Group       string `json:"group"`        // Название группы
	Song        string `json:"song"`         // Название песни
	Text        string `json:"text"`         // Текст песни
	ReleaseDate string `json:"release_date"` // Дата релиза
	Link        string `json:"link"`         // ссылка на песню
}

// UpdateSongRequest представляет тело запроса для изменения данных песни
type UpdateSongRequest struct {
	OldSongName    string `json:"old_song_name"`    // Название песни, которую нужно обновить
	OldGroup       string `json:"old_group"`        // Старое название группы
	NewGroup       string `json:"new_group"`        // Новое название группы
	NewSongName    string `json:"new_song_name"`    // Новое название песни
	NewText        string `json:"new_text"`         // Новый текст песни
	NewReleaseDate string `json:"new_release_date"` // Дата релиза
	NewLink        string `json:"new_link"`         // ссылка на песню
}

// FilterParams представляет параметры фильтрации и пагинации
type FilterParams struct {
	Group       string `json:"group"`        // Название группы
	SongName    string `json:"song"`         // Название песни
	Text        string `json:"text"`         // Текст песни (поиск по включению)
	ReleaseDate string `json:"release_date"` // Дата релиза
	Limit       int    `json:"limit"`        // Количество записей на страницу
	Offset      int    `json:"offset"`       // Смещение для пагинации
}
