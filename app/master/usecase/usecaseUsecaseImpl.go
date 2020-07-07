package usecase

import (
	"enigmamusic/app/master/models"
	"enigmamusic/app/master/repositories"
	"log"
)

type SongUsecaseImpl struct {
	songRepo repositories.SongRepositories
}

func (s SongUsecaseImpl) GetSongs() ([]*models.Song, error) {
	song, err := s.songRepo.GetAllSong()
	if err != nil {
		return nil, err
	}

	return song, nil
}
func (s SongUsecaseImpl) GetGenre(genremusik string) ([]*models.Song, error) {
	genre, err := s.songRepo.GetGenre(genremusik)
	if err != nil {
		log.Fatal(err)
	}
	return genre, nil
}
func (s SongUsecaseImpl) AddDataSong(song models.Song) error {
	err := s.songRepo.AddSong(song)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s SongUsecaseImpl) UpdateDataSong(Song models.Song, id string) (models.Song, error) {
	song, err := s.songRepo.UpdateSong(Song, id)
	if err != nil {
		log.Fatal(err)
	}
	return song, nil
}

func (s SongUsecaseImpl) DeleteDataSong(id string) error {
	err := s.songRepo.DeleteSong(id)
	if err != nil {
		return err
	}

	return nil
}

func InitSongUseCase(songRepo repositories.SongRepositories) SongUseCase {
	return &SongUsecaseImpl{songRepo}
}
