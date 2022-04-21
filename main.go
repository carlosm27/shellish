package main

import (
	//"bufio"
	//"errors"
	"fmt"
	"log"
	"os"

	//"os/exec"

	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pterm/pterm"
	//"path/filepath"
	//"strings"
)

func main() {
	pterm.DisableColor()
	//reader := bufio.NewReader(os.Stdin)
	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Shellish")).Srender()
	pterm.DefaultCenter.Println(s)

	Choices()

	for {
		Choices()
	}
}

func Choices() {
	cmd := "cmd"
	prompt := &survey.Select{
		Message: "Hi, choose a cmd:",
		Options: []string{"dir", "cd", "cd..", "exit", "current"},
		Default: "dir",
	}
	err := survey.AskOne(prompt, &cmd)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if cmd == "dir" {
		FilesTable()

	} else if cmd == "exit" {
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
