package oldall

import (
	"encoding/json"
	"log"
	"os"
	"reflect"

	"github.com/mniak/bench/old_stuff/toolchain"
	"github.com/mniak/bench/pkg/cache"
	"github.com/pkg/errors"
)

var toolchainLoaders = []ToolchainLoader{
	toolchain.NewCPPLoader(),
	toolchain.NewGoLoader(),
}

type (
	ToolchainsList []Toolchain
	Named          interface {
		Name() string
	}
)

func MarshalNamedList[T Named](list []T) ([]byte, error) {
	var result []any
	for _, r := range list {
		v := reflect.ValueOf(r).Elem()
		result = append(result, map[string]any{
			"kind":   r.Name(),
			"params": v.Interface(),
		})
	}
	return json.Marshal(result)
}

func UnmarshalNamedList[T Named](known map[string]reflect.Type, b []byte) ([]T, error) {
	jsonList := make([]struct {
		Kind      string          `json:"kind"`
		RawParams json.RawMessage `json:"params"`
	}, 0)
	if err := json.Unmarshal(b, &jsonList); err != nil {
		return nil, err
	}

	var list []T
	for _, item := range jsonList {
		type_, found := known[item.Kind]
		if !found {
			log.Printf("Kind %q not found", item.Kind)
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
	return MarshalNamedList[Toolchain](list)
}

func (list *ToolchainsList) UnmarshalJSON(b []byte) error {
	known := make(map[string]reflect.Type)
	for _, l := range toolchainLoaders {
		known[l.Name()] = l.ToolchainType().Elem()
	}

	result, err := UnmarshalNamedList[Toolchain](known, b)
	if err != nil {
		return err
	}
	*list = result
	return nil
}

func loadToolchains() ToolchainsList {
	var toolchains ToolchainsList
	for _, loader := range toolchainLoaders {
		r, err := loader.Load()
		if err != nil {
			log.Printf("Failed to load toolchain %q", loader.Name())
			continue
		}
		log.Printf("Toolchain %q loaded", loader.Name())
		toolchains = append(toolchains, r)
	}
	return toolchains
}

func RebuildCache() (ToolchainsList, error) {
	toolchains := loadToolchains()
	err := cache.JSONStore("toolchains.json", toolchains)
	if err != nil {
		return nil, err
	}
	return toolchains, err
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

// ToolchainFor tries to find a suitable toolchain for a specific file
func ToolchainFor(filename string) (Toolchain, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	for _, toolchain := range Toolchains() {
		can := toolchain.CanRun(filename)
		if can {
			return toolchain, nil
		}
	}
	return nil, errors.New("no suitable toolchain found for file")
}
