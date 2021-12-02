package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

func read_input(f string) []int {
 	var depths []int
	
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// reading this scanners code makes my head explode
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, i)
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return depths
}

func solve_part1(depths []int) int {
	var	count, prev int = 0, depths[0]
	for _, d := range depths {
		if d > prev {
			count++
		}
		prev = d
	}
	return count
}

func solve_part2(depths []int) int {
	count := 0
	for i := 0; i < len(depths) - 3; i ++ {
		prev := depths[i] + depths[i+1] + depths[i+2]
		next := depths[i+1] + depths[i+2] + depths[i+3]
		if next > prev {
			count++
		}
	}
	return count
}

func main() {
	depths := read_input("input")
	fmt.Println("part1: ", solve_part1(depths))
	fmt.Println("part2: ", solve_part2(depths))
}
