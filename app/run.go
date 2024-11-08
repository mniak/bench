package app

import (
	"os"

	"github.com/mniak/bench/runners"
)

func Run(filename string, args ...string) error {
	r, err := runners.RunnerFor(filename)
	if err != nil {
		return err
	}
	cmd, err := r.Start(runners.Cmd{
		Path:   filename,
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		return err
	}
	return cmd.Wait()
}


