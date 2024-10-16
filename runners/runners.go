package runners

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
	"reflect"

	"github.com/mniak/bench/cache"
)

var loaders = []Loader{
	NewPythonRunner(),
	NewBinaryRunner(),
}

var knownRunnersMap = func() map[string]reflect.Type {
	knownRunners := []Runner{
		new(PythonRunner),
		new(BinaryRunner),
	}
	result := make(map[string]reflect.Type)
	for _, runner := range knownRunners {
		result[runner.Name()] = reflect.TypeOf(runner).Elem()
	}
	return result
}()

type RunnerList []Runner

func (list RunnerList) MarshalJSON() ([]byte, error) {
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

func (list *RunnerList) UnmarshalJSON(b []byte) error {
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
			log.Println("Runner kind %q not found")
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

func loadRunners() RunnerList {
	var runners RunnerList
	for _, loader := range loaders {
		r, err := loader.Load()
		if err != nil {
			log.Printf("Failed to load %q", loader.Name())
			continue
		}
		runners = append(runners, r)
	}
	return runners
}

func RebuildCache() error {
	runners := loadRunners()
	err := cache.JSONStore("runners.json", runners)
	return err
}

func Runners() RunnerList {
	result, err := cache.JSONCache("runners.json", func() (RunnerList, error) {
		runners := loadRunners()
		return runners, nil
	})
	if err != nil {
		log.Printf("Failed to load runners from cache: %s", err)
	}
	return result
}

func RunnerFor(filename string) (Runner, error) {
	for _, runner := range Runners() {
		can := runner.CanRun(filename)
		if !can {
			continue
		}
		return runner, nil
	}
	return nil, ErrRunnerNotFound
}

type Loader interface {
	Name() string
	Load() (Runner, error)
}

type _StartedRunnerCmd struct {
	cmd *exec.Cmd
}

func newStartedRunnerCmd(cmd *exec.Cmd) *_StartedRunnerCmd {
	return &_StartedRunnerCmd{
		cmd: cmd,
	}
}

type Runner interface {
	Name() string
	CanRun(filename string) bool
	Start(cmd Cmd) (StartedCmd, error)
}

func StartAndWait(r Runner, cmd Cmd) error {
	startedCmd, err := r.Start(cmd)
	if err != nil {
		return err
	}
	return startedCmd.Wait()
}

type StartedCmd interface {
	Wait() error
}

func (c *_StartedRunnerCmd) Wait() error {
	return c.cmd.Wait()
}

type Cmd struct {
	Path   string
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
