package main

import (
	"fmt"
	"os"
	"strings"

	bench "github.com/mniak/bench"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:  "test [flags] -- <program> [<arguments>]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		testName, err := cmd.Flags().GetString("name")
		handle(err)

		input, err := cmd.Flags().GetString("input")
		handle(err)

		expectedOutput, err := cmd.Flags().GetString("output")
		handle(err)

		if testName != "" {
			fmt.Printf("Test %s running...\n", testName)
		} else {
			fmt.Println("Test running...")
		}

		t := bench.Test{
			Program:        args[0],
			Input:          input,
			ExpectedOutput: expectedOutput,
		}
		handle(runTest(t, testName))
	},
}

func runTest(test bench.Test, testName string) error {
	err := test.Start()
	handle(err)

	if test.Input != "" {
		fmt.Println("------------- INPUT -------------")
		fmt.Println(test.Input)
	}

	r, err := test.Wait()

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

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("input", "i", "", "Test input")
	testCmd.Flags().StringP("output", "o", "", "Expected test output")
	testCmd.Flags().StringP("name", "n", "", "Test name")

	testCmd.MarkFlagRequired("input")
	testCmd.MarkFlagRequired("output")
}
