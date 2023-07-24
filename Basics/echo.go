package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Variable declaration methods
// 1. s := ""         short variable declaration, only used within function, but not package-level
// 2. var s string    use default valuation, which is ""
// 3. var s = ""      rarely used except declaring multiple variables
// 4. var s string = ""  full declaration having type, value.

func echo3() {
	fmt.Println("Using echo3 method")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// use flag package
// flag is used to parse given options. Note the parsed return value is pointer
func echo4() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func printargs() {
	fmt.Println(os.Args[0:])
}
func main() {
	echo4()
	printargs()
}
