package impl

import (
	"runtime"

	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/toolchain"
)

var DefaultBuilder = WrapBuilderWithFileFinder(
	NewBuilder(DefaultToolchainFinder),
	DefaultProgramFinder,
)

var DefaultProgramFinder = func() domain.FileFinder {
	filenames := []string{"main"}
	extensions := []string{".py"}

	switch runtime.GOOS {
	case "windows":
		extensions = append(extensions, ".exe")
		extensions = append(extensions, ".bat")
		extensions = append(extensions, ".cmd")
		extensions = append(extensions, ".ps1")
	default:
		extensions = append(extensions, "")
		extensions = append(extensions, ".sh")
	}

	return NewFinderOnDirByFilenamesAndExtensions(filenames, extensions)
}()

var DefaultTester = WrapTesterWithFileFinder(
	NewTester(),
	DefaultProgramFinder,
)

var DefaultToolchainFinder domain.ToolchainFinder = NewToolchainFinderFromToolchains(getToolchainsSkippingErrors()...)

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
