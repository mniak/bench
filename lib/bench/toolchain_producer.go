package bench

import (
	"path/filepath"

	"github.com/mniak/bench/lib/toolchain"
)

type _ToolchainProducer struct{}

func (tp *_ToolchainProducer) Produce(mainfile string) (toolchain.Toolchain, error) {
	switch filepath.Ext(mainfile) {
	case ".cpp", ".c++":
		return toolchain.NewCPP()
	}
	return nil, toolchain.ErrToolchainNotFound
}
