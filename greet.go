package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func createGreetCmd() *cobra.Command {
	var name string

	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Greets the user",
		Run: func(cmd *cobra.Command, args []string) {
			if name != "" {
				fmt.Printf("Hello, %s!\n", name)
			} else {
				fmt.Println("Hello from greet command!")
			}
		},
	}

	// Add the --name flag
	greetCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the person to greet")

	return greetCmd
}