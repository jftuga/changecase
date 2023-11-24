package main

/*
len.go
-John Taylor
March 2019

Return the combined string length of all of given command line arguments.

To compile:
go build -ldflags="-s -w" len.go

MIT License; Copyright (c) 2019 John Taylor
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*/

import (
	"fmt"
	"os"
	"strings"
)

const version = "1.0.0"

func main() {
	if 1 == len(os.Args) {
		fmt.Printf("\nUsage: %s \"string\"\n\n", os.Args[0])
		fmt.Printf("This program assumes that there is only one space between each command line argument.\n")
		fmt.Printf("The most accurate way to get a string length is to surround all of your command line arguments between double-quotes.\n\n")
		os.Exit(1)
	}
	fmt.Println(len(strings.Join(os.Args[1:], " ")))
}
