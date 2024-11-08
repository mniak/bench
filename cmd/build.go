package main

import (
	"github.com/mniak/bench/app"
	"github.com/spf13/cobra"
)

func buildCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "build <folder>",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(app.Build(args[0]))
		},
	}

	return &cmd
}
