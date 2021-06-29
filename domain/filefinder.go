package domain

type FileFinder interface {
	Find(filename string) (string, error)
}
