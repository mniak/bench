package bench

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/mniak/bench/lib/toolchain"
)

var ErrProgramNotFound = errors.New("program not found")

func findMain(filenameOrFolder string) (string, error) {
	info, err := os.Stat(filenameOrFolder)
	if os.IsNotExist(err) {
		return "", ErrProgramNotFound
	}

	if !info.IsDir() {
		return filenameOrFolder, nil
	}

	ignoredExtensions := []string{
		".exe", ".dll", ".obj",
		".o", ".so",
	}
	files, err := ioutil.ReadDir(filenameOrFolder)
	if err != nil {
		return "", err
	}
	for _, fi := range files {
		if fi.IsDir() {
			continue
		}
		filename := fi.Name()
		for _, iext := range ignoredExtensions {
			if iext == path.Ext(filename) {
				continue
			}
		}
		return path.Join(filenameOrFolder, fi.Name()), nil
	}

	return "", ErrProgramNotFound
}

func buildToolchain(mainfile string) (toolchain.Toolchain, error) {
	switch path.Ext(mainfile) {
	case ".cpp", ".c++":
		return toolchain.NewCPP()
	}
	return nil, toolchain.ErrToolchainNotFound
}

func Build(path string) error {
	mainfile, err := findMain(path)
	if err != nil {
		return err
	}
	tchain, err := buildToolchain(mainfile)
	if err != nil {
		return err
	}
	_, err = tchain.Build(mainfile)
	if err != nil {
		return err
	}
	return nil
}
