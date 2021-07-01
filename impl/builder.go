package impl

import "github.com/mniak/bench/domain"

type _BaseBuilder struct {
	toolchainFinder domain.ToolchainFinder
}

func (b *_BaseBuilder) Build(path string) (string, error) {
	tchain, err := b.toolchainFinder.Find(path)
	if err != nil {
		return "", err
	}
	return tchain.Build(path)
}

func NewBuilder(toolchainFinder domain.ToolchainFinder) domain.Builder {
	return &_BaseBuilder{
		toolchainFinder: toolchainFinder,
	}
}
