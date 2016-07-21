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
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

type session struct {
	homeDir     string
	currentName string
}

const current = "CURRENT"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newSession() *session {
	usr, err := user.Current()
	checkError(err)

	homeDir := strings.Join([]string{usr.HomeDir, ".ddc"}, "/")
	session := &session{
		homeDir: homeDir,
	}

	exists, err := session.exists("")
	checkError(err)

	if exists == false {
		err = os.Mkdir(homeDir, os.FileMode(int(0700)))
		checkError(err)
	}

	session.currentName = session.current()

	return session
}

func (s *session) current() string {
	if s.currentName != "" {
		return s.currentName
	}

	exists, err := s.exists(current)
	if exists == false {
		return ""
	}

	session, err := ioutil.ReadFile(s.filename(current))
	checkError(err)

	s.currentName = string(session)
	return s.currentName
}

func (s *session) exists(name string) (bool, error) {
	_, err := os.Stat(s.filename(name))

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func (s *session) filename(name string) string {
	return strings.Join([]string{s.homeDir, name}, "/")
}

func (s *session) setCurrent(newName string) {
	f, err := os.Create(s.filename(current))
	defer f.Close()

	checkError(err)

	_, err = f.WriteString(newName)
	checkError(err)

	err = f.Sync()
	checkError(err)

	s.currentName = newName
}
