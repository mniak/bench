package bench

import (
	"runtime"

	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/impl"
)

var DefaultProgramFinder domain.FileFinder = defaultFileFinder()

func defaultFileFinder() domain.FileFinder {
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

	return impl.NewFinderOnDirByFilenamesAndExtensions(filenames, extensions)
}
