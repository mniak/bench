package bench

type _TesterWithBuilder struct {
	Tester
	builder      Builder
	sourceFinder FileFinder
}

func (t *_TesterWithBuilder) Start(test Test) (started StartedTest, err error) {
	program, err := t.builder.Build(test.Program)
	if err != nil {
		return StartedTest{}, err
	}
	test.Program = program
	return t.Tester.Start(test)
}

func WrapTesterWithBuilder(tester Tester, builder Builder) Tester {
	return &_TesterWithBuilder{
		Tester:  tester,
		builder: builder,
	}
}
