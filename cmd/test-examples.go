package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/go-multierror"
	bench "github.com/mniak/bench/lib"
	"github.com/spf13/cobra"
)

var testExamplesCmd = &cobra.Command{
	Use: "examples <program> [<arguments>]",
	Aliases: []string{
		"ex", "example",
		"sample", "samples",
	},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workingDir, err := cmd.Flags().GetString("dir")
		handle(err)

		examples, err := bench.FindExamples(path.Join(workingDir, "examples"))
		handle(err)

		var testResults error
		for _, ex := range examples {
			fmt.Printf("Test %s running...\n", ex.Name)

			t := bench.Test{
				Program:        args[0],
				Input:          ex.Input,
				ExpectedOutput: ex.ExpectedOutput,
				Args:           args[1:],
				WorkingDir:     workingDir,
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
	testExamplesCmd.MarkPersistentFlagRequired("dir")
}
