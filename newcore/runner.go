package newcore

import "io"

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
