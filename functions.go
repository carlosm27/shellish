package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pterm/pterm"
)

func CurrentPath() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}
	return path, err
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

	for _, s := range fileName {

		sizeFile, err := SizeFile(s)
		if err != nil {
			log.Println(err)
		}

		d = append(d, []string{s, sizeFile})
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
	return err

}
