package bench

import (
	"os"
	"path/filepath"
	"runtime"
)

type ProgramFinder interface {
	Find(string) (string, error)
}

type _ProgramFinder struct {
	extensions []string
	filenames  []string
}

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

	// files, err := ioutil.ReadDir(filenameOrFolder)
	// if err != nil {
	// 	return "", err
	// }
	// for _, fi := range files {
	// 	if fi.IsDir() {
	// 		continue
	// 	}
	// 	filename := fi.Name()
	// 	for _, iext := range pf.extensions {
	// 		if iext == filepath.Ext(filename) {
	// 			continue
	// 		}
	// 	}
	// 	return filepath.Join(filenameOrFolder, fi.Name()), nil
	// }

	folderBaseName := filepath.Base(filenameOrFolder)
	filenames := append(pf.filenames, folderBaseName)

	for _, filename := range filenames {
		for _, extension := range pf.extensions {
			full := filepath.Join(filenameOrFolder, filename+extension)
			_, err := os.Stat(full)
			if err == nil {
				return full, nil
			}
		}
	}

	return "", ErrProgramNotFound
}

var defaultProgramFinder *_ProgramFinder = &_ProgramFinder{
	filenames: []string{"main"},
}
var DefaultProgramFinder ProgramFinder = defaultProgramFinder

func init() {
	switch runtime.GOOS {
	case "windows":
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".exe")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".bat")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".cmd")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".ps1")
	default:
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, "")
	}
}
