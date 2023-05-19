package cmd

func Brancher(cli Cli) {
	if !validateCurrentConfigurationAndAlert(){
		return
	}
	
	handleBranch(cli)
}
