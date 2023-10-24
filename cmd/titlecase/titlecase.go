package main

// use 'titlecase' instead of 'title' because
// Windows has a built-in 'title' command

import (
	"fmt"
	"os"

	"github.com/jftuga/changecase"
)

const pgmName string = "titlecase"

func main() {
	if len(os.Args) == 1 {
		changecase.Usage(pgmName)
		return
	}
	fmt.Printf("%v", changecase.TitleCase(os.Args[1:]))
}
