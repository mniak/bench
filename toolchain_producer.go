package bench

import (
	"path"

	"github.com/mniak/bench/toolchain"
)

type ToolchainProducer interface {
	Produce(string) (toolchain.Toolchain, error)
}

type _ToolchainProducer struct{}

func (tp *_ToolchainProducer) Produce(mainfile string) (toolchain.Toolchain, error) {
	switch path.Ext(mainfile) {
	case ".cpp", ".c++":
		return toolchain.NewCPP()
	}
	return nil, toolchain.ErrToolchainNotFound
}