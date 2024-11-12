package bench

import (
	"github.com/mniak/bench/old_stuff/impl"
)

var DefaultBuilder Builder = impl.DefaultBuilder

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}
