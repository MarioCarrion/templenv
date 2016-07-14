package ddclib

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

type InitCommand struct {
	Name string
}

func (cmd *InitCommand) Execute(overwrite bool) {
	if len(cmd.Name) == 0 {
		fmt.Fprintf(os.Stderr, "error: workspace is empty\n")
		os.Exit(1)
	}

	result, err := cmd.workspaceFilenameExists()
	if result == false && err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if result == false {
		fmt.Printf("create file")
	} else if overwrite == false {
		fmt.Println("workspace file exists, selecting it")
	} else if overwrite == true {
		fmt.Println("workspace file exists, overwriting")
	}

	fmt.Printf("Executing! %s overwrite %t\n", cmd.Name, overwrite)
}

func (cmd *InitCommand) workspaceFilenameExists() (bool, error) {
	usr, err := user.Current()

	if err != nil {
		return false, err
	}

	filename := strings.Join([]string{usr.HomeDir, "/.", cmd.Name, ".ddc"}, "")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else {
		return true, nil
	}
}
