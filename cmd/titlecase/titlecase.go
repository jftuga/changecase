package main

// use 'titlecase' instead of 'title' because
// Windows has a built-in 'title' command

import (
	"fmt"
	"github.com/jftuga/changecase"
	"os"
)

const pgmName string = "titlecase"

func main() {
	if len(os.Args) == 1 {
		changecase.Usage(pgmName)
		return
	}
	fmt.Println(changecase.TitleCase(os.Args[1:]))
}
