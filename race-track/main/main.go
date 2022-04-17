package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/logic"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/race"
)

func main() {
	rand.Seed(time.Now().Unix())

	if err := logic.LoadInData(); err != nil {
		log.Fatal(err)
	}

	rc, err := race.CreateRace(1, "new")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rc)
}
