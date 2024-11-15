package newcore

import (
	"io"

	"github.com/pkg/errors"
)

type Compiler interface {
	Toolchain
	CanCompile(filename string) bool
	CompilerInputExtensions() []string
	Compile(input CompilationInput) error
}
type CompilationInput struct {
	Filename       string
	OutputFilename string
	Stdin          io.Reader
	Stdout         io.Writer
	Stderr         io.Writer
}

// CompilerFor tries to find a suitable compiler for a specific file
func CompilerFor(filename string) (Compiler, error) {
	for compiler := range iterateToolchains[Compiler] {
		can := compiler.CanCompile(filename)
		if can {
			return compiler, nil
		}
	}
	return nil, errors.New("no suitable compiler found for file")
}
