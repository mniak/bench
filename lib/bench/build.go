package bench

import (
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/impl"
)

var DefaultBuilder domain.Builder = impl.DefaultBuilder

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}
