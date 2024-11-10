package newcore

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/pkg/errors"
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
	dir, found := g.findRoot(input.Filename)
	if !found {
		return nil, errors.New("could not find project root")
	}

	outFile, err := os.CreateTemp("", "program*.exe")
	if err != nil {
		return nil, err
	}
	defer os.Remove(outFile.Name())

	cmd := exec.Command("go", "build", "-o", outFile.Name(), input.Filename)
	cmd.Dir = dir
	cmd.Stdin = input.Stdin
	cmd.Stdout = input.Stdout
	cmd.Stderr = input.Stderr

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return &Artifact{}, nil
}

func (g *_GoCompiler) SupportsFile(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".py":
		return true
	default:
		return false
	}
}
