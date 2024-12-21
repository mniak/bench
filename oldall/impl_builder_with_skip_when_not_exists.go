package oldall

import (
	"os"
)

type _BuilderWithSkipWhenNotExist struct {
	Builder
}

func (b *_BuilderWithSkipWhenNotExist) Build(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return filename, nil
	}
	return b.Builder.Build(filename)
}

func WrapBuilderWithSkipWhenNotExist(builder Builder) Builder {
	return &_BuilderWithSkipWhenNotExist{
		Builder: builder,
	}
}
