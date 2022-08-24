package toolchain

import (
	"os/exec"
	"path/filepath"

	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/internal/utils"
)

func init() {
	cppToolchainFactories = append(cppToolchainFactories, GPlusPlusToolchainFactory)
}

type _GPlusPlusToolchain struct {
	compiler string
}

func GPlusPlusToolchainFactory() (domain.Toolchain, error) {
	if gpppath, err := exec.LookPath("g++"); err == nil {
		return &_GPlusPlusToolchain{
			compiler: gpppath,
		}, nil
	}
	return nil, ErrToolchainNotFound
}

func (tc *_GPlusPlusToolchain) Build(request domain.BuildRequest) error {
	workingDir, main, err := utils.SplitDirAndProgram(request.Input)
	if err != nil {
		return err
	}

	outputpath, err := filepath.Abs(request.Output)
	if err != nil {
		return err
	}

	cmd := exec.Command(tc.compiler, main, "-o", outputpath)
	cmd.Stderr = request.Stderr
	cmd.Stdout = request.Stdout

	cmd.Dir = workingDir
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
