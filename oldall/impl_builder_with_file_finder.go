package oldall

type _BuilderWithFileFinder struct {
	Builder
	fileFinder FileFinder
}

func (b *_BuilderWithFileFinder) Build(path string) (string, error) {
	fullpath, err := b.fileFinder.Find(path)
	if err != nil {
		return fullpath, err
	}

	return b.Builder.Build(fullpath)
}

func WrapBuilderWithFileFinder(builder Builder, fileFinder FileFinder) Builder {
	return &_BuilderWithFileFinder{
		Builder:    builder,
		fileFinder: fileFinder,
	}
}
