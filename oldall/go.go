package oldall

import (
	"os/exec"
	"path/filepath"

	"github.com/mniak/bench/old_stuff/internal/utils"
)

type GoToolchainLoader struct{}

func (l GoToolchainLoader) Load() (Toolchain, error) {
	var result _GoToolchain
	var err error

	result.GoPath, err = findBinaryPath("go")
	if err != nil {
		return &result, err
	}

	return &result, nil
}

type _GoToolchain struct {
	GoPath string
}

func (tc _GoToolchain) Name() string {
	return "go"
}

func (tc _GoToolchain) Build(request BuildRequest) error {
	workingDir, main, err := utils.SplitDirAndProgram(request.Input)
	if err != nil {
		return err
	}

	outputpath, err := filepath.Abs(request.Output)
	if err != nil {
		return err
	}

	cmd := exec.Command(tc.GoPath, "build", "-o", outputpath, main)
	cmd.Stderr = request.Stderr
	cmd.Stdout = request.Stdout

	cmd.Dir = workingDir
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
