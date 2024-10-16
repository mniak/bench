package cmd

import (
	"os"

	"github.com/mniak/bench/runners"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "run <filename>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(Run(args[0], args[1:]...))
		},
	}
	return &cmd
}

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

func rebuildRunnersCacheCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "rebuild-runners-cache <filename>",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(runners.RebuildCache())
		},
	}
	return &cmd
}
