[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
![status](https://github.com/pablotrianda/brancher/actions/workflows/go.yml/badge.svg)

# BRANCHER 🔍🌿!!.
## Tool to manage git branches
<p align="center">
   <img src="https://i.imgur.com/vYqF0sz.png" data-canonical-src="https://i.imgur.com/vYqF0sz.png" width="200" height="250" />
</p>

# Usage
* Download the binary from releases section. 
* Run `$ brancher` and select the brach to make the checkout in the menu or you can run `$ brancher branch_name` directly
* Use the command `-n` to create a new branch

# Example
![Bancher](https://media0.giphy.com/media/d6zP9HA60tiG788xkX/giphy.gif?cid=790b7611cf30827b13c0d1d134eb43844f90b94637fa065a&rid=giphy.gif&ct=g)


# Run on development mode
* Clone this repo `git clone https://github.com/pablotrianda/brancher.git`
* Go to folder `cd brancher`
* Install the dependencies `go mod tidy`
* Run `go run *.go`

# Extra - Build using golang with Docker
* Run `make build`

# Next steps
- [x] Show error in a friendly way.  
- [ ] Remember previous branch
- [ ] Create a fix conflicted automated steps. [THIS](https://dev.to/smetankajakub/how-to-resolve-merge-conflicts-in-bitbucket-repository-with-git-bash-34ag)
- [ ] Build a neovim plugin ;) 

