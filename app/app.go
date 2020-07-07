package main

import (
	"enigmamusic/app/master"
	"enigmamusic/config"
)

func main() {

	db, _, dbHost, portServer := config.Connection()
	router := config.CreateRouter()

	master.Init(router, db)
	config.RunServer(router, dbHost, portServer)

}
