package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Cannot take a negative square root of value", float64(e))
}

func Sqrt(x float64) (float64, error) {
	guess := 1.0
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	for i := 0; i < 10; i++ {
		guess -= (guess*guess - x) / (2*guess)
	}
	return guess, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
