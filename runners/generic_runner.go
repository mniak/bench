package runners

import (
	"os/exec"
)

type _GenericRunner struct {
	program string
}

func (r *_GenericRunner) Start(runnerCmd RunnerCmd) (StartedRunnerCmd, error) {
	cmd := exec.Command(r.program, runnerCmd.Path)
	cmd.Stdin = runnerCmd.Stdin
	cmd.Stdout = runnerCmd.Stdout
	cmd.Stderr = runnerCmd.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return NewStartedRunnerCmd(cmd), nil
}

func newGenericRunner(program string) *_GenericRunner {
	return &_GenericRunner{
		program: program,
	}
}
