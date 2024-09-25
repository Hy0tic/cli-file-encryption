package main

import "github.com/spf13/cobra"

func createRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "CLI file encryption",
		Short: "A simple CLI tool to encrypt files",
	}
	return rootCmd
}