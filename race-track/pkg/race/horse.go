package race

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Horse struct {
	Name  string  `json:"name"`
	Speed float64 `json:"speed"`
}

type Horses struct {
	Horses []Horse `json:"horses"`
}

const horsesLocation = "../pkg/race/horses.json"

var loadedHorses = loadInHorses(horsesLocation)

func GetLoadedHorsesVar() Horses {
	return loadedHorses
}

func GetRandomHorse() *Horse {
	return &loadedHorses.Horses[rand.Intn(GetTotalHorsesLoaded())]
}

func GetTotalHorsesLoaded() int {
	return len(loadedHorses.Horses)
}

func loadInHorses(fileLocation string) Horses {
	jsonFile, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var horses Horses

	err = json.Unmarshal(byteValue, &horses)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	return horses
}
