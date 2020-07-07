package models

type Song struct {
	Id          string `json:"id_lagu"`
	Title       string `json:"judul_lagu"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Genre       string `json:"genre"`
	ReleaseYear string `json:"tahun_rilis"`
}
