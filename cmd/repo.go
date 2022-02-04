package cmd

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repo struct {
	gorm.Model
	Repo           string
	Dir            string
	PreviousBranch string
}

func initDB(databaseName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	db.AutoMigrate(&Repo{})
	return db, nil
}

func SaveBranch(branchName string, repoName string) error {
	log.Println(repoName)
	configPath, _ := findConfiguration(CONFIG_DIR)
	db, _ := initDB(configPath + "brancher.db")
	db.Create(&Repo{Repo: repoName, Dir: configPath, PreviousBranch: branchName})
	return nil
}
