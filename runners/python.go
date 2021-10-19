package runners

import (
	"bytes"
	"os/exec"
)

func NewPythonLoader() RunnerLoader {
	return &PythonLoader{}
}

type PythonLoader struct{}

func (l *PythonLoader) Load() (Runner, error) {
	var buffer bytes.Buffer
	cmd := exec.Command("python", "--version")
	cmd.Stdout = &buffer

	err := cmd.Run()
	if err != nil {
		return nil, ErrRunnerNotFound
	}

	return newGenericRunner("python"), nil
}

func (l *PythonLoader) CanRun(extension string) bool {
	switch extension {
	case ".py":
		return true
	default:
		return false
	}
}
