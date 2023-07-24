package main

import "fmt"

var p = f()

func f() *int {
	v := 1
	return &v
}

func main() {
	var p1 = f()
	var p2 = f()
	fmt.Printf("value of p1 is %d p2 is %d\n", p1, p2)
}
