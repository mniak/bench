package impl

import (
	"os"
	"path/filepath"

	"github.com/mniak/bench/domain"
)

type _FinderOnDirByFilenameAndExtensions struct {
	extensions []string
	filenames  []string
}

func (f *_FinderOnDirByFilenameAndExtensions) Find(filenameOrFolder string) (string, error) {
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

	return "", domain.ErrProgramNotFound
}

func NewFinderOnDirByFilenamesAndExtensions(filenames, extensions []string) domain.FileFinder {
	return &_FinderOnDirByFilenameAndExtensions{
		filenames:  filenames,
		extensions: extensions,
	}
}
