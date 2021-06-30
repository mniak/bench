package bench

import (
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/impl"
	"github.com/mniak/bench/toolchain"
)

var DefaultToolchainFinder domain.ToolchainFinder = impl.NewToolchainFinderFromToolchains(getToolchainsSkippingErrors())

func getToolchainsSkippingErrors() []domain.Toolchain {
	factories := []func() (domain.Toolchain, error){
		toolchain.NewCPP,
	}

	result := make([]domain.Toolchain, 0)
	for _, fac := range factories {
		tchain, err := fac()
		if err != nil {
			continue
		}
		result = append(result, tchain)
	}
	return result
}
