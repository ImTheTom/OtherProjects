package main

import (
	"log"
	"tracker/db"
	"tracker/server"
	"tracker/service"
)

func main() {
	log.Println("Starting server")

	dbRepo, err := db.NewTracker()
	if err != nil {
		log.Fatal(err)
	}

	srv := service.NewTracker(dbRepo)

	serv := server.NewTracker(srv)

	serv.Run()
}
