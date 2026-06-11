package util

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func FileExists(p string) bool {
	pwd, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(path.Join(pwd, p))
	if err != nil {
		return false
	}

	return true
}

func EnsureBinary(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("%s not found. Please install it first.", name)
	}
	return nil
}
