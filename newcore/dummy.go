package newcore

func AlwaysLoader(r Runner) Loader {
	return &_AlwaysLoader{r}
}

type _AlwaysLoader struct {
	Runner
}

func (l *_AlwaysLoader) Name() string {
	return l.Runner.Name()
}

func (l *_AlwaysLoader) LoadRunner() (Runner, error) {
	return l.Runner, nil
}
