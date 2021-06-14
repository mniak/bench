package bench

import (
	"errors"
	"io/ioutil"
	"path"

	"github.com/mniak/bench/lib/toolchain"
)

var (
	ErrMainNotFound = errors.New("program main not found")
	ErrNoToolchain  = errors.New("could not get the appropriate")
)

func findMain(folder string) (string, error) {
	ignoredExtensions := []string{
		".exe", ".dll", ".obj",
		".o", ".so",
	}
	files, err := ioutil.ReadDir(folder)
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
		return path.Join(folder, fi.Name()), nil
	}

	return "", ErrMainNotFound
}

func buildToolchain(mainfile string) (toolchain.Toolchain, error) {
	// switch runtime.GOOS {
	// case "windows":
	switch path.Ext(mainfile) {
	case ".cpp":
		return toolchain.NewMSVC()
	}
	// break
	// }
	return nil, ErrNoToolchain
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
	// ran, err := tchain.Run(built)
	return nil
}
