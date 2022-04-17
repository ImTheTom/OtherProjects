package race

import (
	"errors"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/entrant"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/horse"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/jockey"
)

type Race struct {
	Name     string
	Entrants []*entrant.Entrant
	Distance float64
	Step     int
	Finished bool
}

const (
	TwelveHundredMetre    float64 = 1200
	FifthteenHundredMetre float64 = 1500
)

const (
	TwoHorseRace = 2
)

var errNotEnough = errors.New("Not enough horses or jockeys")

func CreateRace(numParticipants int, name string, distance float64) (*Race, error) {
	if numParticipants > horse.GetTotalHorsesLoaded() || numParticipants > jockey.GetTotalJockeysLoaded() {
		return nil, errNotEnough
	}

	entrants := make([]*entrant.Entrant, numParticipants)

	for i := 0; i < numParticipants; i++ {
		hrs := selectHorse(entrants)
		jky := selectJockey(entrants)

		ent := entrant.Entrant{
			Horse:     hrs,
			Jockey:    jky,
			Number:    i + 1,
			Position:  i + 1,
			Travelled: 0,
		}

		entrants[i] = &ent
	}

	rc := &Race{
		Name:     name,
		Entrants: entrants,
		Distance: distance,
		Step:     0,
	}

	return rc, nil
}

func selectHorse(ents []*entrant.Entrant) *horse.Horse {
	horseFound := false

	var hrs *horse.Horse

	for !horseFound {
		hrs = horse.GetRandomHorse()
		horseFound = true

		// Make sure it's unique
		for _, v := range ents {
			if v == nil {
				continue
			}

			if v.Horse == hrs {
				horseFound = false
			}
		}
	}

	return hrs
}

func selectJockey(ents []*entrant.Entrant) *jockey.Jockey {
	jockeyFound := false

	var jky *jockey.Jockey

	for !jockeyFound {
		jky = jockey.GetRandomJockey()
		jockeyFound = true

		// Make sure it's unique
		for _, v := range ents {
			if v == nil {
				continue
			}

			if v.Jockey == jky {
				jockeyFound = false
			}
		}
	}

	return jky
}