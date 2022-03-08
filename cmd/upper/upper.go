package main

import (
	"fmt"
	"github.com/jftuga/changecase"
	"os"
)

const pgmName string = "upper"

func main() {
	if len(os.Args) == 1 {
		changecase.Usage(pgmName)
		return
	}

	fmt.Println(changecase.Upper(os.Args[1:]))
}
