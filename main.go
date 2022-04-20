package main

import (
	//"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

var qs = []*survey.Question{
	{
		Name: "cmd",
		Prompt: &survey.Select{
			Message: "Hi, choose a cmd:",
			Options: []string{"dir", "cd", "cd..", "exit"},
			Default: "dir",
		},
	},
}

func main() {
	//reader := bufio.NewReader(os.Stdin)

	answers := struct {
		Name string // survey will match the question and field names
		cmd  string `survey:"cmd"` // or you can tag fields to match a specific name

	}{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s chose %s.", answers.Name, answers.cmd)
	for {
		fmt.Print("> " + CurrentPath() + " ")

		// Read the keyboad input.

		if answers.cmd == "dir" {
			files, _ := ListFiles()
			fmt.Println(files)

		} else if err = execInput(answers.cmd); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("path required")

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\r\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":

		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])

	case "cd..":

		path, _ := os.Getwd()
		fatherDir := filepath.Dir(path)
		return os.Chdir(fatherDir)

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
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
