package toolchain

import "github.com/mniak/bench/domain"

var cppToolchainFactories = make([]ToolchainFactory, 0)

func NewCPP() (domain.Toolchain, error) {
	for _, factory := range cppToolchainFactories {
		tc, err := factory()
		if err == nil {
			return tc, nil
		}
		if err == ErrToolchainNotFound {
			continue
		}
	}
	return nil, ErrToolchainNotFound
}
