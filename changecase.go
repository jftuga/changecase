package changecase

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

const PgmVersion string = "1.4.0"
const PgmUrl string = "https://github.com/jftuga/changecase"

// Usage - output when no cmd-line args are given
func Usage(pgmName string) {
	fmt.Printf("%s, v%s\n", pgmName, PgmVersion)
	fmt.Println(PgmUrl)
	fmt.Println()
	fmt.Printf("usage: %s [arguments]\n", pgmName)
	fmt.Println("(consider surrounding command-line arguments in double-quotes to preserve spacing)")
	fmt.Println()
}

// Lower - return a lower case string
func Lower(args []string) string {
	output := ""
	for _, arg := range args {
		output += strings.ToLower(arg) + " "
	}
	return output[:len(output)-1]
}

// Upper - return an upper case string
func Upper(args []string) string {
	output := ""
	for _, arg := range args {
		output += strings.ToUpper(arg) + " "
	}
	return output[:len(output)-1]
}

// TitleCase - return a title case string
func TitleCase(args []string) string {
	output := ""
	for _, arg := range os.Args[1:] {
		output += title(arg) + " "
	}
	return output[:len(output)-1]
}

// Adopted from:
// https://www.reddit.com/r/golang/comments/t9288s/how_can_i_use_stringstitle_but_also_include/hzs8s2t/
// title - similar to strings.Title, also allow for underscore
func title(s string) string {
	t := []rune(s)
	inLetters := false
	for i, r := range t {
		if unicode.IsLetter(r) {
			if !inLetters {
				t[i] = unicode.ToTitle(r)
			}
			inLetters = true
		} else {
			inLetters = false
		}
	}
	return string(t)
}
