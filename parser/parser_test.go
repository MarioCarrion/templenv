// Copyright © 2017 Mario Carrion <mario@carrion.ws>
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

package parser_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/MarioCarrion/templenv/parser"
)

func TestParseFileWithGetEnv(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "templenv")
	if err != nil {
		log.Fatal("Error while creating temp file")
	}
	defer os.Remove(file.Name())

	os.Setenv("USER999", "mario")
	text := `Hello {{ getEnv "USER999" }}`

	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal("Error while creating temp file")
	}

	expected := "Hello mario"
	if result, _ := parser.ParseFile(file.Name()); result != expected {
		t.Errorf("Expected result %s, but it was %s instead.", expected, result)
	}
}

func TestParseWithLoadEnvFilename(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "templenv")
	if err != nil {
		log.Fatal("Error while creating temp file")
	}
	defer os.Remove(file.Name())

	text := `# Nothing
     ABC=hello world     x    

INVALID1=     
INVALID2    
HI=1238475`
	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal("Error while creating temp file")
	}

	text = `START
{{ loadEnvFilename " - " "` + file.Name() + `" }}
END`
	expected := `START
 - ABC: hello world     x
 - HI: 1238475
END`
	if result, _ := parser.Parse("filename.tmpl", text); result != expected {
		t.Errorf("Expected result\n%s\nbut it was instead:\n%s", expected, result)
	}
}
