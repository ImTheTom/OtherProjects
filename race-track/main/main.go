package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/logic"
)

func main() {
	rand.Seed(time.Now().Unix())

	if err := logic.LoadInData(); err != nil {
		log.Fatal(err)
	}
}
