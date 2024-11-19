package main

// var testExamplesCmd = &cobra.Command{
// 	Use:     "examples <program> [<arguments>]",
// 	Aliases: []string{"ex"},
// 	Args:    cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		examples, err := oldall.FindExamples(args[0], "examples")
// 		cobra.CheckErr(err)

// 		if len(examples) == 0 {
// 			fmt.Fprintln(os.Stderr, "No examples found")
// 		}

// 		var testResults error
// 		for _, ex := range examples {
// 			fmt.Printf("Test %s running...\n", ex.Name)
// 			t := oldall.Test{
// 				Program:        args[0],
// 				Input:          ex.Input,
// 				ExpectedOutput: ex.ExpectedOutput,
// 			}
// 			err = runTest(t, ex.Name)
// 			if err != nil {
// 				fmt.Fprintln(os.Stderr, err)
// 			}
// 			testResults = multierror.Append(testResults, err)
// 		}

// 		if err != nil {
// 			os.Exit(1)
// 		}
// 	},
// }
