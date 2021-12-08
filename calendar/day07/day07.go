package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"math"
	"sort"
)

// take the mean
func solve_part2(positions []int) int {
	var sum int

	for _, pos := range positions {
		sum += pos
	}

	mean := float64(sum)/float64(len(positions))
	lmean, rmean := int(math.Floor(mean)), int(math.Ceil(mean))
	lfuel, rfuel := 0, 0
	for _, pos := range positions {
		lx := int(math.Abs(float64(lmean-pos)))
		rx := int(math.Abs(float64(rmean-pos)))
		lfuel += (lx*(lx+1))/2 
		rfuel += (rx*(rx+1))/2 
	}
	
	if lfuel > rfuel {
		return rfuel
	}
	return lfuel
	
}

// take the median
func solve_part1(positions []int) int {
	sort.Ints(positions)
	sorted := positions
	num_pos := len(sorted)
	median := sorted[(num_pos/2)]

	var fuel int

	for _, pos := range positions {
		fuel += int(math.Abs(float64(median-pos)))
	}
	
	return fuel
}

func read_input() []int {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	buf := make([]byte, 4096)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	
	str_pos := strings.Split(string(buf[:n]), ",")

	positions := make([]int, len(str_pos))
	
	for i, pos := range str_pos {
		n, _ := strconv.Atoi(pos)
		positions[i] = n
	}
	
	return positions
}

func main() {
	positions := read_input()

	fmt.Println("part1: ", solve_part1(positions))
	fmt.Println("part2: ", solve_part2(positions))
}
