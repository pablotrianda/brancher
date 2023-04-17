package cmd

import (
	"errors"
	"os"
)

func findConfiguration(string) (confPath string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := homeDir + CONFIG_DIR
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

	configDir := homeDir + CONFIG_DIR
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

	configDir := homeDir + CONFIG_DIR + DATABASE_NAME

	_, err = os.Create(configDir)
	err = InitSchemaDB()

	if err != nil {
		return err
	}

	return nil
}

// Validate current config show the alerts
func validateCurrentConfiguration() bool{
	if _, err := findConfiguration(CONFIG_DIR); err != nil {
		if conf := confirmCreateConfig(); !conf {
			showAlert(ERROR_CONFIG, FAIL_ALERT)
			return false
		}
		if err := CreateNewConfig(); err != nil {
			showAlert(ERROR_CRATE_CONFIG, FAIL_ALERT)
			return false
		}
	}

	return true
}

// Validate current config show the alerts
func validateCurrentConfigurationAndAlert() bool{
	if _, err := findConfiguration(CONFIG_DIR); err != nil {
		if conf := confirmCreateConfig(); !conf {
			showAlert(ERROR_CONFIG, FAIL_ALERT)
			return false
		}
		if err := CreateNewConfig(); err != nil {
			showAlert(ERROR_CRATE_CONFIG, FAIL_ALERT)
			return false
		}
	}

	return true
}
