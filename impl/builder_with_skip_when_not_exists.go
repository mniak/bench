package impl

import (
	"os"

	"github.com/mniak/bench/domain"
)

type _BuilderWithSkipWhenNotExist struct {
	domain.Builder
}

func (b *_BuilderWithSkipWhenNotExist) Build(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return filename, nil
	}
	return b.Builder.Build(filename)
}

func WrapBuilderWithSkipWhenNotExist(builder domain.Builder) domain.Builder {
	return &_BuilderWithSkipWhenNotExist{
		Builder: builder,
	}
}
