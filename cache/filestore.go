package cache

import (
	"os"
	"path/filepath"
)

type _FileStore struct {
	BaseDir string
}

func (fs _FileStore) filename(key string) string {
	return filepath.Join(fs.BaseDir, key)
}

func (fs _FileStore) Load(key string) ([]byte, error) {
	data, err := os.ReadFile(key)

	if os.IsNotExist(err) {
		return data, ErrCacheMiss
	}
	return data, err
}

func (fs _FileStore) Store(key string, data []byte) error {
	err := os.MkdirAll(fs.filename(key), 0o777)
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.filename(key), data, 0o655)
	return err
}

func (fs _FileStore) Invalidate(key string) error {
	return os.Remove(fs.filename(key))
}