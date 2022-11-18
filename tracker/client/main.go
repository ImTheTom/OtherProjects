package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type CreateTokenRequest struct {
	Urge       int       `json:"urge"`
	Need       int       `json:"need"`
	CreateTime time.Time `json:"create_time"`
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("expected 2 args")
		return
	}

	urge, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	need, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	loc, err := time.LoadLocation("Australia/Brisbane")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = CreateTokenRequest{
		Urge:       urge,
		Need:       need,
		CreateTime: time.Now().In(loc),
	}

	// Send the args
}
