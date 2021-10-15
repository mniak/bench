package impl

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mniak/bench/domain"
)

type _BaseBuilder struct {
	toolchainFinder domain.ToolchainFinder
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

	req := domain.BuildRequest{
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

func NewBuilder(toolchainFinder domain.ToolchainFinder) domain.Builder {
	return &_BaseBuilder{
		toolchainFinder: toolchainFinder,
	}
}

func getBinaryPath(inputpath string) string {
	return strings.TrimSuffix(inputpath, filepath.Ext(inputpath)) + domain.OSBinaryExtension
}
