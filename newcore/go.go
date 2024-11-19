package newcore

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/pkg/errors"
)

type GoLoader struct{}

func (l *GoLoader) Load() (Toolchain, error) {
	var toolchain _GoToolchain
	err := testProgram("go", "version")
	if err == nil {
		return &toolchain, nil
	}
	return nil, errors.New("toolchain not loaded: Go not found")
}

func (l *GoLoader) ToolchainType() reflect.Type {
	return reflect.TypeOf(_GoToolchain{})
}

type _GoToolchain struct{}

func (g *_GoToolchain) findRoot(filename string) (string, bool) {
	info, err := os.Stat(filename)
	if err != nil {
		return "", false
	}
	if info.IsDir() {
		return "", false
	}
	return "", false
}

func (g *_GoToolchain) Compile(input CompilationInput) error {
	newWorkingDir := filepath.Dir(input.Filename)

	inputFilename, err := filepath.Rel(newWorkingDir, input.Filename)
	if err != nil {
		return err
	}

	outputFilename, err := filepath.Abs(input.OutputFilename)
	if err != nil {
		return err
	}

	cmd := exec.Command(
		"go", "build",
		"-o", outputFilename,
		inputFilename,
	)
	cmd.Dir = newWorkingDir
	cmd.Stdin = input.Stdin
	cmd.Stdout = input.Stdout
	cmd.Stderr = input.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (g *_GoToolchain) CanCompile(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".go":
		return true
	default:
		return false
	}
}

func (g *_GoToolchain) CompilerInputExtensions() []string {
	return []string{".go"}
}
