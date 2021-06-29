package toolchain

import (
	"github.com/mniak/bench/domain"
)

type ToolchainFactory func() (domain.Toolchain, error)
