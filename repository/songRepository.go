package repository

import (
	"database/sql"
	"song-libary/models"
)

type SongRepository struct {
	DB *sql.DB
}

// NewSongRepository создает новый экземпляр репозитория
func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{DB: db}
}

// Save сохраняет песню в базе данных
func (r *SongRepository) Save(song *models.Song) error {
	query := "INSERT INTO songs (group_name, song_name, text) VALUES ($1, $2, $3) RETURNING id, created_at"
	return r.DB.QueryRow(query, song.GroupName, song.SongName, song.Text).Scan(&song.ID, &song.CreatedAt)
}
