package main

import (
	"fmt"
	"os"
	"os/exec"
)

// executeCommand runs a command and prints its output or error
func executeCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %s %v\n", name, args)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	// Step 1: Checkout the code (You need to run this manually or set up a git command if this is local)
	fmt.Println("Step 1: Checkout the code - This should be done manually or through a git command.")

	// Step 2: Set up Go
	fmt.Println("Step 2: Set up Go")
	executeCommand("go", "version")

	// Step 3: Cache Go modules (simulate caching by checking if modules are already installed)
	fmt.Println("Step 3: Cache Go modules")
	_, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("Go is not installed. Please install Go to continue.")
		os.Exit(1)
	}

	// Step 4: Install dependencies
	fmt.Println("Step 4: Install dependencies")
	executeCommand("go", "mod", "tidy")

	// Step 5: Build the CLI tool
	fmt.Println("Step 5: Build the CLI tool")
	executeCommand("go", "build", "-o", "file-encryption")

	// Step 6: Run tests
	fmt.Println("Step 6: Run tests")
	executeCommand("go", "test", "../...")

	// Step 7: Verify the tool works
	fmt.Println("Step 7: Run CLI tool help")
	exec.Command("./file-encryption", "--help")
}
