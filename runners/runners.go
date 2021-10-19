package runners

import (
	"io"
	"os/exec"
)

var loaders = []RunnerLoader{
	NewBinaryLoader(),
	NewPythonLoader(),
}

func RunnerFor(filename string) (Runner, error) {
	for _, loader := range loaders {
		can := loader.CanRun(filename)
		if !can {
			continue
		}
		r, err := loader.Load()
		if err == nil {
			return r, nil
		}
		if err == ErrRunnerNotFound {
			continue
		}
	}
	return nil, ErrRunnerNotFound
}

type RunnerLoader interface {
	Load() (Runner, error)
	CanRun(filename string) bool
}

type _StartedRunnerCmd struct {
	cmd *exec.Cmd
}

func NewStartedRunnerCmd(cmd *exec.Cmd) *_StartedRunnerCmd {
	return &_StartedRunnerCmd{
		cmd: cmd,
	}
}

type Runner interface {
	Start(cmd RunnerCmd) (StartedRunnerCmd, error)
}

func StartAndWait(r Runner, cmd RunnerCmd) error {
	startedCmd, err := r.Start(cmd)
	if err != nil {
		return err
	}
	return startedCmd.Wait()
}

type StartedRunnerCmd interface {
	Wait() error
}

func (c *_StartedRunnerCmd) Wait() error {
	return c.cmd.Wait()
}

type RunnerCmd struct {
	Path   string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
