package main

import (
	"fmt"
	"github.com/lgmonezi/file-ninja-go/repository"
	"github.com/lgmonezi/file-ninja-go/utils"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func dirtyFilesMenu() {
	prompt := "What do you want to do about the Dirty Files"
	options := []string{
		"List Root Folders",
		"Add Root Folder",
		"Remove Root Folder",
		"Return to Previous Menu",
	}

	for {
		input, _ := utils.ChooseFromOptions(prompt, options)
		switch input {
		case 1:
			listDirtyFilesRootFolders()
		case 2:
		case 3:
		case 4:
			return
		}
	}
}

func listDirtyFilesRootFolders() {
	defer fmt.Println()

	folders, err := repository.GetDirtyRootFolders()
	if err != nil {
		log.Fatalln(err)
	}

	if len(folders) == 0 {
		utils.Print2Ln("0 Results")
		return
	}
	var data [][]string
	for _, folder := range folders {
		isClean := "false"
		if folder.Clean {
			isClean = "true"
		}
		data = append(data, []string{folder.Name, folder.Path, folder.Type, isClean})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Path", "Type", "Clean Folder"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
