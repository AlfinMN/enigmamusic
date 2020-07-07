package controller

import (
	"encoding/json"
	"enigmamusic/app/master/models"
	"enigmamusic/app/master/usecase"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SongHandler struct {
	SongUseCase usecase.SongUseCase
}

func SongController(r *mux.Router, service usecase.SongUseCase) {
	SongHandler := SongHandler{service}

	r.HandleFunc("/songs", SongHandler.ListSongs).Methods(http.MethodGet)
	r.HandleFunc("/song", SongHandler.CreateDataSong).Methods(http.MethodPost)
	r.HandleFunc("/song/{id}", SongHandler.UpToDateDataSong).Methods(http.MethodPut)
	r.HandleFunc("/song/{id}", SongHandler.DeleteContentSong).Methods(http.MethodDelete)
	r.HandleFunc("/song/{genre}", SongHandler.FindGenre).Methods(http.MethodGet)
}

func (s SongHandler) ListSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := s.SongUseCase.GetSongs()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "data Song's succes displayed"
	status.Data = songs

	byteOfSongs, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfSongs)
}
func (s SongHandler) CreateDataSong(w http.ResponseWriter, r *http.Request) {

	var song models.Song
	saveData, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(saveData, &song)
	err := s.SongUseCase.AddDataSong(song)

	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "Create data song succes ! "
	status.Data = song

	byteOfSongs, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfSongs)
}

func (s SongHandler) UpToDateDataSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	ID := mux.Vars(r)
	SongId, _ := (ID["id"])
	saveData, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(saveData, &song)
	dataSong, err := s.SongUseCase.UpdateDataSong(song, SongId)

	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "Update data song succes ! "
	status.Data = dataSong

	byteOfSong, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfSong)
}

func (s SongHandler) DeleteContentSong(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)
	err := s.SongUseCase.DeleteDataSong(ID["id"])
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte("Delete Data success"))
}

func (s SongHandler) FindGenre(w http.ResponseWriter, r *http.Request) {
	Genre := mux.Vars(r)
	GenreMusik, _ := (Genre["genre"])
	GenreSong, err := s.SongUseCase.GetGenre(GenreMusik)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "data musik by genre  "
	status.Data = GenreSong

	byteOfGenre, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfGenre)
}
