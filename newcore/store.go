package newcore

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"reflect"

	"github.com/mniak/bench/pkg/cache"
	"github.com/pkg/errors"
)

var loaders = []Loader{
	NewPythonLoader(),
	NewBinaryLoader(),
}

var knownRunnersMap = func() map[string]reflect.Type {
	knownRunners := []Runner{
		new(_PythonRunner),
		new(BinaryRunner),
	}
	result := make(map[string]reflect.Type)
	for _, runner := range knownRunners {
		result[runner.Name()] = reflect.TypeOf(runner).Elem()
	}
	return result
}()

type RunnersList []Runner

func (list RunnersList) MarshalJSON() ([]byte, error) {
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

func (list *RunnersList) UnmarshalJSON(b []byte) error {
	jsonList := make([]struct {
		Kind      string          `json:"kind"`
		RawParams json.RawMessage `json:"params"`
	}, 0)
	if err := json.Unmarshal(b, &jsonList); err != nil {
		return err
	}

	for _, jsonRunner := range jsonList {
		runnerType, found := knownRunnersMap[jsonRunner.Kind]
		if !found {
			log.Printf("Runner kind %q not found", jsonRunner.Kind)
			continue
		}

		runner := reflect.New(runnerType).Interface().(Runner)
		if err := json.Unmarshal(jsonRunner.RawParams, runner); err != nil {
			return err
		}
		*list = append(*list, runner)
	}
	return nil
}

func loadRunners() RunnersList {
	var runners RunnersList
	for _, loader := range loaders {
		r, err := loader.LoadRunner()
		if err != nil {
			log.Printf("Failed to load %q", loader.Name())
			continue
		}
		runners = append(runners, r)
	}
	return runners
}

func RebuildCache() (RunnersList, error) {
	runners := loadRunners()
	err := cache.JSONStore("runners.json", runners)
	return runners, err
}

type _StartedRunnerCmd struct {
	cmd *exec.Cmd
}

func newStartedRunnerCmd(cmd *exec.Cmd) *_StartedRunnerCmd {
	return &_StartedRunnerCmd{
		cmd: cmd,
	}
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
