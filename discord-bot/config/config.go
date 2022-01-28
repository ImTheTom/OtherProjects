package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	BotToken           string `json:"botToken"`
	DatabaseConnection string `json:"databaseConnection"`
	Prefix             string `json:"prefix"`
	GuildID            string `json:"guild_id"`
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

	if len(config.DatabaseConnection) == 0 {
		return errBadConfig
	}

	if len(config.Prefix) == 0 {
		return errBadConfig
	}

	if len(config.GuildID) == 0 {
		return errBadConfig
	}

	return nil
}
