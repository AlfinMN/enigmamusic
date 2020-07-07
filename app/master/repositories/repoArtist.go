package repositories

import "enigmamusic/app/master/models"

type ArtistRepositories interface {
	GetAllArtist() ([]*models.Artist, error)
	AddArtist(models.Artist) error
	UpdateArtist(artist models.Artist, id string) (models.Artist, error)
	DeleteArtist(id string) error
}
