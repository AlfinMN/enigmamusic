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

type ArtistHandler struct {
	ArtistUseCase usecase.ArtistUseCase
}

func ArtistController(r *mux.Router, service usecase.ArtistUseCase) {
	ArtistHandler := ArtistHandler{service}

	r.HandleFunc("/artists", ArtistHandler.ListArtists).Methods(http.MethodGet)
	r.HandleFunc("/artist", ArtistHandler.CreateDataArtist).Methods(http.MethodPost)
	r.HandleFunc("/artist/{id}", ArtistHandler.UpToDateDataArtist).Methods(http.MethodPut)
	r.HandleFunc("/artist/{id}", ArtistHandler.DeleteContentArtist).Methods(http.MethodDelete)
}

func (a ArtistHandler) ListArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := a.ArtistUseCase.GetArtists()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "data artist's succes displayed"
	status.Data = artists

	byteOfArtists, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfArtists)
}
func (a ArtistHandler) CreateDataArtist(w http.ResponseWriter, r *http.Request) {

	var artist models.Artist
	saveData, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(saveData, &artist)
	err := a.ArtistUseCase.AddDataArtist(artist)

	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "Create data artist succes ! "
	status.Data = artist

	byteOfArtists, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfArtists)
}

func (a ArtistHandler) UpToDateDataArtist(w http.ResponseWriter, r *http.Request) {
	var artist models.Artist
	ID := mux.Vars(r)
	ArtistId, _ := (ID["id"])
	saveData, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(saveData, &artist)
	dataArtist, err := a.ArtistUseCase.UpdateDataArtist(artist, ArtistId)

	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "Update data artist succes ! "
	status.Data = dataArtist

	byteOfArtist, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfArtist)
}

func (a ArtistHandler) DeleteContentArtist(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)
	err := a.ArtistUseCase.DeleteDataArtist(ID["id"])
	if err != nil {
		log.Println(err)
	}

	w.Write([]byte("Delete Data artist success"))
}
