package impl

import "github.com/mniak/bench/domain"

type _Builder struct {
	toolchainFinder domain.ToolchainFinder
}

func (b *_Builder) Build(path string) (string, error) {
	tchain, err := b.toolchainFinder.Produce(path)
	if err != nil {
		return "", err
	}
	return tchain.Build(path)
}

func NewBuilder(toolchainFinder domain.ToolchainFinder) domain.Builder {
	return &_Builder{
		toolchainFinder: toolchainFinder,
	}
}
