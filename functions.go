package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pterm/pterm"
)

func Choices() {
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

func CurrentPath() string {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	return path
}

func ListFiles() ([]string, error) {
	var files []string
	path, _ := os.Getwd()
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil

}

func SizeFile(name string) (fileSize string) {

	files, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	fileSize = fmt.Sprint(files.Size())
	return

}
func FilesTable() {
	d := pterm.TableData{{"File Name", "Size(bytes)"}}
	name, _ := ListFiles()

	for _, s := range name {
		d = append(d, []string{s, SizeFile(s)})
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()

}

func BackToParentFolder() (err error) {
	err = os.Chdir("../")
	return err
}
func ChangingDirectory(directory string) (err error) {
	err = os.Chdir(directory)
	return err
}
