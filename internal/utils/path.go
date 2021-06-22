package utils

import (
	"os"
	"path/filepath"
)

func SplitDirAndProgram(dirOrProgram string) (string, string, error) {
	full, err := filepath.Abs(dirOrProgram)
	if err != nil {
		return "", "", err
	}

	info, err := os.Stat(full)
	if err != nil {
		return "", "", err
	}

	if info.IsDir() {
		return full, "", nil
	}

	return filepath.Dir(full), filepath.Base(full), nil
}
