package cmd

import (
	"testing"

	"github.com/pablotrianda/brancher/src/pkg/db"
)

func check(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	newRepo := db.Repo{Repo: "a_new_one", Dir: "/home/user", PreviousBranch: "develop"}

	repo, err := db.InsertOrUpdateRepo(newRepo)
	if repo.Repo != "a_new_one" {
		t.Fatal("Fail to create a new Repo")
	}
	check(err, t)

}

func TestUpdate(t *testing.T) {
	oldRepo := db.Repo{Repo: "a_new_one", Dir: "/home/user/code", PreviousBranch: "master"}
	modifiedRepo, err := db.InsertOrUpdateRepo(oldRepo)
	if modifiedRepo.Dir != "/home/user/code" {
		t.Fatal("No updated the repo")
	}
	check(err, t)

}

func TestFindByRepoName(t *testing.T) {
	currentDb, err := db.InitDB()
	check(err, t)

	repo, err := db.FindByRepoName(currentDb, "a_new_one")
	check(err, t)

	if repo.Repo != "a_new_one" {
		t.Fatal("The repo was not found")
	}

	defer currentDb.Close()
}

func TestGetPreviousBranchName(t *testing.T) {
	currentDb, err := db.InitDB()
	check(err, t)

	prevBranch, err := db.GetPreviousBranchName("a_new_one")
	check(err, t)

	if prevBranch != "master" {
		t.Fatalf("Fail, want: master, got %s", prevBranch)
	}

	defer currentDb.Close()
}
