package bench

import (
	"os"
	"path/filepath"
	"runtime"
)

type ProgramFinder interface {
	Find(filenameOrFolder string) (string, error)
	// IsExecutable(filename string) (bool, error)
}

type _ProgramFinder struct {
	extensions []string
	filenames  []string
}

func (f *_ProgramFinder) Find(filenameOrFolder string) (string, error) {
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

	folderBaseName := filepath.Base(filenameOrFolder)
	filenames := append(f.filenames, folderBaseName)

	for _, filename := range filenames {
		for _, extension := range f.extensions {
			full := filepath.Join(filenameOrFolder, filename+extension)
			_, err := os.Stat(full)
			if err == nil {
				return full, nil
			}
		}
	}

	return "", ErrProgramNotFound
}

// func (f *_ProgramFinder) IsExecutable(filename string) (bool, error) {
// 	fileExtension := filepath.Ext(filename)
// 	for _, extension := range f.extensions {
// 		if fileExtension == extension {
// 			return true, nil
// 		}
// 	}

// 	info, err := os.Stat(filename)
// 	if err != nil {
// 		return false, err
// 	}
// 	if uint32(info.Mode().Perm()) == uint32(73) {
// 		return true, nil
// 	}
// 	return false, nil
// }

var defaultProgramFinder *_ProgramFinder = &_ProgramFinder{
	filenames:  []string{"main"},
	extensions: []string{".py"},
}
var DefaultExecutableFinder ProgramFinder = defaultProgramFinder

func init() {
	switch runtime.GOOS {
	case "windows":
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".exe")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".bat")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".cmd")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".ps1")
	default:
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, "")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".sh")
	}
}

type finderWithBuilder struct {
	sourceFinder ProgramFinder
	binaryFinder ProgramFinder
	builder      Builder
}

func IsExecutableProgram(path string) (bool, error) {
	return false, nil
}

// func (f *finderWithBuilder) Find(path string) (string, error) {
// 	result, err := f.binaryFinder.Find(path)
// 	if err != nil {
// 		return result, nil
// 	}

// 	Build(path)
// }
