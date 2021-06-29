package bench

import (
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/impl"
)

var DefaultBuilder domain.Builder = impl.WrapBuilderWithSourceFinder(
	impl.NewBuilder(impl.NewToolchainProducer()),
	DefaultProgramFinder,
)

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}
