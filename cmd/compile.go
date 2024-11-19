package main

import (
	"github.com/mniak/bench/newcore"
	"github.com/spf13/cobra"
)

func compileCmd() *cobra.Command {
	var flagOut string
	cmd := cobra.Command{
		Use:  "compile <filename>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := Compile(args[0], flagOut)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVarP(&flagOut, "output", "o", "", "Output file name")
	cmd.MarkFlagRequired("output")
	return &cmd
}

func Compile(filename string, output string) error {
	c, err := newcore.CompilerFor(filename)
	if err != nil {
		return err
	}

	err = c.Compile(newcore.CompilationInput{
		Filename:       filename,
		OutputFilename: output,
	})
	if err != nil {
		return err
	}
	return nil
}
