package race

import (
	"fmt"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/entrant"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/horse"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/jockey"
)

type Race struct {
	Name     string
	Entrants []entrant.Entrant
}

func CreateRace(numParticipants int, name string) (*Race, error) {
	if numParticipants > horse.GetTotalHorsesLoaded() || numParticipants > jockey.GetTotalJockeysLoaded() {
		return nil, fmt.Errorf("Not enough horses or jockeys")
	}

	entrants := make([]entrant.Entrant, numParticipants)
	for i := 0; i < numParticipants; i++ {
		horseFound := false
		var hrs horse.Horse
		var jky jockey.Jockey
		for !horseFound {
			hrs = horse.GetRandomHorse()

		}
	}
}
