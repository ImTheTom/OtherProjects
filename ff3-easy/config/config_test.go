package config_test

import (
	"OtherProjects/ff3-easy/config"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp() {
	blankConfig := config.Config{
		Location:  "",
		AuthToken: "",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setUp()

			result, err := config.SetupConfig(tt.fileLocation)
			fmt.Printf("error %v\n", result)
			if tt.exceptError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
