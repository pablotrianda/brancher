package cmd

import (
	"errors"
	"fmt"
	"io"
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

// TODO build a install/update script
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
	return nil
}

// Copy the db file located at /files/ to the directory configured at
// constans.go in the DATABASE_NAME constant
func InitDatabaseFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := homeDir + CONFIG_DIR + DATABASE_NAME

	err = copy("../files/brancher.db", configDir)

	if err != nil {
		return err
	}

	return nil
}

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}
