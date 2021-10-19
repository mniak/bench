package runner

import (
	"os/exec"

	"github.com/mniak/bench/domain"
)

func NewBinaryLoader() domain.RunnerLoader {
	return &PythonLoader{}
}

type (
	BinaryLoader  struct{}
	_BinaryRunner struct{}
)

func (r *_BinaryRunner) Run(runnerCmd domain.RunnerCmd) error {
	cmd := exec.Command(runnerCmd.Path)
	cmd.Stdin = runnerCmd.Stdin
	cmd.Stdout = runnerCmd.Stdout
	cmd.Stderr = runnerCmd.Stderr
	return cmd.Run()
}
