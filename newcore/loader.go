package newcore

type Loader interface {
	Name() string
	LoadRunner() (Runner, error)
}
