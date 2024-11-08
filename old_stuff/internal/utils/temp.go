package utils

import (
	"os"
	"path/filepath"
)

type _TempFile struct {
	*os.File
}

func TempFile(dir, name string) (_TempFile, error) {
	file, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		return _TempFile{}, err
	}
	return _TempFile{
		File: file,
	}, nil
}

func (f *_TempFile) CloseAndRemove() error {
	err := f.File.Close()
	if err != nil {
		return err
	}
	return os.Remove(f.Name())
}
