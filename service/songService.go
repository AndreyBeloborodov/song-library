package service

import (
	"song-libary/models"
	"song-libary/repository"
)

type SongService struct {
	Repo *repository.SongRepository
}

// NewSongService создает новый экземпляр сервиса
func NewSongService(repo *repository.SongRepository) *SongService {
	return &SongService{Repo: repo}
}

// AddSong добавляет новую песню
func (s *SongService) AddSong(group, song, text string) (*models.Song, error) {
	newSong := &models.Song{
		GroupName: group,
		SongName:  song,
		Text:      text,
	}

	// Сохраняем песню в базе данных через репозиторий
	if err := s.Repo.Save(newSong); err != nil {
		return nil, err
	}
	return newSong, nil
}
