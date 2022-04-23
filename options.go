package main

import (
	"fmt"

	"os"

	"github.com/AlecAivazis/survey/v2"
)

func Options() {
	cmd := "cmd"
	prompt := &survey.Select{
		Message: "Hi, choose a cmd:",
		Options: []string{"dir", "cd", "cd..", "current path", "exit"},
		Default: "dir",
	}
	err := survey.AskOne(prompt, &cmd)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch cmd {
	case "dir":
		FilesTable()

	case "cd":
		directory := ""
		prompt := &survey.Input{
			Message: "Write the name of a child directory:",
		}
		survey.AskOne(prompt, &directory)
		ChangingDirectory(directory)
		fmt.Println(CurrentPath())

	case "cd..":
		BackToParentFolder()
		fmt.Println(CurrentPath())

	case "current path":
		fmt.Println(CurrentPath())

	case "exit":
		os.Exit(0)
	}

}
