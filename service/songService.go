package service

import (
	"database/sql"
	"errors"
	"log"
	"song-libary/models"
	"song-libary/repository"
	"strings"
)

var ErrSongNotFound = errors.New("song not found")

type SongService struct {
	Repo *repository.SongRepository
}

func NewSongService(repo *repository.SongRepository) *SongService {
	return &SongService{Repo: repo}
}

func (s *SongService) AddSong(group, song, text, releaseDate, link string) (*models.Song, error) {
	log.Printf("[INFO] Adding new song: group=%s, song=%s", group, song)

	newSong := &models.Song{
		GroupName:   group,
		SongName:    song,
		Text:        text,
		ReleaseDate: releaseDate,
		Link:        link,
	}

	if err := s.Repo.SaveSong(newSong); err != nil {
		log.Printf("[ERROR] Failed to add song: %v", err)
		return nil, err
	}

	log.Printf("[INFO] Song added successfully: %+v", newSong)
	return newSong, nil
}

// DeleteSongByNameAndGroup удаляет песню по названию
func (s *SongService) DeleteSongByNameAndGroup(songName, group string) error {
	log.Printf("[INFO] Deleting song with name: %s", songName)

	if err := s.Repo.DeleteBySongNameAndGroup(songName, group); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[INFO] Song not found: %s", songName)
			return ErrSongNotFound
		}
		log.Printf("[ERROR] Failed to delete song: %v", err)
		return err
	}

	log.Printf("[INFO] Song deleted successfully: %s", songName)
	return nil
}

// UpdateSong изменяет данные песни
func (s *SongService) UpdateSong(req models.UpdateSongRequest) error {
	log.Printf("[INFO] Updating song: oldName=%s, oldGroup=%s, newGroup=%s, newName=%s", req.OldSongName, req.OldGroup, req.NewGroup, req.NewSongName)

	if err := s.Repo.UpdateSong(req); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[INFO] Song not found for update: %s", req.OldSongName)
			return ErrSongNotFound
		}
		log.Printf("[ERROR] Failed to update song: %v", err)
		return err
	}

	log.Printf("[INFO] Song updated successfully: oldName=%s, oldGroup=%s", req.OldSongName, req.OldGroup)
	return nil
}

// GetSongs возвращает песни с учетом фильтров и пагинации
func (s *SongService) GetSongs(params models.FilterParams) ([]*models.Song, error) {
	log.Printf("[INFO] Fetching songs with params: %+v", params)
	return s.Repo.FindSongs(params)
}

// GetSongText возвращает текст песни с учетом пагинации
func (s *SongService) GetSongText(songName, group string, limit, offset int) (models.SongTextResponse, error) {
	log.Printf("[INFO] Fetching text for song: %s with pagination: limit=%d, offset=%d", songName, limit, offset)

	text, err := s.Repo.GetSongTextByNameAndGroup(songName, group)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch song text: %v", err)
		return models.SongTextResponse{}, ErrSongNotFound
	}

	// Разделяем текст на куплеты
	verses := strings.Split(text, "\\n\\n")
	total := len(verses)

	// Применяем пагинацию
	start := offset
	if start >= total {
		start = total
	}
	end := start + limit
	if end > total {
		end = total
	}

	paginatedVerses := verses[start:end]

	response := models.SongTextResponse{
		SongName: songName,
		Group:    group,
		Verses:   paginatedVerses,
		Limit:    limit,
		Offset:   offset,
		Total:    total,
	}

	log.Printf("[INFO] Returning %d verses for song: %s", len(paginatedVerses), songName)
	return response, nil
}

// GetSongInfo получает информацию о песне по группе и названию
func (s *SongService) GetSongInfo(group, songName string) (*models.SongDetail, error) {
	log.Printf("[INFO] Fetching song info for group: %s, song: %s", group, songName)
	return s.Repo.GetSongInfo(group, songName)
}
