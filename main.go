package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pterm/pterm"
)

func main() {
	pterm.DisableColor()

	bigText, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Shellish")).Srender()
	pterm.DefaultCenter.Println(bigText)

	for {

		cmd := ListOptions()
		Cases(cmd)

	}
}
func ListOptions() (cmd string) {
	cmd = ""
	prompt := &survey.Select{
		Message: "Hi, choose a cmd:",
		Options: []string{"dir", "cd", "cd..", "current path", "exit"},
		Default: "dir",
	}
	err := survey.AskOne(prompt, &cmd)

	if err != nil {
		fmt.Println(err.Error())

	}
	return cmd

}
func Cases(cmd string) {

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

		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)

	case "cd..":
		err := os.Chdir("../")
		if err != nil {
			log.Println(err)
		}

		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)

	case "current path":
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)

	case "exit":
		os.Exit(0)
	}

}

func ListFiles() ([]string, error) {

	var files []string

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, err

}

func SizeFile(name string) (string, error) {

	files, err := os.Stat(name)
	if err != nil {
		log.Println(err)
	}

	fileSize := fmt.Sprint(files.Size())
	return fileSize, err

}
func FilesTable() error {

	d := pterm.TableData{{"File Name", "Size(bytes)"}}

	fileName, err := ListFiles()
	if err != nil {
		log.Println(err)
	}

	for _, file := range fileName {

		sizeFile, err := SizeFile(file)
		if err != nil {
			log.Println(err)
		}

		d = append(d, []string{file, sizeFile})
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
	return err

}
