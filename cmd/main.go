package main

import (
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "bench",
	}
	rootCmd.AddCommand(buildCmd())
	rootCmd.AddCommand(runCmd(), rebuildRunnersCacheCmd())
	rootCmd.AddCommand(testCmd())

	rootCmd.Execute()
}
