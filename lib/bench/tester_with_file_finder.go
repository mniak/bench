package bench

type _TesterWithFileFinder struct {
	Tester
	fileFinder FileFinder
}

func (t *_TesterWithFileFinder) Start(test Test) (started StartedTest, err error) {
	test.Program, err = t.fileFinder.Find(test.Program)
	if err != nil {
		return
	}
	return t.Tester.Start(test)
}

func WrapTesterWithFileFinder(tester Tester, fileFinder FileFinder) Tester {
	return &_TesterWithFileFinder{
		Tester:     tester,
		fileFinder: fileFinder,
	}
}
