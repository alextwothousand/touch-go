package main

import (
	"os"

	"github.com/fatih/color"
)

func main() {
	args := os.Args[1:]

	for _, fileName := range args {
		if _, err := os.Stat(fileName); err == nil {
			c := color.New(color.FgBlack, color.BgRed, color.Bold)
			b := color.New(color.Bold)

			c.Printf("WARN:")
			b.Printf(" File %s already exists. \n", fileName)

			continue
		}

		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
}
