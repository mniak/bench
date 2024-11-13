package newcore

import (
	"encoding/json"
	"log"
	"os"
	"reflect"

	"github.com/mniak/bench/pkg/cache"
	"github.com/pkg/errors"
)

var runnerLoaders = []RunnerLoader{
	new(_PythonLoader),
	new(BinaryLoader),
}

type (
	RunnersList []Runner
	Named       interface {
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

func (list RunnersList) MarshalJSON() ([]byte, error) {
	return MarshalNamedList[Runner](list)
}

func (list *RunnersList) UnmarshalJSON(b []byte) error {
	known := make(map[string]reflect.Type)
	for _, l := range runnerLoaders {
		known[l.Name()] = l.RunnerType()
	}

	result, err := UnmarshalNamedList[Runner](known, b)
	if err != nil {
		return err
	}
	*list = result
	return nil
}

func loadRunners() RunnersList {
	var runners RunnersList
	for _, loader := range runnerLoaders {
		r, err := loader.LoadRunner()
		if err != nil {
			log.Printf("Failed to load runner %q", loader.Name())
			continue
		}
		log.Printf("Runner %q loaded", loader.Name())
		runners = append(runners, r)
	}
	return runners
}

var compilerLoaders = []CompilerLoader{
	new(_GoLoader),
}

type CompilersList []Compiler

func (list CompilersList) MarshalJSON() ([]byte, error) {
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

func (list *CompilersList) UnmarshalJSON(b []byte) error {
	known := make(map[string]reflect.Type)
	for _, l := range compilerLoaders {
		known[l.Name()] = l.CompilerType()
	}

	result, err := UnmarshalNamedList[Compiler](known, b)
	if err != nil {
		return err
	}
	*list = result
	return nil
}

func loadCompilers() CompilersList {
	var compilers CompilersList
	for _, loader := range compilerLoaders {
		r, err := loader.LoadCompiler()
		if err != nil {
			log.Printf("Failed to load compiler %q", loader.Name())
			continue
		}
		log.Printf("Compiler %q loaded", loader.Name())
		compilers = append(compilers, r)
	}
	return compilers
}

func RebuildCache() (RunnersList, CompilersList, error) {
	runners := loadRunners()
	err := cache.JSONStore("runners.json", runners)
	if err != nil {
		return nil, nil, err
	}

	compilers := loadCompilers()
	err = cache.JSONStore("compilers.json", compilers)
	if err != nil {
		return nil, nil, err
	}
	return runners, compilers, err
}

// Runners returns a list of runners, using a cache
func Runners() RunnersList {
	result, err := cache.JSONCache("runners.json", func() (RunnersList, error) {
		runners := loadRunners()
		return runners, nil
	})
	if err != nil {
		log.Printf("Failed to load runners from cache: %s", err)
	}
	return result
}

// RunnerFor tries to find a suitable runner for a specific file
func RunnerFor(filename string) (Runner, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	for _, runner := range Runners() {
		can := runner.CanRun(filename)
		if can {
			return runner, nil
		}
	}
	return nil, errors.New("no suitable runner found for file")
}

// Compilers returns a list of compilers, using a cache
func Compilers() CompilersList {
	result, err := cache.JSONCache("compilers.json", func() (CompilersList, error) {
		compilers := loadCompilers()
		return compilers, nil
	})
	if err != nil {
		log.Printf("Failed to load compilers from cache: %s", err)
	}
	return result
}

// CompilerFor tries to find a suitable compiler for a specific file
func CompilerFor(filename string) (Compiler, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	for _, compiler := range Compilers() {
		can := compiler.SupportsFile(filename)
		if can {
			return compiler, nil
		}
	}
	return nil, errors.New("no suitable compiler found for file")
}
