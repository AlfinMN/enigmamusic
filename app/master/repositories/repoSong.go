package repositories

import "enigmamusic/app/master/models"

type SongRepositories interface {
	GetAllSong() ([]*models.Song, error)
	GetGenre(id string) ([]*models.Song, error)
	AddSong(models.Song) error
	UpdateSong(Song models.Song, id string) (models.Song, error)
	DeleteSong(id string) error
}
