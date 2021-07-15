package domain

type Toolchain interface {
	Build(inputpath, outputpath string) (string, error)
}
type ToolchainFactory func() (Toolchain, error)

type ToolchainLoader interface {
	Load() (Toolchain, error)
	InputExtensions() []string
	OutputExtension() string
}
