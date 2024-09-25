package main

func main() {
	rootCmd := createRootCmd()
	rootCmd.AddCommand(createGreetCmd())
	rootCmd.AddCommand(createEncryptCmd())
	rootCmd.AddCommand(createDecryptCmd())

	rootCmd.Execute()
}
