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
