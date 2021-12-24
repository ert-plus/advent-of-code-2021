package main

import (
	"fmt"
	//	"golang.org/x/tour/pic"
	//	"golang.org/x/tour/wc"
	"strings"
)

func Sqrt(x float64) float64 {
	guess := 1.0
	for i := 0; i < 10; i++ {
		guess -= (guess*guess - x) / (2*guess)
	}
	return guess
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 4}
	v3 = Vertex{}
	p = &Vertex{2,3}
)
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

var (
	length = 200
	width = 400
)

func Pic(dx, dy int) [][]uint8 {
	var out [][]uint8
	
	for y := 0; y < dy; y++ {
		out = append(out, make([]uint8, dx))
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			out[x][y] = uint8(x^y)
		}
	}
	
	return out
}

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	fields := strings.Fields(s)
	for _, f := range fields {
		wc[f] += 1
	}
	return wc
}

// adder means a function that adds
// it is a function that takes no input and
// returns a function (a closure). the closer function
// it returns is bound to a variable in the local scope of adder()
// each function is bound to it's own variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
/*
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // calling the closures 
			neg(-2*i), // they are functions, after all
		)
	}
}
*/

func fibonacci() func() int {
	prev := 0
	curr := 1
	return func() int {
		next := curr + prev
		prev, curr = curr, next
		return prev
	}
}
/*
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
*/
