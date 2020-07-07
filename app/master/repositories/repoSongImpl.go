package repositories

import (
	"database/sql"
	"enigmamusic/app/master/models"
	"log"
)

type SongRepoImpl struct {
	db *sql.DB
}

func (s SongRepoImpl) GetAllSong() ([]*models.Song, error) {
	dataSong := []*models.Song{}
	query := "select * from lagu"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		song := models.Song{}
		var err = data.Scan(&song.Id, &song.Title, &song.Artist, &song.Album, &song.Genre, &song.ReleaseYear)
		if err != nil {
			return nil, err
		}
		dataSong = append(dataSong, &song)
	}
	return dataSong, nil
}

func (s SongRepoImpl) AddSong(song models.Song) error {

	tx, err := s.db.Begin()
	if err != nil {
	}
	_, err = tx.Exec("insert into lagu values(?,?,?,?,?,?)", song.Id, song.Title, song.Artist, song.Album, song.Genre, song.ReleaseYear)
	if err != nil {
		tx.Rollback()
		// log.Fatal(err)
	}
	err = tx.Commit()

	return err
}

func (s SongRepoImpl) UpdateSong(song models.Song, id string) (models.Song, error) {

	tx, err := s.db.Begin()
	if err != nil {
		log.Print(err)
	}
	_, err = tx.Exec("UPDATE lagu SET id_lagu=?,judul_lagu=?,artist=?,album=?,genre=?,tahun_rilis=? where id_lagu=?", song.Id, song.Title, song.Artist, song.Album, song.Genre, song.ReleaseYear, id)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	err = tx.Commit()

	return song, err
}

func (s *SongRepoImpl) DeleteSong(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM lagu WHERE id_lagu = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (a SongRepoImpl) GetGenre(genremusik string) (models.Song, error) {
	var genre models.Song
	query := "SELECT * FROM lagu WHERE genre=? group by id"
	err := a.db.QueryRow(query, genremusik).Scan(&genre.Id, &genre.Title, &genre.Artist, &genre.Album, &genre.Genre, &genre.ReleaseYear)

	if err != nil {
		log.Fatal(err)
	}

	return genre, nil
}

func InitSongRepoImpl(db *sql.DB) SongRepositories {
	return &SongRepoImpl{db}
}
