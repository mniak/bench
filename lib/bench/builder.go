package bench

type Builder interface {
	Build(fullpath string) (string, error)
}

type _Builder struct {
	toolchainProducer ToolchainProducer
}

func (b *_Builder) Build(fullpath string) (string, error) {
	// mainfile, err := b.fileFinder.Find(fullpath)
	// if err != nil {
	// 	return "", err
	// }
	tchain, err := b.toolchainProducer.Produce(fullpath)
	if err != nil {
		return "", err
	}
	return tchain.Build(fullpath)
}

type _BuilderWithFileFinder struct {
	Builder
	fileFinder FileFinder
}
