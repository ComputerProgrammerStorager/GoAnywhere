package main

import (
	"fmt"
	"os"
	"runtime"
)

func defer1(x int) {
	fmt.Printf("defer1(%x)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	defer1(x - 1)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	defer printStack()
	defer1(3)
}
