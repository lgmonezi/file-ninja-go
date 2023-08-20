package main

import (
	"github.com/lgmonezi/file-ninja-go/utils"
	"os"
)

func mainMenu() {
	prompt := "What do you want to do"
	options := []string{
		"Clean Folder Settings",
		"Dirty Folder Settings",
		"Scan Files",
		"Organize Files",
		"Quit",
	}

	for {
		input, _ := utils.ChooseFromOptions(prompt, options)
		switch input {
		case 1:
		case 2:
			dirtyFilesMenu()
		case 3:
		case 4:
		case 5:
			quitApplication()
		}
	}
}

func quitApplication() {
	utils.Print2Ln("Bye")
	os.Exit(0)
}
