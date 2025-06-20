// chomp reads from standard input and outputs all content except a trailing newline, mimicking Perl's chomp functionality.
// It preserves internal newlines while removing only the final newline character if present.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jftuga/changecase"
)

const pgmName string = "chomp"

func main() {
	versionFlag := flag.Bool("version", false, "Display version information")
	helpFlag := flag.Bool("help", false, "Display help information")
	flag.Parse()

	if *versionFlag || *helpFlag {
		fmt.Printf("%s, v%s\n", pgmName, changecase.PgmVersion)
		fmt.Println(changecase.PgmUrl)
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  chomp < input.txt")
		fmt.Println("  echo 'text' | chomp")
		fmt.Println()
		fmt.Println("Removes trailing newline from standard input while preserving internal newlines.")
		fmt.Println("Mimics Perl's chomp functionality.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var lastByte byte
	var hasBytes bool

	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}

		if hasBytes {
			if err := writer.WriteByte(lastByte); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
				os.Exit(1)
			}
		}
		lastByte = b
		hasBytes = true
	}

	if hasBytes && lastByte != '\n' {
		if err := writer.WriteByte(lastByte); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
			os.Exit(1)
		}
	}
}
