package cmd

import (
	"github.com/fatih/color"
)

func showAlert(message string, alertCode int) {
	switch alertCode {
	case 1:
		color.Red(message)
	case 2:
		color.Cyan(message)
	default:
		color.White(message)
	}
}
