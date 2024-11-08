package cache

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

var ErrCacheMiss = errors.New("cache miss")

type (
	FallbackFunc[T any] func() (T, error)
	CacheBackend        interface {
		Store(key string, data []byte) error
		Load(key string) ([]byte, error)
		Invalidate(key string) error
	}
	MarshalUnmarshaler[T any] interface {
		Marshal(obj T) ([]byte, error)
		Unmarshal(data []byte) (T, error)
	}
)

var DefaultBackend CacheBackend

func init() {
	path, err := os.UserCacheDir()
	if err != nil {
		log.Printf("Could not load user cache dir: %s", err)
		DefaultBackend = _FileStore{
			BaseDir: ".bench-cache",
		}
	} else {
		DefaultBackend = _FileStore{
			BaseDir: filepath.Join(path, "bench", "cache"),
		}
	}
}

func BinaryCache(key string, fallback FallbackFunc[[]byte]) ([]byte, error) {
	data, err := DefaultBackend.Load(key)
	if !errors.Is(err, ErrCacheMiss) {
		return data, err
	}

	data, err = fallback()
	if err == nil {
		storeErr := DefaultBackend.Store(key, data)
		log.Printf("Cache failed to store %q: %s", key, storeErr)
	}
	return data, err
}

func Cache[T any](marshaler MarshalUnmarshaler[T], key string, fallback FallbackFunc[T]) (T, error) {
	data, err := BinaryCache(key, func() ([]byte, error) {
		obj, err := fallback()
		if err != nil {
			return nil, err
		}
		return marshaler.Marshal(obj)
	})
	if err != nil {
		var obj T
		return obj, err
	}

	return marshaler.Unmarshal(data)
}

func Store[T any](marshaler MarshalUnmarshaler[T], key string, obj T) error {
	data, err := marshaler.Marshal(obj)
	if err != nil {
		return err
	}
	return DefaultBackend.Store(key, data)
}
