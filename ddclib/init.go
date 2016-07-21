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
)

type InitCommand struct {
	Name string
	*session
}

func NewInitCommand(name string) *InitCommand {
	return &InitCommand{
		Name:    name,
		session: newSession(),
	}
}

func (cmd *InitCommand) Execute(overwrite bool) {
	if len(cmd.Name) == 0 {
		log.Println("using workspace:", cmd.current())
		os.Exit(0)
	}

	setCurrent := false
	if (cmd.current() == "") || (cmd.current() != cmd.Name && overwrite == true) {
		setCurrent = true
	} else if cmd.current() != cmd.Name {
		log.Fatal("error: use --overwrite the change the session")
	}

	if setCurrent == true {
		cmd.setCurrent(cmd.Name)
	}

	log.Printf("Executing! %s overwrite %t\n", cmd.Name, overwrite)
}
