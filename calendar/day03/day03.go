package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)


func read_input(filename string) []string {
	var lines []string
	
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return lines
}

func count_ones(lines []string) []int {
	bit_width := len(lines[0])
	one_count := make([]int, bit_width, bit_width)
	for _, line := range lines {
		for i := 0; i < bit_width; i++ {
			if line[i] == '1' {
				one_count[i] += 1
			}
		}
	}
	return one_count
}

func solve_part1(lines []string) int {
	line_count := len(lines)
	bit_width := len(lines[0])
	one_count := count_ones(lines) 
	gamma, epsilon := 0, 0
	
	// calculate gamma and epsilon by stepping through one_count
	for i := 0; i < bit_width; i++ {
		gamma, epsilon = gamma << 1, epsilon << 1
		if float64(one_count[i]) / float64(line_count) > 0.5 {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	
	return gamma*epsilon
}

// filters the list "line" for char "bit" ('0' or '1') at position "pos"
// return may be reduntant, because we are operating on a slice, idk i'm rushing
func filter_list(lines []string, pos int, bit uint8) []string {
	var output []string
	for _, line := range lines {
		if line[pos] == bit {
			output = append(output, line)
		}
	}
	return output
}

/*
func make_filter(b uint8) func([]string) []string {
	pos := 0
	bit := b
	return func(lines []string) []string {
		one_count := count_ones(lines)
		output := make([]string, len(lines))
		for i, line := range lines {

		}
	}
}
*/

func solve_part2(lines []string) int {
	oxy_list := make([]string, len(lines))
	co2_list := make([]string, len(lines))
	copy(oxy_list, lines)
	copy(co2_list, lines)
	bit_width := len(lines[0])
	for i := 0; i < bit_width; i++ {
		fmt.Println("len(oxy_list): ", len(oxy_list))
		oxy_one_count := count_ones(oxy_list)
		oxy_one_frac := float64(oxy_one_count[i]) / float64(len(oxy_list))
		if oxy_one_frac >= .5 {
			oxy_list = filter_list(oxy_list, i, '1')
		} else {
			oxy_list = filter_list(oxy_list, i, '0')
		}
		if len(oxy_list) == 1 {
			break
		}
	}

	for i := 0; i < bit_width; i++ {
		fmt.Println("len(co2_list): ", len(co2_list))
	   	co2_one_count := count_ones(co2_list)
		co2_one_frac := float64(co2_one_count[i]) / float64(len(co2_list))
		if co2_one_frac >= .5 {
			co2_list = filter_list(co2_list, i, '0')
		} else {
			co2_list = filter_list(co2_list, i, '1')
		}
		if len(co2_list) == 1 {
			break
		}
	}
	
	var co2, oxy int
	if len(co2_list) != 1 || len(oxy_list) != 1 {
		log.Fatal("list does not have only one value")
	}
	
	for i := 0; i < bit_width; i++ {
		co2, oxy = co2 << 1, oxy << 1
		if co2_list[0][i] == '1' {
			co2 += 1
		}
		if oxy_list[0][i] == '1' {
			oxy += 1
		}
	}
	return oxy * co2
}

func main() {
	lines := read_input("input.txt")
	fmt.Println("part1: ", solve_part1(lines))
	fmt.Println("part2: ", solve_part2(lines))
}
