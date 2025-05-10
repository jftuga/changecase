/*
eq - String Comparison Utility

eq is a command-line tool that performs Unicode string comparisons and reports
whether two strings match exactly. If they don't match, it reports the position
of the first difference using 1-based indexing.

Usage:
  eq [-i] [-q] [-v] [--version] [string1 string2]

Options:
  -i  Perform case-insensitive comparison
  -q  Quiet mode (no output, only exit code)
  -v  Verbose mode (shows detailed comparison with context)
  --version  Display version information

Input:
  - If two command line arguments are provided, they will be compared
  - If no arguments are provided, two lines will be read from STDIN

Output:
  - Standard mode: Prints the position of the first difference (1-based),
    or "0" if strings match exactly
  - Verbose mode: Shows detailed information about the match/mismatch
    with context around the differing position
  - Quiet mode: No output, only exit code

Exit Codes:
  - 0 if strings match exactly
  - N (position number) if strings differ at position N

Examples:
  eq "hello" "hello"    # Will output "0" and exit with code 0
  eq "hello" "hallo"    # Will output "3" and exit with code 3
  eq -i "Hello" "hello" # Will output "0" and exit with code 0 (case-insensitive)
  eq -v "abc" "abx"     # Will show detailed difference at position 3
  echo -e "str1\nstr2" | eq  # Read strings from stdin
*/

// eq is a command-line utility for string comparison
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jftuga/changecase"
)

const pgmName string = "eq"

// processInput handles command line arguments and stdin to get the strings to compare.
// It returns the two strings to compare.
func processInput() (string, string) {
	// Check if we have non-flag arguments
	args := flag.Args()

	// If two arguments are provided, use them as the strings to compare
	if len(args) == 2 {
		return args[0], args[1]
	}

	// If no arguments are provided, read from stdin
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		// Read first line
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Error reading first line from stdin")
			os.Exit(1)
		}
		str1 := scanner.Text()

		// Read second line
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Error reading second line from stdin")
			os.Exit(1)
		}
		str2 := scanner.Text()

		return str1, str2
	}

	// Invalid number of arguments
	fmt.Fprintln(os.Stderr, "Invalid number of arguments")
	flag.Usage()
	os.Exit(1)
	return "", "" // This will never execute, but needed for compilation
}

// compareStrings compares two strings and returns:
// - 0 if strings match
// - position (1-based) of first mismatch otherwise
func compareStrings(str1, str2 string, caseInsensitive bool) int {
	// For case-insensitive comparison, convert to lowercase
	if caseInsensitive {
		str1 = strings.ToLower(str1)
		str2 = strings.ToLower(str2)
	}

	// Get the runes for proper Unicode handling
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	// Compare the strings character by character
	for i := 0; i < len(runes1) && i < len(runes2); i++ {
		if runes1[i] != runes2[i] {
			return i + 1 // Return 1-based index of mismatch
		}
	}

	// If we got here, either strings match or one is a prefix of the other
	if len(runes1) != len(runes2) {
		return min(len(runes1), len(runes2)) + 1
	}

	// Strings match
	return 0
}

// displayVerboseComparison shows a detailed comparison of the two strings
func displayVerboseComparison(str1, str2 string, position int) {
	if position == 0 {
		fmt.Println("Strings match exactly")
		return
	}

	fmt.Printf("Strings differ at position %d\n", position)

	// Display the characters at the differing position
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	// Show the difference with context
	pos := position - 1 // Convert to 0-based
	fmt.Println("Difference:")

	// Calculate safe ranges for context display
	startPos := max(0, pos-5)
	endPos1 := min(len(runes1), pos+6)
	endPos2 := min(len(runes2), pos+6)

	// Display the context with the difference highlighted
	// Handle the differing character for string 1
	var diffChar1, endChar1 string
	if pos < len(runes1) {
		diffChar1 = string(runes1[pos : pos+1])
	} else {
		diffChar1 = "END"
	}
	if pos+1 < len(runes1) {
		endChar1 = string(runes1[pos+1 : endPos1])
	} else {
		endChar1 = ""
	}

	// Handle the differing character for string 2
	var diffChar2, endChar2 string
	if pos < len(runes2) {
		diffChar2 = string(runes2[pos : pos+1])
	} else {
		diffChar2 = "END"
	}
	if pos+1 < len(runes2) {
		endChar2 = string(runes2[pos+1 : endPos2])
	} else {
		endChar2 = ""
	}

	fmt.Printf("String 1: %s[%s]%s\n",
		string(runes1[startPos:min(pos, len(runes1))]),
		diffChar1,
		endChar1)

	fmt.Printf("String 2: %s[%s]%s\n",
		string(runes2[startPos:min(pos, len(runes2))]),
		diffChar2,
		endChar2)
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Define command-line flags
	caseInsensitiveFlag := flag.Bool("i", false, "Perform case-insensitive comparison")
	quietModeFlag := flag.Bool("q", false, "Quiet mode (no output, only exit code)")
	verboseModeFlag := flag.Bool("v", false, "Verbose mode (shows detailed comparison)")
	versionFlag := flag.Bool("version", false, "Display version information")

	// Add custom usage message
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: eq [-i] [-q] [-v] [--version] [string1 string2]")
		fmt.Fprintln(os.Stderr, "  -i: Perform case-insensitive comparison")
		fmt.Fprintln(os.Stderr, "  -q: Quiet mode (no output, only exit code)")
		fmt.Fprintln(os.Stderr, "  -v: Verbose mode (shows detailed comparison)")
		fmt.Fprintln(os.Stderr, "  --version: Display version information")
		fmt.Fprintln(os.Stderr, "  If strings are not provided, reads two lines from stdin")
	}

	// Parse the flags
	flag.Parse()

	// Check if version flag is set
	if *versionFlag {
		fmt.Printf("%s, v%s\n", pgmName, changecase.PgmVersion)
		fmt.Printf("%s\n", changecase.PgmUrl)
		os.Exit(0)
	}

	// Get the strings to compare
	str1, str2 := processInput()

	// Compare the strings and get the result
	position := compareStrings(str1, str2, *caseInsensitiveFlag)

	// Handle output based on mode
	if !*quietModeFlag {
		if *verboseModeFlag {
			displayVerboseComparison(str1, str2, position)
		} else {
			// Standard output - just position
			fmt.Println(position)
		}
	}

	// Exit with the appropriate code
	if position == 0 {
		os.Exit(0)
	} else {
		os.Exit(position)
	}
}
