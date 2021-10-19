package runner

import (
	"os/exec"

	"github.com/mniak/bench/domain"
)

type _GenericRunner struct {
	program string
}

func (r *_GenericRunner) Start(runnerCmd domain.RunnerCmd) (domain.StartedRunnerCmd, error) {
	cmd := exec.Command(r.program, runnerCmd.Path)
	cmd.Stdin = runnerCmd.Stdin
	cmd.Stdout = runnerCmd.Stdout
	cmd.Stderr = runnerCmd.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return domain.NewStartedRunnerCmd(cmd), nil
}

func newGenericRunner(program string) *_GenericRunner {
	return &_GenericRunner{
		program: program,
	}
}
