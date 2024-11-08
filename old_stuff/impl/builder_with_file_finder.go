package impl

import "github.com/mniak/bench/old_stuff/domain"

type _BuilderWithFileFinder struct {
	domain.Builder
	fileFinder domain.FileFinder
}

func (b *_BuilderWithFileFinder) Build(path string) (string, error) {
	fullpath, err := b.fileFinder.Find(path)
	if err != nil {
		return fullpath, err
	}

	return b.Builder.Build(fullpath)
}

func WrapBuilderWithFileFinder(builder domain.Builder, fileFinder domain.FileFinder) domain.Builder {
	return &_BuilderWithFileFinder{
		Builder:    builder,
		fileFinder: fileFinder,
	}
}
