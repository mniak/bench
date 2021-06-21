package toolchain

import "errors"

type (
	Toolchain interface {
		Build(mainfile string) (string, error)
	}

	ToolchainFactory func() (Toolchain, error)
)

var ErrToolchainNotFound = errors.New("toolchain was not found")
