package cmd

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/lib/bench"
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

			t := domain.Test{
				Program:        args[0],
				Input:          ex.Input,
				ExpectedOutput: ex.ExpectedOutput,
			}
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
