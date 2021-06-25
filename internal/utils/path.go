package utils

import (
	"os"
	"path/filepath"
	"strings"
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

func ChangeExtension(filename, newExtension string) string {
	ext := filepath.Ext(filename)
	filenameWithoutExtension := strings.TrimSuffix(filename, ext)
	return filenameWithoutExtension + newExtension
}
