package runners

type Loader interface {
	Name() string
	LoadRunner() (Runner, error)
}
