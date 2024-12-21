package newcore

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/exp/slices"
)

type _WindowsBinaryRunner struct {
	Extensions []string
}

func loadPlatformSpecificBinaryRunner() _WindowsBinaryRunner {
	pathext := os.Getenv("PATHEXT")
	extensions := strings.Split(strings.ToUpper(pathext), ";")
	return _WindowsBinaryRunner{
		Extensions: extensions,
	}
}

func (bin _WindowsBinaryRunner) CanRun(filename string) bool {
	if len(bin.Extensions) == 0 {
		return false
	}
	ext := filepath.Ext(filename)
	hasExt := slices.Contains(bin.Extensions, strings.ToUpper(ext))
	if !hasExt {
		return false
	}

	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}
	return true
}

func (bin _WindowsBinaryRunner) Start(program string, a RunArgs) (Waiter, error) {
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

func (bin _WindowsBinaryRunner) RunnerInputExtensions() []string {
	return bin.Extensions
}
