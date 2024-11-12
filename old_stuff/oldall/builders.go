package oldall

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
	"reflect"

	"github.com/mniak/bench/pkg/cache"
)

var loaders = []Loader{
	NewGoBuilder(),
}

var knownBuildersMap = func() map[string]reflect.Type {
	knownBuilders := []Builder{
		new(GoBuilder),
	}
	result := make(map[string]reflect.Type)
	for _, builder := range knownBuilders {
		result[builder.Name()] = reflect.TypeOf(builder).Elem()
	}
	return result
}()

type BuilderList []Builder

func (list BuilderList) MarshalJSON() ([]byte, error) {
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

func (list *BuilderList) UnmarshalJSON(b []byte) error {
	jsonList := make([]struct {
		Kind      string          `json:"kind"`
		RawParams json.RawMessage `json:"params"`
	}, 0)
	if err := json.Unmarshal(b, &jsonList); err != nil {
		return err
	}

	for _, jsonBuilder := range jsonList {
		builderType, found := knownBuildersMap[jsonBuilder.Kind]
		if !found {
			log.Printf("Builder kind %q not found", jsonBuilder.Kind)
			continue
		}

		builder := reflect.New(builderType).Interface().(Builder)
		if err := json.Unmarshal(jsonBuilder.RawParams, builder); err != nil {
			return err
		}
		*list = append(*list, builder)
	}
	return nil
}

func loadBuilders() BuilderList {
	var builders BuilderList
	for _, loader := range loaders {
		r, err := loader.Load()
		if err != nil {
			log.Printf("Failed to load %q", loader.Name())
			continue
		}
		builders = append(builders, r)
	}
	return builders
}

func RebuildCache() error {
	builders := loadBuilders()
	err := cache.JSONStore("builders.json", builders)
	return err
}

func Builders() BuilderList {
	result, err := cache.JSONCache("builders.json", func() (BuilderList, error) {
		builders := loadBuilders()
		return builders, nil
	})
	if err != nil {
		log.Printf("Failed to load builders from cache: %s", err)
	}
	return result
}

func BuilderFor(filename string) (Builder, error) {
	for _, builder := range Builders() {
		can := builder.CanRun(filename)
		if !can {
			continue
		}
		return builder, nil
	}
	return nil, ErrBuilderNotFound
}

type Loader interface {
	Name() string
	Load() (Builder, error)
}

type _StartedBuilderCmd struct {
	cmd *exec.Cmd
}

func newStartedBuilderCmd(cmd *exec.Cmd) *_StartedBuilderCmd {
	return &_StartedBuilderCmd{
		cmd: cmd,
	}
}

type Builder interface {
	Name() string
	CanRun(filename string) bool
	Start(cmd Cmd) (StartedCmd, error)
}

func StartAndWait(r Builder, cmd Cmd) error {
	startedCmd, err := r.Start(cmd)
	if err != nil {
		return err
	}
	return startedCmd.Wait()
}

type StartedCmd interface {
	Wait() error
}

func (c *_StartedBuilderCmd) Wait() error {
	return c.cmd.Wait()
}

type Cmd struct {
	Path   string
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
