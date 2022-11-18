package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	"tracker/models"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("expected 2 args")
		return
	}

	req, err := createRequest(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = sendRequest(req)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Was accepted")
	}
}

func createRequest(args []string) (*models.CreateTrackRequest, error) {
	urge, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, err
	}

	need, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation("Australia/Brisbane")
	if err != nil {
		return nil, err
	}

	return &models.CreateTrackRequest{
		Urge:       urge,
		Need:       need,
		CreateTime: time.Now().In(loc),
	}, nil
}

func sendRequest(createReq *models.CreateTrackRequest) error {
	marshedReq, err := json.Marshal(createReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/create_track", bytes.NewBuffer(marshedReq))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		var errResponse models.CreateTrackErrorResponse

		err := json.Unmarshal(body, &errResponse)
		if err != nil {
			return err
		}

		return errors.New(errResponse.Error)
	}

	return nil
}
