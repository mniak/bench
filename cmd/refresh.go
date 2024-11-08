package main

import (
	"github.com/mniak/bench/runners"
	"github.com/spf13/cobra"
)

func refreshCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "refresh",
		Short: "Rebuild runners cache",
		Run: func(cmd *cobra.Command, args []string) {
			err := runners.RebuildCache()
			cobra.CheckErr(err)
		},
	}
	return &cmd
}
