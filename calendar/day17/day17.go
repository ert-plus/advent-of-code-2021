package main

import (
	"fmt"
)

type Bound struct {
	xMin, xMax, yMin, yMax int
}

type Vector struct {
	X, Y int
}

func read_input() Bound {
	var output Bound
	// yahh... not gonna write a parser for 4 integers
	/*
	output.xMin = 20
	output.xMax = 30
	output.yMin = -10
	output.yMax = -5
	*/
	
	output.xMin = 119
	output.xMax = 176
	output.yMin = -141
	output.yMax = -84
	
	return output
}

func solve_part2(b Bound) int {
	count := 0
	for x := 0; x <= b.xMax; x ++ {
		for y := b.yMin; y <= -b.yMin; y ++ {
			vec := Vector{x,y}
			traject := shootProbe(vec, b)
			last := traject[len(traject)-1]
			if last.X <= b.xMax && last.X >= b.xMin &&
				last.Y <= b.yMax && last.Y >= b.yMin {
				// fmt.Println(vec)
				count ++ 
			}
		}
	}
	return count
}

// returns a slice of Vector positions that is the trajectory of the probe
// when you shoot it with initial velocity v. Stops when positions are no longer
// are larger than b.yMin and smaller than xMax
func shootProbe(v Vector, b Bound) []Vector {
	out := make([]Vector, 0)
	pos := Vector{0,0}
	out = append(out, pos)
	for step := 0; ; step ++ {
		pos.X += v.X
		pos.Y += v.Y
		if pos.X > b.xMax || pos.Y < b.yMin {
			break
		}
		out = append(out, pos)
		v.Y -= 1
		switch {
		case v.X > 0:
			v.X -= 1
		case v.X < 0:
			v.X += 1
		}	
	}
	return out
}

func solve_part1(b Bound) int {
	// so this works assuming yMin is negative
	// because the trajectorys y values will be the same going up and down
	// and the max height is achieved when y=0 on the way down
	// it reaches yMin in the next step. then initial vector's y is -(yMin+1)
	// and this is the height, the -(yMin+1)th trianglar number or whatever 
	return (-b.yMin-1)*(-b.yMin)/2
}

func main() {
	bound := read_input()
	fmt.Println("part1: ", solve_part1(bound))
	fmt.Println("part2: ", solve_part2(bound))
}
