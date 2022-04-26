package main

import (
	"fmt"
	"log"

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

		err := os.Chdir(directory)
		if err != nil {
			log.Println(err)
		}

		path, err := CurrentPath()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)

	case "cd..":
		err := os.Chdir("../")
		if err != nil {
			log.Println(err)
		}

		path, err := CurrentPath()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)

	case "current path":
		path, err := CurrentPath()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)

	case "exit":
		os.Exit(0)
	}

}
