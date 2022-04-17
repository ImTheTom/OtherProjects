package horse

import "math/rand"

type Horse struct {
	Name  string  `json:"name"`
	Speed float64 `json:"speed"`
}

type Horses struct {
	Horses []Horse `json:"horses"`
}

var loadedHorses Horses

func GetLoadedHorsesVar() Horses {
	return loadedHorses
}

func GetRandomHorse() Horse {
	return loadedHorses.Horses[rand.Intn(GetTotalHorsesLoaded())]
}

func GetTotalHorsesLoaded() int {
	return len(loadedHorses.Horses)
}
