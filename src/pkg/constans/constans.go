package constans

// Commands
const GITCOMMAND = "git for-each-ref --sort=committerdate refs/heads/ --format='%(refname:short),'"
const GIT_GET_NAME = "git rev-parse --abbrev-ref HEAD"
const GIT_GET_DIR = "git rev-parse --show-toplevel"

// Erros messages
const ERROR_CREATE = "Error when tried to create a new branch"
const ERROR_CHANGE = "Error when tried to change to another branch"
const ERROR_NOT_BRANCHES = "The current repo hasn't git branches."
const ERROR_SAVE_BRANCH = "Can't get the branch actual name "
const ERROR_DELETE_BRANCH = "Can't delete the branch "
const ERROR_CONFIG = "Can't save the configuration"
const ERROR_CRATE_CONFIG = "Can't save the configuration"
const ERROR_FETCH_BRANCHES = "Error when tried to fetch all branches from remote"

// Alert codes
const FAIL_ALERT = 1
const MAXIM_ALERT = 1
const SUCCESS_ALERT = 2

// Directory and Database
const CONFIG_DIR = "/.config/brancher/"
const DATABASE_NAME = "brancher.db"
