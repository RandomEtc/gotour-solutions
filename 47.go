package main

import "fmt"
import "math/cmplx"
import "math"

func Cbrt(x complex128) complex128 {
	z := 1+0i
	for i := 0; i < 10; i++ {
		z = z - (cmplx.Pow(z,3) - x) / (3 * cmplx.Pow(z,2))
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(Cbrt(27))
	fmt.Println(math.Cbrt(8))
}