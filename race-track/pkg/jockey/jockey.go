package jockey

import "math/rand"

type Jockey struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type Jockeys struct {
	Jockeys []Jockey `json:"jockeys"`
}

var loadedJockeys Jockeys

func GetLoadedJockeysVar() Jockeys {
	return loadedJockeys
}

func GetRandomJockey() Jockey {
	return loadedJockeys.Jockeys[rand.Intn(GetTotalJockeysLoaded())]
}

func GetTotalJockeysLoaded() int {
	return len(loadedJockeys.Jockeys)
}
