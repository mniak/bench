package main

import (
	"fmt"

	"github.com/mniak/bench/runners"
	"github.com/spf13/cobra"
)

func refreshCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "refresh",
		Short: "Rebuild runners cache",
		Run: func(cmd *cobra.Command, args []string) {
			list, err := runners.RebuildCache()
			fmt.Printf("%d runners detected:\n", len(list))
			for _, runner := range list {
				fmt.Printf(" - %s\n", runner.Name())
			}
			cobra.CheckErr(err)
		},
	}
	return &cmd
}
