package ddclib

import "fmt"

type InitCommand struct {
  Name string
}

func (i *InitCommand) Execute(overwrite bool) {
  fmt.Printf("Executing! %s overwrite %t\n", i.Name, overwrite)
}
