package usecase

import "enigmamusic/app/master/models"

type ArtistUseCase interface {
	GetArtists() ([]*models.Artist, error)
	AddDataArtist(models.Artist) error
	UpdateDataArtist(artist models.Artist, id string) (models.Artist, error)
	DeleteDataArtist(id string) error
}
