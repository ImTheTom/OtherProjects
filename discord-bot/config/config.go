package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	BotToken           string `json:"botToken"`
	DatabaseConnection string `json:"databaseConnection"`
	Prefix             string `json:"prefix"`
}

const EnvConfigName = "CONFIG_LOCATION"

var (
	_config       = SetupConfig()
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

func SetupConfig() Config {
	defaultConfig := Config{}

	configFile := os.Getenv(EnvConfigName)
	if len(configFile) == 0 {
		logrus.Warn(errNoConfig)

		return defaultConfig
	}

	raw, err := ioutil.ReadFile(configFile)
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
	if len(c.BotToken) == 0 {
		return errBadConfig
	}

	if len(c.DatabaseConnection) == 0 {
		return errBadConfig
	}

	if len(c.Prefix) == 0 {
		return errBadConfig
	}

	return nil
}
