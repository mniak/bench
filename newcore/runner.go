package newcore

import (
	"io"
	"os/exec"
)

type Runner interface {
	Name() string
	CanRun(filename string) bool
	Start(cmd Cmd) (StartedCmd, error)
}

func StartAndWait(r Runner, cmd Cmd) error {
	startedCmd, err := r.Start(cmd)
	if err != nil {
		return err
	}
	return startedCmd.Wait()
}

type StartedCmd interface {
	Wait() error
}

func (c *_StartedRunnerCmd) Wait() error {
	return c.cmd.Wait()
}

type Cmd struct {
	Path   string
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type _StartedRunnerCmd struct {
	cmd *exec.Cmd
}

func newStartedRunnerCmd(cmd *exec.Cmd) *_StartedRunnerCmd {
	return &_StartedRunnerCmd{
		cmd: cmd,
	}
}
