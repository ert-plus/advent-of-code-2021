package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
	"strconv"
)

func read_input(filename string) []string {
	var commands []string

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		commands = append(commands, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return commands
	
}

func solve_part1(commands []string) int {
	x, y := 0, 0

	for _, s := range commands {
		cmd := strings.Split(s, " ")
		i, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatal(err)
		}

		switch cmd[0] {
		case "forward":
			x += i
		case "down":
			y += i
		case "up":
			y -= i
		}
	}

	return x*y
}

func solve_part2(commands []string) int {
	aim, x, y := 0, 0, 0

	for _, s := range commands {
		cmd := strings.Split(s, " ")
		i, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatal(err)
		}

		switch cmd[0] {
		case "forward":
			x += i
			y += aim*i
		case "down":
			aim += i
		case "up":
			aim -= i
		}
	}

	return x*y
}

func main() {
	commands := read_input("input.txt")
	
	fmt.Println("part1: ", solve_part1(commands))
	fmt.Println("part2: ", solve_part2(commands))
}

