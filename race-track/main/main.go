package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/race"
)

func main() {
	rand.Seed(time.Now().Unix())

	rc, err := race.CreateRace(1, "new")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rc)
}
