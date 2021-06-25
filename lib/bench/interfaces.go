package bench

import "github.com/mniak/bench/lib/toolchain"

type ToolchainProducer interface {
	Produce(string) (toolchain.Toolchain, error)
}

type FileFinder interface {
	Find(filename string) (string, error)
}

type Builder interface {
	Build(path string) (string, error)
}

type Tester interface {
	Start(t Test) (started StartedTest, err error)
	Wait(started StartedTest) (result TestResult, err error)
}
