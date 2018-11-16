package utils

import (
	"os"
	"path/filepath"
)

func ExecPath() (string, error) {
	e, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(e), nil
}
