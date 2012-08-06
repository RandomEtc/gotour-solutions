package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f >= 0.0 {
		return math.Sqrt(f), nil
	}
	return 0.0, ErrNegativeSqrt(f)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}