package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Location  string `json:"location"`
	AuthToken string `json:"auth_token"`
}

const ConfigLocation = "./config/config.json"

var (
	_config, _    = SetupConfig(ConfigLocation)
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

func SetupConfig(location string) (Config, error) {
	defaultConfig := Config{}

	raw, err := ioutil.ReadFile(location)
	if err != nil {
		fmt.Println(err)
		return defaultConfig, err
	}

	if err = json.Unmarshal(raw, &defaultConfig); err != nil {
		fmt.Println(err)
		return defaultConfig, err
	}

	err = defaultConfig.sanityCheckValues()

	if err == nil {
		isConfigSetup = true
	}

	return defaultConfig, err
}

func (c *Config) sanityCheckValues() error {
	if len(c.Location) == 0 {
		return errBadConfig
	}

	if len(c.AuthToken) == 0 {
		return errBadConfig
	}

	return nil
}
