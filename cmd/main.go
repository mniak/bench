package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "bench",
	}
	cmd.AddCommand(refreshCmd())
	cmd.AddCommand(runCmd())
	cmd.AddCommand(compileCmd())
	cmd.AddCommand(testCmd())

	cmd.Execute()
}
