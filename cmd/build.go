package cmd

import (
	"fmt"

	"github.com/mniak/bench/lib/bench"
	"github.com/spf13/cobra"
)

func buildCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "build <folder>",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(Build(args[0]))
		},
	}

	return &cmd
}

func Build(folder string) error {
	builtPath, err := bench.Build(folder)
	if err == nil {
		fmt.Println(builtPath)
	}
	return err
}
