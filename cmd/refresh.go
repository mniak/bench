package main

import (
	"log"
	"reflect"

	"github.com/mniak/bench/newcore"
	"github.com/spf13/cobra"
)

func refreshCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "refresh",
		Short: "Rebuild toolchains cache",
		Run: func(cmd *cobra.Command, args []string) {
			toolchains, err := newcore.RebuildCache()
			log.Printf("%d toolchains detected:\n", len(toolchains))
			for _, r := range toolchains {
				log.Printf(" - %s\n", reflect.TypeOf(r).Elem().Name())
			}
			cobra.CheckErr(err)
		},
	}
	return &cmd
}
