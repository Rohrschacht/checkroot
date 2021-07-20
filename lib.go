package checkroot

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Isroot checks whether the current program is running as root and returns true
// if it is
func Isroot() (bool, error) {
	stdout, err := exec.Command("id").Output()
	if err != nil {
		return false, fmt.Errorf("isroot: could not execute 'id': %v", err)
	}

	if strings.TrimSpace(string(stdout)) == "uid=0(root) gid=0(root) groups=0(root)" {
		return true, nil
	}

	return false, nil
}

// RootOrExit prints an error and exits the program with exit code 1 if it is
// not running as root
func RootOrExit() {
	if b, err := Isroot(); err != nil || !b {
		fmt.Fprintf(os.Stderr, "This script needs root access!\n")
		os.Exit(1)
	}
}
