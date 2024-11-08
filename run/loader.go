package run

type Loader interface {
	Name() string
	LoadRunner() (Runner, error)
}
