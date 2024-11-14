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
			toolchains, err := newcore.RebuildCache()
			fmt.Printf("%d toolchains detected:\n", len(toolchains))
			for _, r := range toolchains {
				fmt.Printf(" - %s\n", r.Name())
			}
			cobra.CheckErr(err)
		},
	}
	return &cmd
}
