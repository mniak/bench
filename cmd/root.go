package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use: "bench",
	}
	rootCmd.AddCommand(buildCmd())
	rootCmd.AddCommand(runCmd(), rebuildRunnersCacheCmd())
	rootCmd.AddCommand(testCmd())

	rootCmd.Execute()
}
