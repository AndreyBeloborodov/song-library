package repository

import "song-libary/models"

type SongRepository interface {
	SaveSong(song *models.Song) error
	DeleteBySongNameAndGroup(songName, group string) error
	UpdateSong(req models.UpdateSongRequest) error
	FindSongs(params models.FilterParams) ([]*models.Song, error)
	GetSongTextByNameAndGroup(songName, group string) (string, error)
	GetSongInfo(group, songName string) (*models.SongDetail, error)
}
