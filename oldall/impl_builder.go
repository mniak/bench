package oldall

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type _BaseBuilder struct {
	toolchainFinder ToolchainFinder
}

func (b *_BaseBuilder) Build(inputpath string) (string, error) {
	tloader, err := b.toolchainFinder.Find(inputpath)
	if err != nil {
		return "", err
	}
	tchain, err := tloader.Load()
	if err != nil {
		return "", err
	}
	outputpath := getBinaryPath(inputpath)
	var buffout bytes.Buffer
	var bufferr bytes.Buffer

	req := BuildRequest{
		Stdout: &buffout,
		Stderr: &bufferr,
		Input:  inputpath,
		Output: outputpath,
	}
	err = tchain.Build(req)
	if err != nil {
		io.Copy(os.Stdout, &buffout)
		io.Copy(os.Stderr, &bufferr)
	}
	return outputpath, err
}

func NewBuilder(toolchainFinder ToolchainFinder) Builder {
	return &_BaseBuilder{
		toolchainFinder: toolchainFinder,
	}
}

func getBinaryPath(inputpath string) string {
	return strings.TrimSuffix(inputpath, filepath.Ext(inputpath)) + OSBinaryExtension
}
