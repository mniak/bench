package newcore

import (
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

type Runner interface {
	Toolchain
	CanRun(filename string) bool
	Start(program string, a RunArgs) (Waiter, error)
	RunnerInputExtensions() []string
}

func StartAndWait(r Runner, program string, a RunArgs) error {
	startedCmd, err := r.Start(program, a)
	if err != nil {
		return err
	}
	return startedCmd.Wait()
}

type Waiter interface {
	Wait() error
}

func (c *startedProgramImpl) Wait() error {
	return c.cmd.Wait()
}

type RunArgs struct {
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type startedProgramImpl struct {
	cmd *exec.Cmd
}

func newStartedRunnerCmd(cmd *exec.Cmd) *startedProgramImpl {
	return &startedProgramImpl{
		cmd: cmd,
	}
}

// RunnerFor tries to find a suitable runner for a specific file
func RunnerFor(filename string) (Runner, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	for runner := range iterateToolchains[Runner] {
		can := runner.CanRun(filename)
		if can {
			return runner, nil
		}
	}

	binaryRunner := BinaryRunner()
	if binaryRunner.CanRun(filename) {
		return binaryRunner, nil
	}
	return nil, errors.New("no suitable runner found for file")
}
