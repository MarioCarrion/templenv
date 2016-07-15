// Copyright © 2016 Mario Carrion <mario@carrion.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ddclib

import (
	"log"
	"os"
	"os/user"
	"strings"
)

type InitCommand struct {
	Name    string
	homeDir string
}

func NewInitCommand(name string) *InitCommand {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return &InitCommand{
		Name:    name,
		homeDir: usr.HomeDir,
	}
}

func (cmd *InitCommand) Execute(overwrite bool) {
	if len(cmd.Name) == 0 {
		log.Fatal("error: workspace name is required")
	}

	result, err := cmd.workspaceFilenameExists()
	if result == false && err != nil {
		log.Fatal(err)
	}

	if result == false || overwrite == true {
		_, err = os.Create(cmd.filename(cmd.Name))
		if err != nil {
			log.Fatal(err)
		}

		if overwrite == true {
			log.Println("workspace overwritten")
		} else {
			log.Println("workspace created")
		}
	} else if overwrite == false {
		log.Println("workspace file exists, selecting it")
	}

	log.Printf("Executing! %s overwrite %t\n", cmd.Name, overwrite)
}

func (cmd *InitCommand) filename(name string) string {
	return strings.Join([]string{cmd.homeDir, "/.", name, ".ddc"}, "")
}

func (cmd *InitCommand) workspaceFilenameExists() (bool, error) {
	if _, err := os.Stat(cmd.filename(cmd.Name)); os.IsNotExist(err) {
		return false, nil
	} else {
		return true, nil
	}
}
