package master

import (
	"database/sql"
	"enigmamusic/app/master/controller"
	"enigmamusic/app/master/repositories"
	"enigmamusic/app/master/usecase"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	songRepo := repositories.InitSongRepoImpl(db)
	songUseCase := usecase.InitSongUseCase(songRepo)
	controller.SongController(r, songUseCase)

	artistRepo := repositories.InitArtistRepoImpl(db)
	artistUseCase := usecase.InitArtistUseCase(artistRepo)
	controller.ArtistController(r, artistUseCase)
}
