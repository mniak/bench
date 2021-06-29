package impl

import (
	"path/filepath"

	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/toolchain"
)

type _ToolchainProducer struct{}

func (tp *_ToolchainProducer) Produce(mainfile string) (domain.Toolchain, error) {
	switch filepath.Ext(mainfile) {
	case ".cpp", ".c++":
		return toolchain.NewCPP()
	}
	return nil, toolchain.ErrToolchainNotFound
}

func NewToolchainProducer() domain.ToolchainProducer {
	return &_ToolchainProducer{}
}
