package impl

import (
	"runtime"

	"github.com/mniak/bench/old_stuff/domain"
	"github.com/mniak/bench/old_stuff/toolchain"
)

var (
	DefaultSourceFinder    domain.FileFinder
	DefaultProgramFinder   domain.FileFinder
	DefaultToolchainFinder domain.ToolchainFinder

	DefaultBuilder domain.Builder
	DefaultTester  domain.Tester
)

func createSourceFinder(toolchainLoaders []domain.ToolchainLoader) domain.FileFinder {
	filenames := []string{"main"}
	extensions := []string{}

	for _, tcl := range toolchainLoaders {
		extensions = append(extensions, tcl.InputExtensions()...)
	}

	return NewFinderOnDirByFilenamesAndExtensions(filenames, extensions)
}

func createProgramFinder() domain.FileFinder {
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
}

func init() {
	toolchainLoaders := []domain.ToolchainLoader{
		toolchain.NewCPPLoader(),
		toolchain.NewGoLoader(),
	}

	DefaultSourceFinder = createSourceFinder(toolchainLoaders)
	DefaultProgramFinder = createProgramFinder()
	DefaultToolchainFinder = NewToolchainFinderFromToolchainLoaders(toolchainLoaders...)

	DefaultBuilder = WrapBuilderWithFileFinder(
		NewBuilder(DefaultToolchainFinder),
		DefaultSourceFinder,
	)

	DefaultTester = WrapTesterWithFileFinder(
		WrapTesterWithBuilder(
			NewTester(),
			WrapBuilderWithSkipWhenNotExist(
				NewBuilder(DefaultToolchainFinder),
			),
		),
		DefaultSourceFinder,
	)
}
