package bench

import (
	"io/ioutil"
	"os"
	"path"
)

type ProgramFinder interface {
	Find(string) (string, error)
}
type _ProgramFinder struct{}

func (pf *_ProgramFinder) Find(filenameOrFolder string) (string, error) {
	info, err := os.Stat(filenameOrFolder)
	if os.IsNotExist(err) {
		return filenameOrFolder, nil
	}
	if err != nil {
		return "", err
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

var DefaultProgramFinder ProgramFinder = new(_ProgramFinder)
