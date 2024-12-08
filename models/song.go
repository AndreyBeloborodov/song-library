package models

import "time"

// Song представляет сущность песни
type Song struct {
	ID          string    `json:"id"` // UUID
	GroupName   string    `json:"group_name"`
	SongName    string    `json:"song_name"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
	ReleaseDate string    `json:"release_date"`
	Link        string    `json:"link"`
}
