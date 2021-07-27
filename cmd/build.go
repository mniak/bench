package cmd

import (
	"fmt"

	"github.com/mniak/bench/lib/bench"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:  "build <folder>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handle(Build(args[0]))
	},
}

func Build(folder string) error {
	builtPath, err := bench.Build(folder)
	if err == nil {
		fmt.Println(builtPath)
	}
	return err
}

func init() {
	rootCmd.AddCommand(runCmd)
}
