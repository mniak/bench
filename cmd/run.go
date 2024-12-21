package main

import (
	"os"

	"github.com/mniak/bench/newcore"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "run <filename>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := Run(args[0], args[1:]...)
			cobra.CheckErr(err)
		},
	}
	return &cmd
}

func Run(program string, args ...string) error {
	r, err := newcore.RunnerFor(program)
	if err != nil {
		return err
	}
	a := newcore.RunArgs{
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmd, err := r.Start(program, a)
	if err != nil {
		return err
	}
	return cmd.Wait()
}
