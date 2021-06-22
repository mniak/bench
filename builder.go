package bench

type Builder interface {
	Build(string) (string, error)
}

type _Builder struct {
	programFinder     ProgramFinder
	toolchainProducer ToolchainProducer
}

func (b *_Builder) Build(path string) (string, error) {
	mainfile, err := b.programFinder.Find(path)
	if err != nil {
		return "", err
	}
	tchain, err := b.toolchainProducer.Produce(mainfile)
	if err != nil {
		return "", err
	}
	return tchain.Build(mainfile)
}

var DefaultBuilder Builder = &_Builder{
	toolchainProducer: new(_ToolchainProducer),
	programFinder:     new(_ProgramFinder),
}

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}
