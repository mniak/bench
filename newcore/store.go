package newcore

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/mniak/bench/pkg/cache"
)

type ToolchainLoader interface {
	Load() (Toolchain, error)
	ToolchainType() reflect.Type
}

type (
	ToolchainsList []Toolchain
	Toolchain      interface{}
)

var toolchainLoaders = []ToolchainLoader{
	new(GoLoader),
	new(PythonLoader),
}

func MarshalList[T any](list []T) ([]byte, error) {
	var result []any
	for _, r := range list {
		v := reflect.ValueOf(r).Elem()
		result = append(result, map[string]any{
			"name":   v.Type().Name(),
			"params": v.Interface(),
		})
	}
	return json.Marshal(result)
}

func UnmarshalList[T any](known map[string]reflect.Type, b []byte) ([]T, error) {
	jsonList := make([]struct {
		Type      string          `json:"name"`
		RawParams json.RawMessage `json:"params"`
	}, 0)
	if err := json.Unmarshal(b, &jsonList); err != nil {
		return nil, err
	}

	var list []T
	for _, item := range jsonList {
		type_, found := known[item.Type]
		if !found {
			log.Printf("Type %q not found", item.Type)
			continue
		}

		instance := reflect.New(type_).Interface().(T)
		if err := json.Unmarshal(item.RawParams, instance); err != nil {
			return nil, err
		}
		list = append(list, instance)
	}
	return list, nil
}

func (list ToolchainsList) MarshalJSON() ([]byte, error) {
	return MarshalList[Toolchain](list)
}

func (list *ToolchainsList) UnmarshalJSON(b []byte) error {
	known := make(map[string]reflect.Type)
	for _, l := range toolchainLoaders {
		t := l.ToolchainType()
		known[t.Name()] = t
	}

	result, err := UnmarshalList[Toolchain](known, b)
	if err != nil {
		return err
	}
	*list = result
	return nil
}

func RebuildCache() (ToolchainsList, error) {
	toolchains := loadToolchains()
	err := cache.JSONStore("toolchains.json", toolchains)
	if err != nil {
		return nil, err
	}
	return toolchains, err
}

func loadToolchains() ToolchainsList {
	var toolchains ToolchainsList
	for _, loader := range toolchainLoaders {
		loaderTypeName := reflect.TypeOf(loader).Elem().Name()
		r, err := loader.Load()
		if err != nil {
			log.Printf("Failed to load toolchain %T", loaderTypeName)
			continue
		}
		log.Printf("Toolchain %s loaded", loaderTypeName)
		toolchains = append(toolchains, r)
	}
	return toolchains
}

// Toolchains returns a list of toolchains, using a cache
func Toolchains() ToolchainsList {
	result, err := cache.JSONCache("toolchains.json", func() (ToolchainsList, error) {
		toolchains := loadToolchains()
		return toolchains, nil
	})
	if err != nil {
		log.Printf("Failed to load toolchains from cache: %s", err)
	}
	return result
}

func iterateToolchains[T Toolchain](yield func(T) bool) {
	for _, toolchain := range Toolchains() {
		typed, ok := toolchain.(T)
		if !ok {
			continue
		}
		end := yield(typed)
		if end {
			return
		}
	}
}
