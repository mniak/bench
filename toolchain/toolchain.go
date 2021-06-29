package toolchain

import (
	"errors"

	"github.com/mniak/bench/domain"
)

type (
	ToolchainFactory func() (domain.Toolchain, error)
)

var ErrToolchainNotFound = errors.New("toolchain was not found")
