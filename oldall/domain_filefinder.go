package oldall

type FileFinder interface {
	Find(filename string) (string, error)
}
