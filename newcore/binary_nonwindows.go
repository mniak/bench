//go:build !windows

package newcore

import (
	"os"
	"os/exec"
)

type _NonWindowsBinaryRunner struct{}

func loadPlatformSpecificBinaryRunner() _NonWindowsBinaryRunner {
	return _NonWindowsBinaryRunner{}
}

func (bin _NonWindowsBinaryRunner) CanRun(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}

	executable := info.Mode()&0o100 != 0
	return executable
}

func (bin _NonWindowsBinaryRunner) Start(program string, a RunArgs) (StartedCmd, error) {
	c := exec.Command(program, a.Args...)
	c.Stdin = a.Stdin
	c.Stdout = a.Stdout
	c.Stderr = a.Stderr
	err := c.Start()
	if err != nil {
		return nil, err
	}
	return newStartedRunnerCmd(c), nil
}

func (bin _NonWindowsBinaryRunner) RunnerInputExtensions() []string {
	return nil
}
