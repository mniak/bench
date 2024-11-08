package toolchain

import (
	"os/exec"
	"path/filepath"

	"github.com/mniak/bench/old_stuff/domain"
	"github.com/mniak/bench/old_stuff/internal/utils"
)

var goToolchainFactories = make([]domain.ToolchainFactory, 0)

func init() {
	goToolchainFactories = append(goToolchainFactories, GoToolchainFactory)
}

func NewGoLoader() domain.ToolchainLoader {
	return NewLoaderFromFactories(
		goToolchainFactories,
		[]string{".go"},
		domain.OSBinaryExtension,
	)
}

func GoToolchainFactory() (domain.Toolchain, error) {
	var result _GoToolchain
	var err error

	result.gopath, err = findBinaryPath("go")
	if err != nil {
		return &result, err
	}

	return &result, nil
}

type _GoToolchain struct {
	gopath string
}

func (tc _GoToolchain) Build(request domain.BuildRequest) error {
	workingDir, main, err := utils.SplitDirAndProgram(request.Input)
	if err != nil {
		return err
	}

	outputpath, err := filepath.Abs(request.Output)
	if err != nil {
		return err
	}

	cmd := exec.Command(tc.gopath, "build", "-o", outputpath, main)
	cmd.Stderr = request.Stderr
	cmd.Stdout = request.Stdout

	cmd.Dir = workingDir
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
