package oldall

type Builder interface {
	Build(inputpath string) (string, error)
}
