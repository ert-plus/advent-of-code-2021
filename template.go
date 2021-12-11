package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func read_input() []string {
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

func solve_part2() int {
	return 0
}

func solve_part1() int {
	return 0
}

func main() {
	s := read_input()
	fmt.Println("part1: ", solve_part1())
	fmt.Println("part2: ", solve_part2())
}
