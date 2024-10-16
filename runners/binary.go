package runners

import (
	"os"
	"os/exec"
)

func NewBinaryRunner() *BinaryRunner {
	return &BinaryRunner{}
}

type BinaryRunner struct{}

func (bin *BinaryRunner) Name() string {
	return "Binary"
}

func (bin *BinaryRunner) Load() (Runner, error) {
	return bin, nil
}

func (bin *BinaryRunner) CanRun(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	executable := info.Mode()&0o100 != 0
	return executable
}

func (bin *BinaryRunner) Start(cmd Cmd) (StartedCmd, error) {
	c := exec.Command(cmd.Path, cmd.Args...)
	c.Stdin = cmd.Stdin
	c.Stdout = cmd.Stdout
	c.Stderr = cmd.Stderr
	err := c.Start()
	if err != nil {
		return nil, err
	}
	return newStartedRunnerCmd(c), nil
}
