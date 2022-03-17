package runners

import (
	"os/exec"
)

func NewBinaryLoader() RunnerLoader {
	return &PythonLoader{}
}

type (
	BinaryLoader  struct{}
	_BinaryRunner struct{}
)

func (r *_BinaryRunner) Run(runnerCmd RunnerCmd) error {
	cmd := exec.Command(runnerCmd.Path)
	cmd.Stdin = runnerCmd.Stdin
	cmd.Stdout = runnerCmd.Stdout
	cmd.Stderr = runnerCmd.Stderr
	return cmd.Run()
}
