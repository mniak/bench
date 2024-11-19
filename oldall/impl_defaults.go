package oldall

import (
	"runtime"

	"github.com/mniak/bench/old_stuff/toolchain"
)

var DefaultSourceFinder FileFinder

func createSourceFinder(toolchainLoaders []ToolchainLoader) FileFinder {
	filenames := []string{"main"}
	extensions := []string{}

	for _, tcl := range toolchainLoaders {
		extensions = append(extensions, tcl.InputExtensions()...)
	}

	return NewFinderOnDirByFilenamesAndExtensions(filenames, extensions)
}

func createProgramFinder() FileFinder {
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
	toolchainLoaders := []ToolchainLoader{
		// toolchain.NewCPPLoader(),
		toolchain.GoToolchainLoader{},
	}

	DefaultSourceFinder = createSourceFinder(toolchainLoaders)
	DefaultProgramFinder = createProgramFinder()
	DefaultToolchainFinder = NewToolchainFinderFromToolchainLoaders(toolchainLoaders...)

	// DefaultBuilder = WrapBuilderWithFileFinder(
	// 	NewBuilder(DefaultToolchainFinder),
	// 	DefaultSourceFinder,
	// )

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
