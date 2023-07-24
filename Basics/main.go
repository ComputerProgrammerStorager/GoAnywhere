package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello 世界!")
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
	fmt.Printf("MaxFloat32 is %g\n", math.MaxFloat32)
	var f float32 = 16777216
	fmt.Println(f == f+1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f", x, math.Exp(float64(x)))
	}
	fmt.Println()
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	nan := math.NaN()
	fmt.Println(math.IsNaN(nan))

	var x_ complex128 = complex(1, 2)
	var y_ complex128 = complex(3, 4)
	fmt.Println(x_ * y_)
	fmt.Println(real(x_ * y_))
	fmt.Println(imag(x_ * y_))
	fmt.Println("\a")
}
