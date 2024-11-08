package domain

type Builder interface {
	Build(inputpath string) (string, error)
}
