package main

import (
	"OtherProjects/ff3-easy/config"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	if !config.IsConfigSetup() {
		return fmt.Errorf("Invalid config setup")
	}

	fmt.Println("Welcome to FF3-Easy")

	return nil
}
