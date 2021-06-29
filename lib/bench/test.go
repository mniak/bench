package bench

import (
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/impl"
)

var DefaultTester domain.Tester = impl.WrapTesterWithFileFinder(
	impl.NewTester(),
	DefaultProgramFinder,
)

func StartTest(t domain.Test) (started domain.StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started domain.StartedTest) (result domain.TestResult, err error) {
	return DefaultTester.Wait(started)
}
