package oldall

import (
	"io"
	"reflect"
)

type ToolchainFinder interface {
	Find(string) (ToolchainLoader, error)
}

type ToolchainLoader interface {
	Name() string
	Load() (Toolchain, error)
	InputExtensions() []string
	OutputExtension() string
	ToolchainType() reflect.Type
}

type BuildRequest struct {
	Stdout io.Writer
	Stderr io.Writer
	Input  string
	Output string
}

type Toolchain interface {
	Name() string
	Build(request BuildRequest) error
}

type ToolchainFactory func() (Toolchain, error)
