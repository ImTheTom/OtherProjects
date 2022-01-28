package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	BotToken string `json:"botToken"`
}

const configFile = "./config/config.json"

var (
	config       Config
	errBadConfig = errors.New("Bad config was loaded in")
)

func Init() error {
	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		return err
	}

	return sanityCheckValues()
}

func GetConfig() Config {
	return config
}

func sanityCheckValues() error {
	if len(config.BotToken) == 0 {
		return errBadConfig
	}

	return nil
}
