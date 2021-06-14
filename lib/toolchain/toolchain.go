package toolchain

import "errors"

type Toolchain interface {
	Build(mainfile string) (string, error)
}

var (
	ErrToolchainOSUnsupported = errors.New("toolchain is not supported in this OS")
	ErrToolchainNotFound      = errors.New("toolchain was not found")
)
