package newcore

import "reflect"

type RunnerLoader interface {
	Name() string
	LoadRunner() (Runner, error)
	RunnerType() reflect.Type
}
type CompilerLoader interface {
	Name() string
	LoadCompiler() (Compiler, error)
	CompilerType() reflect.Type
}
