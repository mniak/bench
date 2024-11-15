package newcore

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/pkg/errors"
)

type PythonLoader struct{}

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

func (py *PythonLoader) Load() (Toolchain, error) {
	var toolchain _PythonToolchain
	for _, program := range []string{
		"python3", "python",
	} {
		err := testProgram(program, "--version")
		if err == nil {
			toolchain.Command = program
			return &toolchain, nil
		}
	}
	return nil, errors.New("toolchain not loaded: python not found")
}

func (py *PythonLoader) ToolchainType() reflect.Type {
	return reflect.TypeOf(_PythonToolchain{})
}

type _PythonToolchain struct {
	Command string
}

func (py *_PythonToolchain) CanRun(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".py":
		return true
	default:
		return false
	}
}

func (py *_PythonToolchain) Start(programPath string, a RunArgs) (StartedProgram, error) {
	args := append([]string{programPath}, a.Args...)
	c := exec.Command("python", args...)
	c.Stdin = a.Stdin
	c.Stdout = a.Stdout
	c.Stderr = a.Stderr
	err := c.Start()
	if err != nil {
		return nil, err
	}
	return newStartedRunnerCmd(c), nil
}

func (py *_PythonToolchain) RunnerInputExtensions() []string {
	return []string{".py"}
}
