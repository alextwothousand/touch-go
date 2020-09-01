package main

import (
	"os"
	"regexp"

	"github.com/fatih/color"
)

var (
	executableRegex *regexp.Regexp = regexp.MustCompile(`.+\.exe`)
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		c := color.New(color.FgBlack, color.BgRed, color.Bold)
		b := color.New(color.Bold)

		c.Printf("WARN:")
		b.Printf(" No files to touch specified.\n")

		return
	}

	for _, fileName := range args {
		if executableRegex.Match([]byte(fileName)) {
			c := color.New(color.FgBlack, color.BgRed, color.Bold)
			b := color.New(color.Bold)

			c.Printf("WARN:")
			b.Printf(" .exe could not be used as a file extension for %s\n", fileName)

			continue
		}

		if _, err := os.Stat(fileName); err == nil {
			c := color.New(color.FgBlack, color.BgRed, color.Bold)
			b := color.New(color.Bold)

			c.Printf("ERROR:")
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
