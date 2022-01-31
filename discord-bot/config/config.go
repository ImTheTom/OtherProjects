package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	BotToken           string `json:"botToken"`
	DatabaseConnection string `json:"databaseConnection"`
	Prefix             string `json:"prefix"`
}

const EnvConfigName = "CONFIG_LOCATION"

var (
	config       Config
	errBadConfig = errors.New("Bad config was loaded in")
	errNoConfig  = errors.New("Can't find config file")
)

func GetConfig() Config {
	return config
}

func SetConfig(cfg Config) {
	config = cfg
}

func Init() error {
	configFile := os.Getenv(EnvConfigName)
	if len(configFile) == 0 {
		return errNoConfig
	}

	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		return err
	}

	return sanityCheckValues()
}

func sanityCheckValues() error {
	if len(config.BotToken) == 0 {
		return errBadConfig
	}

	if len(config.DatabaseConnection) == 0 {
		return errBadConfig
	}

	if len(config.Prefix) == 0 {
		return errBadConfig
	}

	return nil
}
