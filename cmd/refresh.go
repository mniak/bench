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
			runners, compilers, err := newcore.RebuildCache()
			fmt.Printf("%d runners detected:\n", len(runners))
			for _, r := range runners {
				fmt.Printf(" - %s\n", r.Name())
			}
			fmt.Printf("%d compilers detected:\n", len(compilers))
			for _, c := range compilers {
				fmt.Printf(" - %s\n", c.Name())
			}
			cobra.CheckErr(err)
		},
	}
	return &cmd
}
