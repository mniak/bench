package cmd

import (
	"fmt"
	"os"
	"strings"

	bench "github.com/mniak/bench/lib"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:  "test [flags] -- <program> [<arguments>]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		testName, err := cmd.Flags().GetString("name")
		handle(err)

		input, err := cmd.Flags().GetString("input")
		handle(err)

		expectedOutput, err := cmd.Flags().GetString("output")
		handle(err)

		workingDir, err := cmd.Flags().GetString("dir")
		handle(err)

		if testName != "" {
			fmt.Printf("Running test " + testName + "...\n")
		} else {
			fmt.Println("Running test...")
		}

		t := bench.Test{
			Program:        args[0],
			Input:          input,
			ExpectedOutput: expectedOutput,
			Args:           args[1:],
			WorkingDir:     workingDir,
		}
		err = t.Start()
		handle(err)

		if t.Input != "" {
			fmt.Println("------------- INPUT -------------")
			fmt.Println(t.Input)
		}

		r, err := t.Wait()

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
			fmt.Println(t.ExpectedOutput)

			const spaces = "  "
			fmt.Fprintf(os.Stderr, template, "FAIL",
				strings.ReplaceAll(err.Error(), "\n", "\n"+spaces),
			)
			os.Exit(2)
		}
		fmt.Printf(template, "PASS", testName)

		if t.Input != "" || r.Output != "" || r.ErrorOutput != "" {
			fmt.Println("-------------- END --------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("input", "i", "", "Test input")
	testCmd.Flags().StringP("output", "o", "", "Expected test output")
	testCmd.Flags().StringP("dir", "d", "", "Working directory")
	testCmd.Flags().StringP("name", "n", "", "Test name")

	testCmd.MarkFlagRequired("input")
	testCmd.MarkFlagRequired("output")
}
