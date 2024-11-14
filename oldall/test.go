package oldall

var DefaultTester = oldall.DefaultTester

func FindTest()

func StartTest(t Test) (started StartedTest, err error) {
	return DefaultTester.Start(t)
}

func WaitTest(started StartedTest) (result TestResult, err error) {
	return DefaultTester.Wait(started)
}
