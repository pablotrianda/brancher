package cmd

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repo struct {
	gorm.Model
	Repo           string
	Dir            string
	PreviousBranch string
}

func initDB() (*gorm.DB, error) {
	configPath, _ := findConfiguration(CONFIG_DIR)
	databaseName := configPath + "brancher.db"

	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	db.AutoMigrate(&Repo{})

	return db, nil
}

func SaveBranch(branchName string, repoName string, repoPwd string) error {
	db, _ := initDB()

	repo, err := findByRepoName(db, repoName)
	if err != nil {
		return err
	}

	if repo.PreviousBranch != "" {
		db.Model(&repo).Update("PreviousBranch", branchName)
	} else {
		db.Create(&Repo{Repo: repoName, Dir: repoPwd, PreviousBranch: branchName})
	}

	return nil
}

func findByRepoName(db *gorm.DB, repoName string) (Repo, error) {
	var repo Repo
	result := db.First(&repo, "repo = ?", repoName)
	return repo, result.Error
}

func GetPreviousBranchName(repoName string) (string, error) {
	db, _ := initDB()
	// Read prevBranch
	repo, err := findByRepoName(db, repoName)
	if err != nil {
		return "", err
	}

	return repo.PreviousBranch, nil
}
