package runners

import (
	"bytes"
	"os/exec"
	"path/filepath"
)

func NewPythonRunner() Loader {
	return &PythonRunner{}
}

type PythonRunner struct{}

func (py *PythonRunner) Name() string {
	return "Python"
}

func (py *PythonRunner) Load() (Runner, error) {
	var buffer bytes.Buffer
	cmd := exec.Command("python", "--version")
	cmd.Stdout = &buffer

	err := cmd.Run()
	if err != nil {
		return nil, ErrRunnerNotFound
	}

	return py, nil
}

func (py *PythonRunner) CanRun(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".py":
		return true
	default:
		return false
	}
}

func (py *PythonRunner) Start(cmd Cmd) (StartedCmd, error) {
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
