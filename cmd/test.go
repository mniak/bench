package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mniak/bench/domain"
	"github.com/mniak/bench/lib/bench"
	"github.com/spf13/cobra"
)

func testCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "test [flags] -- <program> [<arguments>]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			testName, err := cmd.Flags().GetString("name")
			cobra.CheckErr(err)

			input, err := cmd.Flags().GetString("input")
			cobra.CheckErr(err)

			expectedOutput, err := cmd.Flags().GetString("output")
			cobra.CheckErr(err)

			if testName != "" {
				fmt.Printf("Test %s running...\n", testName)
			} else {
				fmt.Println("Test running...")
			}

			t := domain.Test{
				Program:        args[0],
				Input:          input,
				ExpectedOutput: expectedOutput,
			}
			cobra.CheckErr(runTest(t, testName))
		},
	}

	cmd.Flags().StringP("input", "i", "", "Test input")
	cmd.Flags().StringP("output", "o", "", "Expected test output")
	cmd.Flags().StringP("name", "n", "", "Test name")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("output")
	return &cmd
}

func runTest(test domain.Test, testName string) error {
	started, err := bench.StartTest(test)
	cobra.CheckErr(err)

	if test.Input != "" {
		fmt.Println("------------- INPUT -------------")
		fmt.Println(test.Input)
	}

	r, err := bench.WaitTest(started)

	if r.Output != "" {
		fmt.Println("------------- OUTPUT ------------")
		fmt.Println(r.Output)
	}
	if r.ErrorOutput != "" {
		fmt.Println("------------- ERROR -------------")
		fmt.Println(r.ErrorOutput)
	}

	var template string
	if testName != "" {
		template = "Test " + testName + " - %s: %s\n"
	} else {
		template = "Test - %s: %s\n"
	}

	if err != nil {
		fmt.Println("-------- EXPECTED OUTPUT --------")
		fmt.Println(test.ExpectedOutput)

		const spaces = "  "
		fmt.Fprintf(os.Stderr, template, "FAIL",
			strings.ReplaceAll(err.Error(), "\n", "\n"+spaces),
		)
		os.Exit(2)
	}
	fmt.Printf(template, "PASS", testName)

	if test.Input != "" || r.Output != "" || r.ErrorOutput != "" {
		fmt.Println("-------------- END --------------")
	}
	return nil
}
