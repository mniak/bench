package impl

import "github.com/mniak/bench/domain"

type _TesterWithFileFinder struct {
	domain.Tester
	fileFinder domain.FileFinder
}

func (t *_TesterWithFileFinder) Start(test domain.Test) (started domain.StartedTest, err error) {
	test.Program, err = t.fileFinder.Find(test.Program)
	if err != nil {
		return
	}
	return t.Tester.Start(test)
}

func WrapTesterWithFileFinder(tester domain.Tester, fileFinder domain.FileFinder) domain.Tester {
	return &_TesterWithFileFinder{
		Tester:     tester,
		fileFinder: fileFinder,
	}
}
