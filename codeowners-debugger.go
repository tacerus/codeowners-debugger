package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hairyhenderson/go-codeowners"
)

func main() {
	var directory string
	var file string

	flag.StringVar(&directory, "directory", "", "Path to the Git repository")
	flag.StringVar(&file, "file", "", "Path to a file to test, relative to the directory")
	flag.Parse()

	if directory == "" || file == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Checking %s/%s\n", directory, file)

	c, err := codeowners.FromFile(directory)
	if err != nil {
		fmt.Printf("Failed to read CODEOWNERS: %s\n", err.Error())
		os.Exit(1)
	}

	owners := c.Owners(file)
	for i, o := range owners {
		fmt.Printf("Owner #%d is %s\n", i, o)
	}
}
