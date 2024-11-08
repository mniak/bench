package newcore

type RunnerLoader interface {
	Name() string
	LoadRunner() (Runner, error)
}
type CompilerLoader interface {
	Name() string
	LoadCompiler() (Runner, error)
}
