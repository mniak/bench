package bench

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mniak/bench/lib/toolchain"
)

func findPlausibleSourceExtensions(toolchains []toolchain.Toolchain, filename string) []string {
	fileExt := filepath.Ext(filename)
	for _, tc := range toolchains {
		if fileExt == tc.OutputExtension() {
			return tc.InputExtensions()
		}
	}
	return []string{}
}

func changeExtension(filename, newExtension string) string {
	ext := filepath.Ext(filename)
	filenameWithoutExtension := strings.TrimSuffix(filename, ext)
	return filenameWithoutExtension + newExtension
}

func findPlausibleSource(toolchains []toolchain.Toolchain, filename string) (string, error) {
	extensions := findPlausibleSourceExtensions(toolchains, filename)
	for _, ext := range extensions {
		filenameWithNewExtension := changeExtension(filename, ext)
		_, err := os.Stat(filenameWithNewExtension)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
	}
	return "", nil
}
