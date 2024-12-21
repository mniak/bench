package cache

import "encoding/json"

type jsonMarshaler[T any] struct{}

func (jsonMarshaler[T]) Marshal(obj T) ([]byte, error) {
	return json.MarshalIndent(obj, "", "  ")
}

func (jsonMarshaler[T]) Unmarshal(data []byte) (T, error) {
	var obj T
	err := json.Unmarshal(data, &obj)
	return obj, err
}

func JSONCache[T any](key string, fallback FallbackFunc[T]) (T, error) {
	var m MarshalUnmarshaler[T] = jsonMarshaler[T]{}
	return Cache(m, key, fallback)
}

func JSONStore[T any](key string, obj T) error {
	var m MarshalUnmarshaler[T] = jsonMarshaler[T]{}
	return Store(m, key, obj)
}
