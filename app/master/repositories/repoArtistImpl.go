package repositories

import (
	"database/sql"
	"enigmamusic/app/master/models"
	"log"
)

type ArtistRepoImpl struct {
	db *sql.DB
}

func (a ArtistRepoImpl) GetAllArtist() ([]*models.Artist, error) {
	dataArtist := []*models.Artist{}
	query := "select * from artist"
	data, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		artist := models.Artist{}
		var err = data.Scan(&artist.Id, &artist.Name)
		if err != nil {
			return nil, err
		}
		dataArtist = append(dataArtist, &artist)
	}
	return dataArtist, nil
}

func (a ArtistRepoImpl) AddArtist(artist models.Artist) error {

	tx, err := a.db.Begin()
	if err != nil {
	}
	_, err = tx.Exec("insert into artist values(?,?)", artist.Id, artist.Name)
	if err != nil {
		tx.Rollback()
		// log.Fatal(err)
	}
	err = tx.Commit()

	return err
}

func (a ArtistRepoImpl) UpdateArtist(artist models.Artist, id string) (models.Artist, error) {

	tx, err := a.db.Begin()
	if err != nil {
		log.Print(err)
	}
	_, err = tx.Exec("UPDATE artist SET id_artist=?,nama_artist=? where id_artist=?", artist.Id, artist.Name, id)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	err = tx.Commit()

	return artist, err
}

func (a *ArtistRepoImpl) DeleteArtist(id string) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM artist WHERE id_artist = ?")
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

func InitArtistRepoImpl(db *sql.DB) ArtistRepositories {
	return &ArtistRepoImpl{db}
}
