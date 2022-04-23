package main

import (
	"github.com/pterm/pterm"
)

func main() {
	pterm.DisableColor()

	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Shellish")).Srender()
	pterm.DefaultCenter.Println(s)

	Choices()

	for {
		Choices()
	}
}
