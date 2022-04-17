package jockey

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Jockey struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type Jockeys struct {
	Jockeys []Jockey `json:"jockeys"`
}

const jockeyLocation = "../pkg/jockey/jockey.json"

var loadedJockeys = loadInJockeys(jockeyLocation)

func GetLoadedJockeysVar() Jockeys {
	return loadedJockeys
}

func GetRandomJockey() *Jockey {
	return &loadedJockeys.Jockeys[rand.Intn(GetTotalJockeysLoaded())]
}

func GetTotalJockeysLoaded() int {
	return len(loadedJockeys.Jockeys)
}

func loadInJockeys(fileLocation string) Jockeys {
	jsonFile, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var jockeys Jockeys

	err = json.Unmarshal(byteValue, &jockeys)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	return jockeys
}
