package domain

type Toolchain interface {
	Build(mainfile string) (string, error)
	OutputExtension() string
	InputExtensions() []string
}
