package main

import (
	"fmt"
	"github.com/jftuga/changecase"
	"os"
)

const pgmName string = "lower"

func main() {
	if len(os.Args) == 1 {
		changecase.Usage(pgmName)
		return
	}

	fmt.Println(changecase.Lower(os.Args[1:]))
}
