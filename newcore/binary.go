package newcore

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/exp/slices"
)

func NewBinaryLoader() *BinaryLoader {
	return &BinaryLoader{}
}

type BinaryLoader struct{}

func (bin *BinaryLoader) Name() string {
	return "Binary"
}

func (bin *BinaryLoader) LoadRunner() (Runner, error) {
	var runner BinaryRunner
	if runtime.GOOS == "windows" {
		pathext := os.Getenv("PATHEXT")
		extensions := strings.Split(strings.ToUpper(pathext), ";")
		runner.Extensions = extensions
	}
	return &runner, nil
}

type BinaryRunner struct {
	Extensions []string
}

func (bin *BinaryRunner) Name() string {
	return "Binary"
}

func (bin *BinaryRunner) CanRun(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	if runtime.GOOS == "windows" {
		ext := filepath.Ext(filename)
		return slices.Contains(bin.Extensions, strings.ToUpper(ext))
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
