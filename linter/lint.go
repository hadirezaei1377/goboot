package linter

import "fmt"

func PrintSomething() {
	fmt.Println("Hello, World")
}

func main() {
	PrintSomething()
}

// codes problem:
/*
PrintSomething must be camleCase in golang
.
.
.
*/

// start linter
/*
golint

main.go:6:1: comment on exported function PrintSomething should be of the form "PrintSomething ..."
main.go:6:6: func name will be used as main.PrintSomething by other packages, and that stutters; consider calling this Something

*/

// modified code:
/*

package main

import "fmt"

func printSomething() {
    fmt.Println("Hello, World")
}

func main() {
    printSomething()
}

*/
