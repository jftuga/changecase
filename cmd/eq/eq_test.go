package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

// TestStringComparison tests the basic functionality of string comparison
func TestStringComparison(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected int
	}{
		{"identical strings", "hello", "hello", 0},
		{"simple difference", "hello", "hallo", 2},
		{"different length (prefix)", "abc", "abcdef", 4},
		{"different length (shorter second)", "abcdef", "abc", 4},
		{"empty strings", "", "", 0},
		{"empty first string", "", "abc", 1},
		{"empty second string", "abc", "", 1},
		{"unicode characters", "café", "cafè", 4},
		{"multi-byte unicode", "日本語", "日本話", 3},
		{"emoji comparison", "👋😊", "👋😢", 2},
		{"whitespace difference", "hello world", "hello  world", 7},
		{"newline difference", "hello\nworld", "hello world", 6},
		{"tab difference", "hello\tworld", "hello world", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary build of the binary
			buildCmd := exec.Command("go", "build", "-o", "eq_test_binary", "eq.go")
			if err := buildCmd.Run(); err != nil {
				t.Fatalf("Failed to build test binary: %v", err)
			}
			defer os.Remove("eq_test_binary")

			// Run the binary with the test inputs
			cmd := exec.Command("./eq_test_binary", tt.str1, tt.str2)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err := cmd.Run()

			// Check exit code
			exitCode := 0
			if exiterr, ok := err.(*exec.ExitError); ok {
				exitCode = exiterr.ExitCode()
			}

			// Check standard output
			output := strings.TrimSpace(stdout.String())
			expected := strconv.Itoa(tt.expected)

			if exitCode != tt.expected {
				t.Errorf("Expected exit code %d, got %d", tt.expected, exitCode)
			}

			if output != expected {
				t.Errorf("Expected output '%s', got '%s'", expected, output)
			}
		})
	}
}

// TestCaseInsensitiveComparison tests the -i flag for case-insensitive comparison
func TestCaseInsensitiveComparison(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected int
	}{
		{"same case", "hello", "hello", 0},
		{"different case", "Hello", "hello", 0},
		{"mixed case", "HeLLo", "hEllO", 0},
		{"case diff but char diff", "Hello", "Hallo", 2},
		{"unicode case", "Café", "café", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary build of the binary
			buildCmd := exec.Command("go", "build", "-o", "eq_test_binary", "eq.go")
			if err := buildCmd.Run(); err != nil {
				t.Fatalf("Failed to build test binary: %v", err)
			}
			defer os.Remove("eq_test_binary")

			// Run the binary with the -i flag
			cmd := exec.Command("./eq_test_binary", "-i", tt.str1, tt.str2)
			var stdout bytes.Buffer
			cmd.Stdout = &stdout
			err := cmd.Run()

			// Check exit code
			exitCode := 0
			if exiterr, ok := err.(*exec.ExitError); ok {
				exitCode = exiterr.ExitCode()
			}

			if exitCode != tt.expected {
				t.Errorf("Case-insensitive: Expected exit code %d, got %d", tt.expected, exitCode)
			}
		})
	}
}

// TestQuietMode tests the -q flag for quiet output
func TestQuietMode(t *testing.T) {
	// Create a temporary build of the binary
	buildCmd := exec.Command("go", "build", "-o", "eq_test_binary", "eq.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}
	defer os.Remove("eq_test_binary")

	// Test with matching strings
	t.Run("matching strings quiet mode", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-q", "hello", "hello")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		err := cmd.Run()

		exitCode := 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			exitCode = exiterr.ExitCode()
		}

		if exitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", exitCode)
		}

		if output := strings.TrimSpace(stdout.String()); output != "" {
			t.Errorf("Expected no output in quiet mode, got '%s'", output)
		}
	})

	// Test with non-matching strings
	t.Run("non-matching strings quiet mode", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-q", "hello", "hallo")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		err := cmd.Run()

		exitCode := 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			exitCode = exiterr.ExitCode()
		}

		if exitCode != 2 {
			t.Errorf("Expected exit code 2, got %d", exitCode)
		}

		if output := strings.TrimSpace(stdout.String()); output != "" {
			t.Errorf("Expected no output in quiet mode, got '%s'", output)
		}
	})
}

// TestVerboseMode tests the -v flag for verbose output
func TestVerboseMode(t *testing.T) {
	// Create a temporary build of the binary
	buildCmd := exec.Command("go", "build", "-o", "eq_test_binary", "eq.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}
	defer os.Remove("eq_test_binary")

	// Test with matching strings
	t.Run("matching strings verbose mode", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-v", "hello", "hello")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		err := cmd.Run()

		exitCode := 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			exitCode = exiterr.ExitCode()
		}

		if exitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", exitCode)
		}

		output := strings.TrimSpace(stdout.String())
		if !strings.Contains(output, "match") {
			t.Errorf("Expected verbose output to mention 'match', got '%s'", output)
		}
	})

	// Test with non-matching strings
	t.Run("non-matching strings verbose mode", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-v", "hello", "hallo")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		err := cmd.Run()

		exitCode := 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			exitCode = exiterr.ExitCode()
		}

		if exitCode != 2 {
			t.Errorf("Expected exit code 2, got %d", exitCode)
		}

		output := strings.TrimSpace(stdout.String())
		if !strings.Contains(output, "differ") {
			t.Errorf("Expected verbose output to mention 'differ', got '%s'", output)
		}
		if !strings.Contains(output, "[l]") && !strings.Contains(output, "[a]") {
			t.Errorf("Expected verbose output to highlight difference, got '%s'", output)
		}
	})
}

// TestStdinInput tests reading input from stdin
func TestStdinInput(t *testing.T) {
	// Create a temporary build of the binary
	buildCmd := exec.Command("go", "build", "-o", "eq_test_binary", "eq.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}
	defer os.Remove("eq_test_binary")

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"matching strings", "hello\nhello", 0},
		{"different strings", "hello\nhallo", 2},
		// Removed the empty lines test from here and put it in its own test function
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("./eq_test_binary")

			// Provide input via stdin
			cmd.Stdin = strings.NewReader(tt.input)

			var stdout bytes.Buffer
			cmd.Stdout = &stdout
			err := cmd.Run()

			exitCode := 0
			if exiterr, ok := err.(*exec.ExitError); ok {
				exitCode = exiterr.ExitCode()
			}

			if exitCode != tt.expected {
				t.Errorf("Expected exit code %d, got %d", tt.expected, exitCode)
			}
		})
	}
}

// TestCombinedFlags tests using multiple flags together
func TestCombinedFlags(t *testing.T) {
	// Create a temporary build of the binary
	buildCmd := exec.Command("go", "build", "-o", "eq_test_binary", "eq.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}
	defer os.Remove("eq_test_binary")

	// Test case-insensitive and verbose together
	t.Run("case-insensitive and verbose", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-i", "-v", "Hello", "hello")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		err := cmd.Run()

		exitCode := 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			exitCode = exiterr.ExitCode()
		}

		if exitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", exitCode)
		}

		output := strings.TrimSpace(stdout.String())
		if !strings.Contains(output, "match") {
			t.Errorf("Expected output to contain 'match', got '%s'", output)
		}
	})

	// Test case-insensitive and quiet together
	t.Run("case-insensitive and quiet", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-i", "-q", "Hello", "hello")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		err := cmd.Run()

		exitCode := 0
		if exiterr, ok := err.(*exec.ExitError); ok {
			exitCode = exiterr.ExitCode()
		}

		if exitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", exitCode)
		}

		output := strings.TrimSpace(stdout.String())
		if output != "" {
			t.Errorf("Expected no output in quiet mode, got '%s'", output)
		}
	})

	// Test all flags together
	t.Run("all flags together", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-i", "-q", "-v", "Hello", "hello")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		if err := cmd.Run(); err != nil {
			// Only check if there's an error, we're expecting success
			t.Errorf("Unexpected error running with empty lines: %v", err)
		}

		// Even with -v, quiet mode should take precedence
		output := strings.TrimSpace(stdout.String())
		if output != "" {
			t.Errorf("Expected no output with quiet mode, even with verbose flag, got '%s'", output)
		}
	})
}

// TestErrorCases tests handling of error cases
func TestErrorCases(t *testing.T) {
	// Create a temporary build of the binary
	buildCmd := exec.Command("go", "build", "-o", "eq_test_binary", "eq.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}
	defer os.Remove("eq_test_binary")

	// Test with insufficient arguments
	t.Run("insufficient arguments", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "only_one_arg")
		if err := cmd.Run(); err == nil {
			t.Errorf("Expected an error with insufficient arguments, got none")
		}
	})

	// Test with unknown flag
	t.Run("unknown flag", func(t *testing.T) {
		cmd := exec.Command("./eq_test_binary", "-z", "hello", "hello")
		if err := cmd.Run(); err == nil {
			t.Errorf("Expected an error with unknown flag, got none")
		}
	})
}
