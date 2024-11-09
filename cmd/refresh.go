package main

import (
	"fmt"

	"github.com/mniak/bench/newcore"
	"github.com/spf13/cobra"
)

func refreshCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "refresh",
		Short: "Rebuild runners cache",
		Run: func(cmd *cobra.Command, args []string) {
			list, err := newcore.RebuildCache()
			fmt.Printf("%d run detected:\n", len(list))
			for _, runner := range list {
				fmt.Printf(" - %s\n", runner.Name())
			}
			cobra.CheckErr(err)
		},
	}
	return &cmd
}
