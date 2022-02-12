package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/stretchr/testify/assert"
)

func setUp() {
	blankConfig := config.Config{
		BotToken:           "",
		DatabaseConnection: "",
		Prefix:             "",
	}

	config.SetConfig(&blankConfig)
}

func TestInit(t *testing.T) {
	tests := []struct {
		name         string
		fileLocation string
		exceptError  bool
	}{
		{
			name:         "Good config returns nil",
			fileLocation: "../test/config/good.json",
			exceptError:  false,
		},
		{
			name:         "Empty discord token returns an error",
			fileLocation: "../test/config/bad.json",
			exceptError:  true,
		},
		{
			name:         "Missing database url returns an error",
			fileLocation: "../test/config/bad-2.json",
			exceptError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUp()

			os.Setenv(config.EnvConfigName, tt.fileLocation)
			result := config.Init()
			fmt.Printf("error %v\n", result)
			if tt.exceptError {
				assert.NotNil(t, result)
			} else {
				assert.Nil(t, result)
			}
		})
	}
}
