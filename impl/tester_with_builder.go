package impl

import "github.com/mniak/bench/domain"

type _TesterWithBuilder struct {
	domain.Tester
	builder      domain.Builder
	sourceFinder domain.FileFinder
}

func (t *_TesterWithBuilder) Start(test domain.Test) (started domain.StartedTest, err error) {
	program, err := t.builder.Build(test.Program)
	if err != nil {
		return nil, err
	}
	test.Program = program
	return t.Tester.Start(test)
}

func WrapTesterWithBuilder(tester domain.Tester, builder domain.Builder) domain.Tester {
	return &_TesterWithBuilder{
		Tester:  tester,
		builder: builder,
	}
}
