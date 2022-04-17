package logic

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ImTheTom/OtherProjects/race-track/pkg/horse"
	"github.com/ImTheTom/OtherProjects/race-track/pkg/jockey"
)

const (
	horsesLocation = "./pkg/horse/horses.json"
	jockeyLocation = "./pkg/jockey/jockey.json"
)

var (
	horsesVar  = horse.GetLoadedHorsesVar()
	jockeysVar = jockey.GetLoadedJockeysVar()
)

func LoadInData() error {
	if err := loadInDataSpecified(
		horsesLocation,
		horsesVar,
	); err != nil {
		return err
	}

	if err := loadInDataSpecified(
		jockeyLocation,
		jockeysVar,
	); err != nil {
		return err
	}

	return nil
}

func loadInDataSpecified(fileLocation string, loadVar interface{}) error {
	jsonFile, err := os.Open(fileLocation)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, &loadVar)
}
