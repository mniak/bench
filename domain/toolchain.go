package domain

type Toolchain interface {
	Build(mainfile string) (string, error)
}
type ToolchainFactory func() (Toolchain, error)

type ToolchainLoader interface {
	Load() (Toolchain, error)
	OutputExtension() string
	InputExtensions() []string
}
