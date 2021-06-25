package bench

import "runtime"

var (
	DefaultProgramFinder FileFinder
	DefaultBuilder       Builder = WrapBuilderWithSourceFinder(
		NewBuilder(new(_ToolchainProducer)),
		DefaultProgramFinder,
	)
	DefaultTester Tester = WrapTesterWithFileFinder(
		NewTester(),
		DefaultProgramFinder,
	)
)

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}

func StartTest(t Test) (started StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started StartedTest) (result TestResult, err error) {
	return DefaultTester.Wait(started)
}

func defaultFileFinder() FileFinder {
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
	DefaultProgramFinder = defaultFileFinder()
}
