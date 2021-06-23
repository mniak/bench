package bench

type Builder interface {
	Build(path string) (string, error)
}

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

type _BuilderWithProgramFinder struct {
	Builder
	programFinder FileFinder
}

func (b *_BuilderWithProgramFinder) Build(path string) (string, error) {
	fullpath, err := b.programFinder.Find(path)
	if err != nil {
		return "", err
	}

	return b.Builder.Build(fullpath)
}

func NewBuilderWithProgramFinder(builder Builder, sourceFinder FileFinder) Builder {
	return &_BuilderWithProgramFinder{
		Builder:       builder,
		programFinder: sourceFinder,
	}
}
