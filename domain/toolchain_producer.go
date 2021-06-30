package domain

type ToolchainFinder interface {
	Produce(string) (Toolchain, error)
}
