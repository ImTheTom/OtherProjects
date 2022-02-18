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

var _config Config

var (
	errBadConfig = errors.New("Bad config was loaded in")
	errNoConfig  = errors.New("Can't find config file")
)

func GetConfig() *Config {
	return &_config
}

func SetConfig(cfg *Config) {
	_config = *cfg
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

	if err = json.Unmarshal(raw, &_config); err != nil {
		return err
	}

	return _config.sanityCheckValues()
}

func (c *Config) sanityCheckValues() error {
	if len(_config.BotToken) == 0 {
		return errBadConfig
	}

	if len(_config.DatabaseConnection) == 0 {
		return errBadConfig
	}

	if len(_config.Prefix) == 0 {
		return errBadConfig
	}

	return nil
}
