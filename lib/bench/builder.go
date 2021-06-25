package bench

type _Builder struct {
	toolchainProducer ToolchainProducer
}

func (b *_Builder) Build(path string) (string, error) {
	tchain, err := b.toolchainProducer.Produce(path)
	if err != nil {
		return "", err
	}
	return tchain.Build(path)
}

func NewBuilder(toolchainProducer ToolchainProducer) Builder {
	return &_Builder{
		toolchainProducer: toolchainProducer,
	}
}
