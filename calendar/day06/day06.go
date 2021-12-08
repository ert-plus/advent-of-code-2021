package main

import (
	"fmt"
	//	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_input() [9]int {
	var fish [9]int

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	
	ages := strings.Split(string(buf[:n]), ",")
	
	for _, age := range ages {
		n, _ := strconv.Atoi(age)
		fish[n] += 1
	}
	
	return fish
}

func solve(fish [9]int, days int) int {
	for d := 0; d < days; d ++ {
		new := fish[0]
		for i := 1; i < 9; i ++ {
			fish[i-1] = fish[i]
		}
		fish[6] += new
		fish[8] = new
	}
	
	var sum int
	for i := 0; i < 9; i ++ {
		sum += fish[i]
	}
	return sum
}

func print_fish(fish [9]int) {
	for i, f := range fish {
		fmt.Println("age", i, ":", f)
	}
}

func main() {
	fish := read_input()
	// fmt.Println("starting ages:"
	// print_fish(fish)
	fmt.Println("part1: ", solve(fish, 80))
	fmt.Println("part2: ", solve(fish, 256))
}

