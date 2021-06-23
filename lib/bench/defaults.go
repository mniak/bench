package bench

import "runtime"

var DefaultBuilder Builder = &_BuilderWithFileFinder{
	Builder: &_Builder{
		toolchainProducer: new(_ToolchainProducer),
	},
	fileFinder: new(_FileFinder),
}

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}

var DefaultTester Tester = NewTester(defaultProgramFinder)

func StartTest(t Test) (started StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started StartedTest) (result TestResult, err error) {
	return DefaultTester.Wait(started)
}

var defaultProgramFinder *_FileFinder = &_FileFinder{
	filenames:  []string{"main"},
	extensions: []string{".py"},
}
var DefaultProgramFinder FileFinder = defaultProgramFinder

func init() {
	switch runtime.GOOS {
	case "windows":
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".exe")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".bat")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".cmd")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".ps1")
	default:
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, "")
		defaultProgramFinder.extensions = append(defaultProgramFinder.extensions, ".sh")
	}
}
