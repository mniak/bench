package newcore

import (
	"io"
	"os"
	"reflect"
)

type CompilerLoader interface {
	Name() string
	LoadCompiler() (Compiler, error)
	CompilerType() reflect.Type
}
type Compiler interface {
	Name() string
	SupportsFile(filename string) bool
	Compile(input CompilerInput) (*Artifact, error)
}
type CompilerInput struct {
	Filename       string
	OutputFilename string
	Stdin          io.Reader
	Stdout         io.Writer
	Stderr         io.Writer
}

type Artifact struct {
	OutputFilename string
}

func (a *Artifact) Filename() string {
	return a.OutputFilename
}

func (a *Artifact) Free() error {
	return os.RemoveAll(a.OutputFilename)
}
