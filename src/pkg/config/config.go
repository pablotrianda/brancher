package config

import (
	"errors"
	"os"

	"github.com/pablotrianda/brancher/src/pkg/constans"
	"github.com/pablotrianda/brancher/src/pkg/db"
	"github.com/pablotrianda/brancher/src/pkg/prompt"
)

func findConfiguration(string) (confPath string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := homeDir + constans.CONFIG_DIR
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return "", errors.New("Configuration not found")
	}
	return configDir, nil
}

func CreateNewConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := homeDir + constans.CONFIG_DIR
	err = os.Mkdir(configDir, 0755)
	if err != nil {
		return err
	}

	InitDatabaseFile()
	return nil
}

// Create a new db file on config directorty with a empty schema
func InitDatabaseFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := homeDir + constans.CONFIG_DIR + constans.DATABASE_NAME

	_, err = os.Create(configDir)
	err = db.InitSchemaDB()

	if err != nil {
		return err
	}

	return nil
}

// Validate current config show the alerts
func ValidateCurrentConfigurationAndAlert() bool {
	if _, err := findConfiguration(constans.CONFIG_DIR); err != nil {
		if conf := prompt.ConfirmCreateConfig(); !conf {
			prompt.ShowAlert(constans.ERROR_CONFIG, constans.FAIL_ALERT)
			return false
		}
		if err := CreateNewConfig(); err != nil {
			prompt.ShowAlert(constans.ERROR_CRATE_CONFIG, constans.FAIL_ALERT)
			return false
		}
	}

	return true
}
