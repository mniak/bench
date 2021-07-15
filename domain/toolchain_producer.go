package domain

type ToolchainFinder interface {
	Find(string) (ToolchainLoader, error)
}
