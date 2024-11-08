package domain

import "io"

type ToolchainFinder interface {
	Find(string) (ToolchainLoader, error)
}

type ToolchainLoader interface {
	Load() (Toolchain, error)
	InputExtensions() []string
	OutputExtension() string
}

type BuildRequest struct {
	Stdout io.Writer
	Stderr io.Writer
	Input  string
	Output string
}

type Toolchain interface {
	Build(request BuildRequest) error
}

type ToolchainFactory func() (Toolchain, error)
