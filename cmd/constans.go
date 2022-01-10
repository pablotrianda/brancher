package cmd

// Commands
const GITCOMMAND = "git for-each-ref --sort=committerdate refs/heads/ --format='%(refname:short),'"

// Erros messages
const ERROR_CREATE = "Error when tried to create a new branch"
const ERROR_CHANGE = "Error when tried to change to another branch"
const ERROR_NOT_BRANCHES = "The current repo hasn't git branches."

// Alert codes
const FAIL_ALERT = 1
