package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func readInput() []string {
	var output []string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func solvePart2() int {
	return 0
}

func solvePart1() int {
	return 0
}

func main() {
	s := readInput()
	fmt.Println("part1: ", solvePart1())
	fmt.Println("part2: ", solvePart2())
}
