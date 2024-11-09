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

func Run(filename string, args ...string) error {
	r, err := newcore.RunnerFor(filename)
	if err != nil {
		return err
	}
	cmd, err := r.Start(newcore.Cmd{
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
