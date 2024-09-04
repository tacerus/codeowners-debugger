/*
   CODEOWNERS debugger
   Copyright (C) 2024  SUSE LLC <georg.pfuetzenreuter@suse.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
