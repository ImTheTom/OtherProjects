package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/race"
)

func main() {
	rand.Seed(time.Now().Unix())

	rc, err := race.CreateRace(race.TwoHorseRace, "new", race.TwelveHundredMetre)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rc)

	for !rc.Finished {
		rc.StepRace()
	}
}
