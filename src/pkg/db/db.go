package db

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pablotrianda/brancher/src/pkg/constans"
)

type Repo struct {
	Repo           string
	Dir            string
	PreviousBranch string
}

func InitDB() (*sql.DB, error) {
	homeDir, err := os.UserHomeDir()
	databaseDir := homeDir + constans.CONFIG_DIR
	databaseName := databaseDir + constans.DATABASE_NAME

	db, err := sql.Open("sqlite3", databaseName)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveBranch(repoName string, repoDir string, previosBranch string) error {
	repo := Repo{Repo: repoName, Dir: repoDir, PreviousBranch: previosBranch}
	_, err := InsertOrUpdateRepo(repo)
	if err != nil {
		return err
	}
	return nil
}

func FindByRepoName(db *sql.DB, repoName string) (Repo, error) {
	var repo Repo

	row, err := db.Query("SELECT * FROM Repos WHERE repo = ?", repoName)

	if err != nil {
		return Repo{}, err
	}

	for row.Next() {
		row.Scan(&repo.Repo, &repo.Dir, &repo.PreviousBranch)
	}

	return repo, err
}

func GetPreviousBranchName(repoName string) (string, error) {
	db, _ := InitDB()
	// Read prevBranch
	repo, err := FindByRepoName(db, repoName)
	defer db.Close()
	if err != nil {
		return "", err
	}

	return repo.PreviousBranch, nil
}

func InsertOrUpdateRepo(repo Repo) (Repo, error) {
	db, err := InitDB()

	if err != nil {
		return Repo{}, nil
	}

	findedRepo, err := FindByRepoName(db, repo.Repo)
	if err != nil {
		return Repo{}, nil
	}

	if findedRepo.Repo == "" {
		err := insertNewRepo(db, repo)
		if err != nil {
			return repo, errors.New("Fail to insert a new repo")
		}
	} else {
		err := updateRepo(db, repo)
		if err != nil {
			return repo, errors.New("Fail to create a new repo")
		}
	}

	defer db.Close()
	return repo, nil
}

func insertNewRepo(db *sql.DB, repo Repo) error {
	stmt, err := db.Prepare("INSERT INTO Repos(repo, dir, previosBranch) values(?,?,?)")
	_, err = stmt.Exec(repo.Repo, repo.Dir, repo.PreviousBranch)

	if err != nil {
		return errors.New("Cant inser a new repo")
	}

	return nil
}

func updateRepo(db *sql.DB, repo Repo) error {
	_, err := db.Exec(
		"UPDATE Repos SET repo = ?, dir = ?, previosBranch = ? WHERE repo = ? ",
		repo.Repo,
		repo.Dir,
		repo.PreviousBranch,
		repo.Repo,
	)

	if err != nil {
		return errors.New("Cant update the repo")
	}

	return nil
}

func InitSchemaDB() error {
	db, err := InitDB()
	if err != nil {
		return err
	}
	q := `CREATE TABLE IF NOT EXISTS "Repos" (
			repo TEXT,
			dir	 TEXT,
			previosBranch TEXT
		);`

	_, err = db.Exec(q)
	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
