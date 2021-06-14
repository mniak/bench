package cmd

import (
	bench "github.com/mniak/bench/lib"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:  "build <folder>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handle(bench.Build(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
