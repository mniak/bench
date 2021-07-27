package impl

import (
	"runtime"

	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/toolchain"
)

var DefaultBuilder = WrapBuilderWithFileFinder(
	NewBuilder(DefaultToolchainFinder),
	DefaultSourceFinder,
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

var DefaultSourceFinder = func() domain.FileFinder {
	filenames := []string{"main"}
	extensions := []string{}

	var toolchainLoaders []domain.ToolchainLoader
	for _, tcl := range toolchainLoaders {
		extensions = append(extensions, tcl.InputExtensions()...)
	}

	return NewFinderOnDirByFilenamesAndExtensions(filenames, extensions)
}()

var DefaultTester = WrapTesterWithFileFinder(
	WrapTesterWithBuilder(
		NewTester(),
		WrapBuilderWithSkipWhenNotExist(
			NewBuilder(DefaultToolchainFinder),
		),
	),
	DefaultProgramFinder,
)

var DefaultToolchainFinder domain.ToolchainFinder = NewToolchainFinderFromToolchainLoaders(toolchainLoaders...)

var toolchainLoaders []domain.ToolchainLoader

func init() {
	toolchainLoaders = []domain.ToolchainLoader{
		toolchain.NewCPPLoader(),
	}
}
