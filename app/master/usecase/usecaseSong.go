package usecase

import "enigmamusic/app/master/models"

type SongUseCase interface {
	GetSongs() ([]*models.Song, error)
	AddDataSong(models.Song) error
	UpdateDataSong(Song models.Song, id string) (models.Song, error)
	DeleteDataSong(id string) error
	GetGenre(id string) (models.Song, error)
}
