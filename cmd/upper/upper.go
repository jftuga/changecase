package main

import (
	"fmt"
	"os"

	"github.com/jftuga/changecase"
)

const pgmName string = "upper"

func main() {
	if len(os.Args) == 1 {
		changecase.Usage(pgmName)
		return
	}

	fmt.Printf("%v", changecase.Upper(os.Args[1:]))
}
