package newcore

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

type _GoLoader struct{}

func (l *_GoLoader) Name() string {
	return "Go"
}

func (l *_GoLoader) LoadCompiler() (Compiler, error) {
	return &_GoCompiler{}, nil
}

func (l *_GoLoader) CompilerType() reflect.Type {
	return reflect.TypeOf(_GoCompiler{})
}

type _GoCompiler struct{}

func (g *_GoCompiler) Name() string {
	return "Go"
}

func (g *_GoCompiler) findRoot(filename string) (string, bool) {
	info, err := os.Stat(filename)
	if err != nil {
		return "", false
	}
	if info.IsDir() {
		return "", false
	}
	return "", false
}

func (g *_GoCompiler) Compile(input CompilerInput) (*Artifact, error) {
	// dir, found := g.findRoot(input.Filename)
	// if !found {
	// 	return nil, errors.New("could not find project root")
	// }

	newWorkingDir := filepath.Dir(input.Filename)

	inputFilename, err := filepath.Rel(newWorkingDir, input.Filename)
	if err != nil {
		return nil, err
	}

	outputFilename, err := filepath.Abs(input.OutputFilename)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(
		"go", "build",
		"-o", outputFilename,
		inputFilename,
	)
	// cmd.Dir = dir
	cmd.Dir = newWorkingDir
	cmd.Stdin = input.Stdin
	cmd.Stdout = input.Stdout
	cmd.Stderr = input.Stderr

	fmt.Println("Go Command:", strings.Join(cmd.Args, " "))

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return &Artifact{}, nil
}

func (g *_GoCompiler) SupportsFile(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".go":
		return true
	default:
		return false
	}
}
