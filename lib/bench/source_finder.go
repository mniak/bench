package bench

import (
	"os"
	"path/filepath"

	"github.com/mniak/bench/internal/utils"
	"github.com/mniak/bench/lib/toolchain"
)

type _SourceFinderByToolchain struct {
	toolchains []toolchain.Toolchain
}

func (f *_SourceFinderByToolchain) Find(filename string) (string, error) {
	extensions := f.FindExtensions(filename)
	for _, ext := range extensions {
		filenameWithNewExtension := utils.ChangeExtension(filename, ext)
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

func (f *_SourceFinderByToolchain) FindExtensions(filename string) []string {
	fileExt := filepath.Ext(filename)
	for _, tc := range f.toolchains {
		if fileExt == tc.OutputExtension() {
			return tc.InputExtensions()
		}
	}
	return []string{}
}
