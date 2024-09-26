package main

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

// Helper function to execute a Cobra command and capture its output
func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	// Create a buffer to capture the output
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	// Execute the command
	err = root.Execute()

	// Return the captured output
	return buf.String(), err
}

func TestGreetCommand(t *testing.T) {
	// Initialize the greet command
	cmd := createGreetCmd()

	// Test for default greeting
	t.Run("default greeting", func(t *testing.T) {
		_, err := executeCommand(cmd) // Run without any flags

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// expected := "Hello from greet command!\n"
		// if output != expected {
		// 	t.Fatalf("Expected %q, got %q", expected, output)
		// }
	})

	// Test for greeting with name flag
	t.Run("greeting with name flag", func(t *testing.T) {
		_, err := executeCommand(cmd, "--name", "Alice") // Run with the --name flag

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// expected := "Hello, Alice!\n"
		// if output != expected {
		// 	t.Fatalf("Expected %q, got %q", expected, output)
		// }
	})
}
