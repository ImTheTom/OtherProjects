package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type Config struct {
	DatabaseConnection string `json:"databaseConnection"`
}

var (
	_config       = SetupConfig("config.json")
	isConfigSetup = false
)

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

func IsConfigSetup() bool {
	return isConfigSetup
}

func SetupConfig(location string) Config {
	defaultConfig := Config{}

	raw, err := ioutil.ReadFile(location)
	if err != nil {
		logrus.Warn(err)

		return defaultConfig
	}

	if err = json.Unmarshal(raw, &defaultConfig); err != nil {
		logrus.Warn(err)

		return defaultConfig
	}

	if err = defaultConfig.sanityCheckValues(); err != nil {
		isConfigSetup = false
	} else {
		isConfigSetup = true
	}

	return defaultConfig
}

func (c *Config) sanityCheckValues() error {
	if len(c.DatabaseConnection) == 0 {
		return errBadConfig
	}

	return nil
}
