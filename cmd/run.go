package main

import (
	"github.com/mniak/bench/app"
	"github.com/mniak/bench/runners"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "run <filename>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := app.Run(args[0], args[1:]...)
			cobra.CheckErr(err)
		},
	}
	return &cmd
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
