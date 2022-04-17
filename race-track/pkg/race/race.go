package race

import (
	"errors"
)

type Race struct {
	Name     string
	Entrants []*Entrant
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
	if numParticipants > GetTotalHorsesLoaded() || numParticipants > GetTotalJockeysLoaded() {
		return nil, errNotEnough
	}

	entrants := make([]*Entrant, numParticipants)

	for i := 0; i < numParticipants; i++ {
		hrs := selectHorse(entrants)
		jky := selectJockey(entrants)

		ent := Entrant{
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

func selectHorse(ents []*Entrant) *Horse {
	horseFound := false

	var hrs *Horse

	for !horseFound {
		hrs = GetRandomHorse()
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

func selectJockey(ents []*Entrant) *Jockey {
	jockeyFound := false

	var jky *Jockey

	for !jockeyFound {
		jky = GetRandomJockey()
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
