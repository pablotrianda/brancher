package cmd

// Commands
const GITCOMMAND = "git for-each-ref --sort=committerdate refs/heads/ --format='%(refname:short),'"
const GIT_GET_NAME = "git rev-parse --abbrev-ref HEAD"
const GIT_GET_DIR = "git rev-parse --show-toplevel"

// Erros messages
const ERROR_CREATE = "Error when tried to create a new branch"
const ERROR_CHANGE = "Error when tried to change to another branch"
const ERROR_NOT_BRANCHES = "The current repo hasn't git branches."
const ERROR_SAVE_BRANCH = "Can't get the branch actual name "
const ERROR_CONFIG = "Can't save the configuration"
const ERROR_CRATE_CONFIG = "Can't save the configuration"

// Alert codes
const FAIL_ALERT = 1

// Directory and Database
const CONFIG_DIR = "/.config/brancher/"
const DATABASE_NAME = "brancher.db"
