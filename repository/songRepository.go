package repository

import (
	"database/sql"
	"log"
	"song-libary/models"
)

type SongRepository struct {
	DB *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{DB: db}
}

func (r *SongRepository) SaveSong(song *models.Song) error {
	log.Printf("[INFO] Saving song to database: %+v", song)

	query := "INSERT INTO songs (group_name, song_name, text, release_date, link) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at"
	err := r.DB.QueryRow(query, song.GroupName, song.SongName, song.Text, song.ReleaseDate, song.Link).Scan(&song.ID, &song.CreatedAt)
	if err != nil {
		log.Printf("[ERROR] Failed to save song: %v", err)
		return err
	}

	log.Printf("[DEBUG] Song saved with ID: %s", song.ID)
	return nil
}

// DeleteBySongNameAndGroup удаляет песню по её названию
func (r *SongRepository) DeleteBySongNameAndGroup(songName, group string) error {
	log.Printf("[INFO] Deleting song with name: %s", songName)

	query := "DELETE FROM songs WHERE song_name = $1 AND group_name = $2"
	result, err := r.DB.Exec(query, songName, group)
	if err != nil {
		log.Printf("[ERROR] Failed to delete song: %v", err)
		return err
	}

	// Проверяем количество удалённых записей
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[ERROR] Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("[INFO] No song found with name: %s", songName)
		return sql.ErrNoRows
	}

	log.Printf("[INFO] Song deleted successfully: %s", songName)
	return nil
}

// UpdateSong изменяет данные песни
func (r *SongRepository) UpdateSong(req models.UpdateSongRequest) error {
	log.Printf("[INFO] Updating song: oldName=%s, oldGroup=%s, newGroup=%s, newName=%s", req.OldSongName, req.OldGroup, req.NewGroup, req.NewSongName)

	query := `
		UPDATE songs
		SET group_name = $1, song_name = $2, text = $3, release_date = $6, link = $7
		WHERE song_name = $4 AND group_name = $5
	`
	result, err := r.DB.Exec(query, req.NewGroup, req.NewSongName, req.NewText, req.OldSongName, req.OldGroup, req.NewReleaseDate, req.NewLink)
	if err != nil {
		log.Printf("[ERROR] Failed to update song: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[ERROR] Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("[INFO] No song found with name: %s", req.OldSongName)
		return sql.ErrNoRows
	}

	log.Printf("[INFO] Song updated successfully: oldName=%s, oldGroup=%s", req.OldSongName, req.OldGroup)
	return nil
}

// FindSongs фильтрует и возвращает песни с учетом параметров пагинации
func (r *SongRepository) FindSongs(params models.FilterParams) ([]*models.Song, error) {
	log.Printf("[INFO] Fetching songs with filters: %+v", params)

	query := `
		SELECT id, group_name, song_name, text, created_at, release_date, link
		FROM songs
		WHERE ($1 = '' OR group_name ILIKE '%' || $1 || '%')
		  AND ($2 = '' OR song_name ILIKE '%' || $2 || '%')
		  AND ($3 = '' OR text ILIKE '%' || $3 || '%')
		  AND ($4 = '' OR release_date < $4)
		LIMIT $5 OFFSET $6
	`

	rows, err := r.DB.Query(query, params.Group, params.SongName, params.Text, params.ReleaseDate, params.Limit, params.Offset)
	if err != nil {
		log.Printf("[ERROR] Failed to execute query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var songs []*models.Song
	for rows.Next() {
		song := &models.Song{}
		if err := rows.Scan(&song.ID, &song.GroupName, &song.SongName, &song.Text, &song.CreatedAt, &song.ReleaseDate, &song.Link); err != nil {
			log.Printf("[ERROR] Failed to scan row: %v", err)
			return nil, err
		}
		songs = append(songs, song)
	}

	log.Printf("[INFO] Found %d songs", len(songs))
	return songs, nil
}

// GetSongTextByNameAndGroup получает текст песни по названию
func (r *SongRepository) GetSongTextByNameAndGroup(songName, group string) (string, error) {
	log.Printf("[INFO] Fetching song text for song: %s, group: %s", songName, group)

	query := "SELECT text FROM songs WHERE song_name = $1 AND group_name = $2"
	var text string
	err := r.DB.QueryRow(query, songName, group).Scan(&text)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch song text: %v", err)
		return "", err
	}

	log.Printf("[INFO] Successfully fetched text for song: %s", songName)
	return text, nil
}

// GetSongInfo получает информацию о песне по имени группы и названию песни
func (r *SongRepository) GetSongInfo(group, songName string) (*models.SongDetail, error) {
	log.Printf("[INFO] Fetching song info for group: %s, song: %s", group, songName)

	query := `
		SELECT release_date, text, link
		FROM songs
		WHERE group_name = $1 AND song_name = $2
	`

	var songDetail models.SongDetail
	err := r.DB.QueryRow(query, group, songName).Scan(&songDetail.ReleaseDate, &songDetail.Text, &songDetail.Link)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch song info: %v", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully fetched song info: %s", songName)
	return &songDetail, nil
}
