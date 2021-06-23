package bench

import "runtime"

var DefaultProgramFinder FileFinder = &_FileFinder{
	filenames:  []string{"main"},
	extensions: []string{".py"},
}

var DefaultBuilder Builder = WrapBuilderWithProgramFinder(
	NewBuilder(new(_ToolchainProducer)),
	DefaultProgramFinder,
)
var DefaultTester Tester = NewTester(DefaultProgramFinder)

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}

func StartTest(t Test) (started StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started StartedTest) (result TestResult, err error) {
	return DefaultTester.Wait(started)
}

func init() {
	programFinder := DefaultProgramFinder.(*_FileFinder)
	switch runtime.GOOS {
	case "windows":
		programFinder.extensions = append(programFinder.extensions, ".exe")
		programFinder.extensions = append(programFinder.extensions, ".bat")
		programFinder.extensions = append(programFinder.extensions, ".cmd")
		programFinder.extensions = append(programFinder.extensions, ".ps1")
	default:
		programFinder.extensions = append(programFinder.extensions, "")
		programFinder.extensions = append(programFinder.extensions, ".sh")
	}
}
