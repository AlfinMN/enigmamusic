package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
func RunServer(router *mux.Router, dbHost string, portServer string) {

	fmt.Println("succes connect to port : " + portServer)
	err := http.ListenAndServe(dbHost+":"+portServer, router)
	if err != nil {
		log.Fatal(err)
	}
}
