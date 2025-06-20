package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestChomp(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Test cases
		{"Hello, World!\n", "Hello, World!"},
		{"Hello, World!", "Hello, World!"},
		{"\n", ""},
		{"", ""},
		{"Line 1\nLine 2\n", "Line 1\nLine 2"},
		{"Line 1\nLine 2", "Line 1\nLine 2"},
	}

	for _, test := range tests {
		// Run the chomp program with the test input
		output, err := runChomp(test.input)
		if err != nil {
			t.Fatalf("Error running chomp: %v", err)
		}

		// Compare the output with the expected result
		if output != test.expected {
			t.Errorf("Input: %q\nExpected: %q\nGot: %q", test.input, test.expected, output)
		}
	}
}

// runChomp simulates running the chomp program with the given input
func runChomp(input string) (string, error) {
	// Create a command to run the chomp program
	cmd := exec.Command("go", "run", "chomp.go")

	// Create pipes for stdin and stdout
	var stdout bytes.Buffer
	cmd.Stdin = bytes.NewReader([]byte(input))
	cmd.Stdout = &stdout

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Return the output as a string
	return stdout.String(), nil
}
