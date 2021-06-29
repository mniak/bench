package bench

import (
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/impl"
	"github.com/mniak/bench/toolchain"
)

var DefaultToolchainProducer domain.ToolchainProducer = impl.NewToolchainProducerFromExtensionMap(map[string]domain.ToolchainFactory{
	".cpp": toolchain.NewCPP,
	".c++": toolchain.NewCPP,
})
