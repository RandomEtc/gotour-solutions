package main

import (
	"fmt"
	"math"
)

const Epsilon = 0.00000001

func Sqrt(x float64) float64 {
	z := 1.0
	prevz := 0.0
	for z - prevz > Epsilon || z - prevz < -Epsilon {
		prevz = z
		z = z - (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println("Sqrt(%d) == %g", 1, Sqrt(1))
	fmt.Println("Sqrt(%d) == %g", 2, Sqrt(2))
	fmt.Println("Sqrt(%d) == %g", 3, Sqrt(3))
	fmt.Println("Sqrt(%d) == %g", 4, Sqrt(4))
	fmt.Println("math.Sqrt(%d) - Sqrt(%d) == %g", 1, 1, math.Sqrt(1) - Sqrt(1))
	fmt.Println("math.Sqrt(%d) - Sqrt(%d) == %g", 2, 2, math.Sqrt(2) - Sqrt(2))
	fmt.Println("math.Sqrt(%d) - Sqrt(%d) == %g", 3, 3, math.Sqrt(3) - Sqrt(3))
	fmt.Println("math.Sqrt(%d) - Sqrt(%d) == %g", 4, 4, math.Sqrt(4) - Sqrt(4))
}
