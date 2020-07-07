package usecase

import (
	"enigmamusic/app/master/models"
	"enigmamusic/app/master/repositories"
	"log"
)

type ArtistUsecaseImpl struct {
	artistRepo repositories.ArtistRepositories
}

func (a ArtistUsecaseImpl) GetArtists() ([]*models.Artist, error) {
	artist, err := a.artistRepo.GetAllArtist()
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (a ArtistUsecaseImpl) AddDataArtist(artist models.Artist) error {
	err := a.artistRepo.AddArtist(artist)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (a ArtistUsecaseImpl) UpdateDataArtist(artist models.Artist, id string) (models.Artist, error) {
	artist, err := a.artistRepo.UpdateArtist(artist, id)
	if err != nil {
		log.Fatal(err)
	}
	return artist, nil
}

func (a ArtistUsecaseImpl) DeleteDataArtist(id string) error {
	err := a.artistRepo.DeleteArtist(id)
	if err != nil {
		return err
	}

	return nil
}

func InitArtistUseCase(artistRepo repositories.ArtistRepositories) ArtistUseCase {
	return &ArtistUsecaseImpl{artistRepo}
}
