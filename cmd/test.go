package cmd

import (
	"fmt"
	"strings"

	"github.com/mniak/bench/lib"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:  "test <program> [<arguments>]",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		testName, err := cmd.Flags().GetString("test-name")
		if err != nil {
			return err
		}
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}
		expectedOutput, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		workingDir, err := cmd.Flags().GetString("working-dir")
		if err != nil {
			return err
		}

		fmt.Printf("Running test %s...\n", testName)
		t := lib.Test{
			Program:        args[0],
			Input:          input,
			ExpectedOutput: expectedOutput,
			Args:           args[1:],
			WorkingDir:     workingDir,
		}
		err = t.Start()
		if err != nil {
			return err
		}

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
			return fmt.Errorf(template, "FAIL",
				strings.ReplaceAll(err.Error(), "\n", "\n"+spaces),
			)
		}
		fmt.Printf(template, "PASS", testName)

		if t.Input != "" || r.Output != "" || r.ErrorOutput != "" {
			fmt.Println("-------------- END --------------")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("working-dir", "w", "", "Working directory")
	testCmd.Flags().StringP("input", "i", "", "Test input")
	testCmd.Flags().StringP("output", "o", "", "Test output")

	testCmd.MarkFlagRequired("input")
	testCmd.MarkFlagRequired("output")
}
