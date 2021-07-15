package toolchain

import "github.com/mniak/bench/domain"

func MultiFactory(loaders []domain.ToolchainLoader) domain.ToolchainFactory {
	result := func() (domain.Toolchain, error) {
		return multiFactoryImpl(loaders)
	}
	return result
}

func multiFactoryImpl(loaders []domain.ToolchainLoader) (domain.Toolchain, error) {
	for _, loader := range loaders {
		tc, err := loader.Load()
		if err == nil {
			return tc, nil
		}
		if err == ErrToolchainNotFound {
			continue
		}
	}
	return nil, ErrToolchainNotFound
}
