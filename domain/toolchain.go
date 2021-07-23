package domain

type ToolchainFinder interface {
	Find(string) (ToolchainLoader, error)
}

type ToolchainLoader interface {
	Load() (Toolchain, error)
	InputExtensions() []string
	OutputExtension() string
}

type Toolchain interface {
	Build(inputpath, outputpath string) error
}

type ToolchainFactory func() (Toolchain, error)
