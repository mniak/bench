package domain

type ToolchainFinder interface {
	Find(string) (Toolchain, error)
}
