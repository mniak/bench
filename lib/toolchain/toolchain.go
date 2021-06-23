package toolchain

import "errors"

type (
	Toolchain interface {
		Build(mainfile string) (string, error)
		OutputExtension() string
		InputExtensions() []string
	}

	ToolchainFactory func() (Toolchain, error)
)

var ErrToolchainNotFound = errors.New("toolchain was not found")
