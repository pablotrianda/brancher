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

func createNewConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := homeDir + CONFIG_DIR
	err = os.Mkdir(configDir, 0755)
	if err != nil {
		return err
	}
	return nil
}
