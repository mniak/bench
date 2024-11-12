package bench

import (
	"github.com/mniak/bench/old_stuff/impl"
)

var DefaultTester = impl.DefaultTester

func StartTest(t Test) (started StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started StartedTest) (result TestResult, err error) {
	return DefaultTester.Wait(started)
}
