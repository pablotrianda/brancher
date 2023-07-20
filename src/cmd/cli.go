package cmd

type Cli struct {
	HasArgument          bool
	NameNewBranch        string
	NameDeleteBranch     string
	BackToPreviousBranch bool
	MakeStash            bool
}

func (cli *Cli) FillDefaults() {
	cli.HasArgument = false
	cli.NameNewBranch = ""
	cli.NameDeleteBranch = ""
	cli.BackToPreviousBranch = false
	cli.MakeStash = false
}
