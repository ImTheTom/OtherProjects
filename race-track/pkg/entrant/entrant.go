package entrant

import (
	"github.com/ImTheTom/OtherProjects/race-track/pkg/horse"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/jockey"
)

type Entrant struct {
	Horse  horse.Horse
	Jockey jockey.Jockey
	Number int
}
