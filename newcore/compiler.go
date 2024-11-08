package newcore

import "os"

type Compiler interface {
	Name() string
	Build() (*Artifact, error)
}

type Artifact struct {
	filename string
}

func (a *Artifact) Filename() string {
	return a.filename
}

func (a *Artifact) Free() error {
	return os.RemoveAll(a.filename)
}
