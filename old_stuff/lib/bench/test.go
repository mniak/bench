package bench

import (
	"github.com/mniak/bench/old_stuff/domain"
	"github.com/mniak/bench/old_stuff/impl"
)

var DefaultTester = impl.DefaultTester

func StartTest(t domain.Test) (started domain.StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started domain.StartedTest) (result domain.TestResult, err error) {
	return DefaultTester.Wait(started)
}
