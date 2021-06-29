package impl

import "github.com/mniak/bench/domain"

type _Builder struct {
	toolchainProducer domain.ToolchainProducer
}

func (b *_Builder) Build(path string) (string, error) {
	tchain, err := b.toolchainProducer.Produce(path)
	if err != nil {
		return "", err
	}
	return tchain.Build(path)
}

func NewBuilder(toolchainProducer domain.ToolchainProducer) domain.Builder {
	return &_Builder{
		toolchainProducer: toolchainProducer,
	}
}
