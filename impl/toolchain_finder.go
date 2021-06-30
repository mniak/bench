package impl

import (
	"os"
	"path/filepath"

	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/internal/utils"
	"github.com/mniak/bench/toolchain"
)

type _ToolchainFinder struct {
	toolchains []domain.Toolchain
}

func (tp *_ToolchainFinder) Produce(filename string) (domain.Toolchain, error) {
	ext := filepath.Ext(filename)
	for _, tchain := range tp.toolchains {
		if tchain.OutputExtension() == ext {
			for _, inExt := range tchain.InputExtensions() {
				if _, err := os.Stat(utils.ChangeExtension(filename, inExt)); err == nil {
					return tchain, nil
				}
			}
		}

		// for _, inext  := range tchain.InputExtensions() {
		// 	if ext == inext {

		// 	}
		// }
	}
	// if factory, ok := tp.toolchains[ext]; ok {
	// 	return factory()
	// }
	return nil, toolchain.ErrToolchainNotFound
}

func NewToolchainFinderFromToolchains(toolchains []domain.Toolchain) domain.ToolchainFinder {
	return &_ToolchainFinder{
		toolchains: toolchains,
	}
}
