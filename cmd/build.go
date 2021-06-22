package main

import (
	"fmt"

	bench "github.com/mniak/bench"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:  "build <folder>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		builtPath, err := bench.Build(args[0])
		handle(err)
		fmt.Println(builtPath)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
