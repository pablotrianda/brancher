package cmd

import (
	"testing"

	"github.com/pablotrianda/brancher/cmd"
	//"github.com/pablotrianda/brancher/cmd"
)

func check(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	newRepo := cmd.Repo{Repo: "a_new_one", Dir: "/home/user", PreviousBranch: "develop"}

	repo, err := cmd.InsertOrUpdateRepo(newRepo)
	if repo.Repo != "a_new_one" {
		t.Fatal("Fail to create a new Repo")
	}
	check(err, t)

}

func TestUpdate(t *testing.T) {
	oldRepo := cmd.Repo{Repo: "a_new_one", Dir: "/home/user/code", PreviousBranch: "master"}
	modifiedRepo, err := cmd.InsertOrUpdateRepo(oldRepo)
	if modifiedRepo.Dir != "/home/user/code" {
		t.Fatal("No updated the repo")
	}
	check(err, t)

}

func TestFindByRepoName(t *testing.T) {
	db, err := cmd.InitDB()
	check(err, t)

	repo, err := cmd.FindByRepoName(db, "a_new_one")
	check(err, t)

	if repo.Repo != "a_new_one" {
		t.Fatal("The repo was not found")
	}

	defer db.Close()
}

func TestGetPreviousBranchName(t *testing.T) {
	db, err := cmd.InitDB()
	check(err, t)

	prevBranch, err := cmd.GetPreviousBranchName("a_new_one")
	check(err, t)

	if prevBranch != "master" {
		t.Fatalf("Fail, want: master, got %s", prevBranch)
	}

	defer db.Close()
}
