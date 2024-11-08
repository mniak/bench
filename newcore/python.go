package newcore

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

func NewPythonLoader() Loader {
	return &_PythonLoader{}
}

type _PythonLoader struct{}

func (py *_PythonLoader) Name() string {
	return "Python"
}

func (py *_PythonLoader) LoadRunner() (Runner, error) {
	var buffer bytes.Buffer
	cmd := exec.Command("python", "--version")
	cmd.Stdout = &buffer

	err := cmd.Run()
	if err != nil {
		return nil, errors.New("runner not loaded: python not found")
	}

	return new(_PythonRunner), nil
}

type _PythonRunner struct{}

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
