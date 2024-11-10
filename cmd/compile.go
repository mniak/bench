package main

import (
	"fmt"

	"github.com/mniak/bench/newcore"
	"github.com/spf13/cobra"
)

func compileCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "compile <filename>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := Compile(args[0])
			cobra.CheckErr(err)
		},
	}
	return &cmd
}

func Compile(filename string) error {
	c, err := newcore.CompilerFor(filename)
	if err != nil {
		return err
	}
	artifact, err := c.Compile(newcore.CompilerInput{
		Filename: filename,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Output file: %s\n", artifact.OutputFilename)
	return nil
}
