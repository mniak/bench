package newcore

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/pkg/errors"
)

type _PythonLoader struct{}

func (py *_PythonLoader) Name() string {
	return "Python"
}

func testProgram(program string, moreArgs ...string) error {
	var buffer bytes.Buffer
	cmd := exec.Command(program, moreArgs...)
	cmd.Stdout = &buffer

	err := cmd.Run()
	if err != nil {
		return errors.New("not found")
	}

	return nil
}

func (py *_PythonLoader) LoadRunner() (Runner, error) {
	var runner _PythonRunner
	for _, program := range []string{
		"python3", "python",
	} {
		err := testProgram(program, "--version")
		if err == nil {
			runner.Command = program
			return &runner, nil
		}
	}
	return nil, errors.New("runner not loaded: python not found")
}

func (py *_PythonLoader) RunnerType() reflect.Type {
	return reflect.TypeOf(_PythonRunner{})
}

type _PythonRunner struct {
	Command string
}

func (py *_PythonRunner) Name() string {
	return "Python"
}

func (py *_PythonRunner) CanRun(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".py":
		return true
	default:
		return false
	}
}

func (py *_PythonRunner) Start(cmd Cmd) (StartedCmd, error) {
	args := append([]string{cmd.Path}, cmd.Args...)
	c := exec.Command("python", args...)
	c.Stdin = cmd.Stdin
	c.Stdout = cmd.Stdout
	c.Stderr = cmd.Stderr
	err := c.Start()
	if err != nil {
		return nil, err
	}
	return newStartedRunnerCmd(c), nil
}
