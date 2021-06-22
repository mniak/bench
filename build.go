package bench

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/mniak/bench/toolchain"
)

type (
	ProgramFinder interface {
		Find(string) (string, error)
	}
	ToolchainProducer interface {
		Produce(string) (toolchain.Toolchain, error)
	}
	Builder interface {
		Build(string) (string, error)
	}
)

type _Builder struct {
	programFinder     ProgramFinder
	toolchainProducer ToolchainProducer
}

func (b *_Builder) Build(path string) (string, error) {
	mainfile, err := b.programFinder.Find(path)
	if err != nil {
		return "", err
	}
	tchain, err := b.toolchainProducer.Produce(mainfile)
	if err != nil {
		return "", err
	}
	return tchain.Build(mainfile)
}

type _ProgramFinder struct{}

var ErrProgramNotFound = errors.New("program not found")

func (pf *_ProgramFinder) Find(filenameOrFolder string) (string, error) {
	info, err := os.Stat(filenameOrFolder)
	if os.IsNotExist(err) {
		return "", ErrProgramNotFound
	}

	if !info.IsDir() {
		return filenameOrFolder, nil
	}

	ignoredExtensions := []string{
		".exe", ".dll", ".obj",
		".o", ".so",
	}
	files, err := ioutil.ReadDir(filenameOrFolder)
	if err != nil {
		return "", err
	}
	for _, fi := range files {
		if fi.IsDir() {
			continue
		}
		filename := fi.Name()
		for _, iext := range ignoredExtensions {
			if iext == path.Ext(filename) {
				continue
			}
		}
		return path.Join(filenameOrFolder, fi.Name()), nil
	}

	return "", ErrProgramNotFound
}

type _ToolchainProducer struct{}

func (tp *_ToolchainProducer) Produce(mainfile string) (toolchain.Toolchain, error) {
	switch path.Ext(mainfile) {
	case ".cpp", ".c++":
		return toolchain.NewCPP()
	}
	return nil, toolchain.ErrToolchainNotFound
}

var DefaultBuilder Builder = &_Builder{
	toolchainProducer: new(_ToolchainProducer),
	programFinder:     new(_ProgramFinder),
}

func Build(path string) (string, error) {
	return DefaultBuilder.Build(path)
}
