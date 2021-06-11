package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "bench",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
