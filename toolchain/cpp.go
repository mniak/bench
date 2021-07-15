package toolchain

import "github.com/mniak/bench/domain"

var cppToolchainLoaders = make([]domain.ToolchainLoader, 0)

func NewCPPFactory() domain.ToolchainFactory {
	return MultiFactory(cppToolchainLoaders)
}
