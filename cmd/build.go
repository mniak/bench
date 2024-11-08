package main

import (
	"fmt"

	"github.com/mniak/bench/old_stuff/lib/bench"
	"github.com/spf13/cobra"
)

func buildCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "build <folder>",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			builtPath, err := bench.Build(args[0])
			cobra.CheckErr(err)
			if err == nil {
				fmt.Println(builtPath)
			}
		},
	}

	return &cmd
}
