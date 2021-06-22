package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
	bench "github.com/mniak/bench"
	"github.com/spf13/cobra"
)

var testExamplesCmd = &cobra.Command{
	Use:     "examples <program> [<arguments>]",
	Aliases: []string{"ex"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		examples, err := bench.FindExamples(args[0], "examples")
		handle(err)

		if len(examples) == 0 {
			fmt.Fprintln(os.Stderr, "No examples found")
		}

		var testResults error
		for _, ex := range examples {
			fmt.Printf("Test %s running...\n", ex.Name)

			t := bench.NewTest(args[0], ex.Input, ex.ExpectedOutput)
			err = runTest(t, ex.Name)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			testResults = multierror.Append(testResults, err)
		}

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	testCmd.AddCommand(testExamplesCmd)
}
