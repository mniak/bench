package toolchain

import "github.com/mniak/bench/domain"

var cppToolchainFactories = make([]domain.ToolchainFactory, 0)

func NewCPPLoader() domain.ToolchainLoader {
	return NewLoaderFromFactories(
		cppToolchainFactories,
		[]string{".cpp", ".cxx", ".c++"},
		domain.OSBinaryExtension,
	)
}
