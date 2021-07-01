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

func (tp *_ToolchainFinder) Find(filename string) (domain.Toolchain, error) {
	ext := filepath.Ext(filename)
	for _, tchain := range tp.toolchains {
		inputExtensions := tchain.InputExtensions()

		if tchain.OutputExtension() == ext {
			for _, inExt := range inputExtensions {
				if _, err := os.Stat(utils.ChangeExtension(filename, inExt)); err == nil {
					return tchain, nil
				}
			}
		}

		for _, inExt := range inputExtensions {
			if inExt == ext {
				if _, err := os.Stat(filename); err == nil {
					return tchain, nil
				}
			}
		}
	}
	return nil, toolchain.ErrToolchainNotFound
}

func NewToolchainFinderFromToolchains(toolchains ...domain.Toolchain) domain.ToolchainFinder {
	return &_ToolchainFinder{
		toolchains: toolchains,
	}
}
