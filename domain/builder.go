package domain

type Builder interface {
	Build(path string) (string, error)
}
