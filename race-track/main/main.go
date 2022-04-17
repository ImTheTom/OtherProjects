package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/gamble"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/race"
)

func main() {
	rand.Seed(time.Now().Unix())

	usr := gamble.CreateNewUser("Tom")

	rc, err := race.CreateRace(race.TwoHorseRace, "new", race.TwelveHundredMetre)
	if err != nil {
		log.Fatal(err)
	}

	gmb := gamble.CreateNewGamble(
		usr.ID,
		rc.ID,
		rc.Entrants[0].Number,
		10.0,
		2.0,
	)

	log.Println(rc)

	for !rc.Finished {
		rc.StepRace()
	}

	gamble.PayoutWinnings(rc)

	log.Println(usr)

	log.Println(gmb)
}
