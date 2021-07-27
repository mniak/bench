package toolchain

import "github.com/mniak/bench/domain"

type _LoaderFromFactories struct {
	factories       []domain.ToolchainFactory
	inputExtensions []string
	outputExtension string
}

// NewLoaderFromFactories creates a new toolchain loader from a list of toolchain factories
func NewLoaderFromFactories(factories []domain.ToolchainFactory, inputExtensions []string, outputExtension string) domain.ToolchainLoader {
	return &_LoaderFromFactories{
		factories:       factories,
		inputExtensions: inputExtensions,
		outputExtension: outputExtension,
	}
}

func (l *_LoaderFromFactories) Load() (domain.Toolchain, error) {
	for _, loader := range l.factories {
		tc, err := loader()
		if err == nil {
			return tc, nil
		}
		if err == ErrToolchainNotFound {
			continue
		}
	}
	return nil, ErrToolchainNotFound
}

func (l *_LoaderFromFactories) InputExtensions() []string {
	return l.inputExtensions
}

func (l *_LoaderFromFactories) OutputExtension() string {
	return l.outputExtension
}
