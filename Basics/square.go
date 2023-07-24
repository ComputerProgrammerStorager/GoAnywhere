package main

import "fmt"

func squares() func() int {
	var x int
	fmt.Printf("Squares: x value is %d and address is %x\n", x, &x)
	return func() int {
		fmt.Printf("Anonymous function: x value is %d and address is %x", x, &x)
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
