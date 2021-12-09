package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
	"strings"
	"sort"
)

func solve_part1(signal_patterns, output_values [][]string) int {
	var easy_digit_freq int
	for _, four_digits := range output_values {
		for _, digit := range four_digits {
			switch len(digit) {
			case 2:
				easy_digit_freq++
			case 3:
				easy_digit_freq++
			case 4:
				easy_digit_freq++
			case 7:
				easy_digit_freq++
			default:
				continue
			}
		}
	}
	return easy_digit_freq
}

// finds segments common between two patterns
func pattern_intersect(p1 string, p2 string) string {
	out := ""
	for _, char := range p1 {
		if c := string(char); strings.Index(p2, c) != -1 {
			out += c
		}
	}
	return out
}

func sort_pattern(pattern string) string {
	chars := strings.Split(pattern, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// code is a map and something like {"ab" => 1, "fbea" => 4 ... }
func decode_pattern(patterns []string) map[string]int {
	code := make(map[string]int) 
	rev := make(map[int]string) 

	// sort the patterns
	for i, patt := range patterns {
		patterns[i] = sort_pattern(patt)
	}
	
	// get the easy numbers
	for _, patt := range patterns {
		switch len(patt) {
		case 2:
			code[patt] = 1
			rev[1] = patt
		case 3:
			code[patt] = 7
			rev[7] = patt
		case 4:
			code[patt] = 4
			rev[4] = patt
		case 7:
			code[patt] = 8
			rev[8] = patt
		}
	}
	
	// solve the patterns of length 6
	// 9's segments are a superset of 4 (when 0 and 6 aren't)
	// 7's segments are not a superset of 7's (when 9 and 0s are)
	// the only remaining is 0
	var patt_save string
	for _, patt := range patterns {
		if len(patt) == 6 {	
			switch {
			case len(pattern_intersect(rev[4], patt)) == 4: 
				code[patt] = 9
				rev[9] = patt
			case len(pattern_intersect(rev[7], patt)) == 2:
				code[patt] = 6
				rev[6] = patt
			default:
				patt_save = patt
			}
		}
	}
	code[patt_save] = 0
	rev[0] = patt_save


	// solve the patterns of length 5
	for _, patt := range patterns {
		if len(patt) == 5 {
			switch {
			case len(pattern_intersect(rev[4], patt)) == 2:
				code[patt] = 2
				rev[2] = patt
			case len(pattern_intersect(rev[6], patt)) == 5:
				code[patt] = 5
				rev[5] = patt
			case len(pattern_intersect(rev[7], patt)) == 3:
				code[patt] = 3
				rev[5] = patt
			}
		}
	}

	return code
}

func solve_part2(signal_patterns, output_values [][]string) int {
	var sum int
	for i, pattern := range signal_patterns {
		code := decode_pattern(pattern)
		output := 0
		for _, o_val := range output_values[i] {
			ov := sort_pattern(o_val)
			output = code[ov] + output*10
		}
		fmt.Println(code, output_values[i], output)
		sum += output
	}
	return sum
}


func read_input() ([][]string, [][]string) {
	var signal_patterns, output_values [][]string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split_line := strings.Split(scanner.Text(), "|")
		signal_patterns = append(signal_patterns,
			strings.Split(strings.TrimRight(split_line[0], " "), " "))
		output_values = append(output_values,
			strings.Split(strings.TrimLeft(split_line[1], " "), " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return signal_patterns, output_values
}

func main() {
	
	signal_patterns, output_value := read_input()
	/*
	for i, n := range signal_patterns {
		fmt.Println(n, output_value[i])
	}
	*/
	fmt.Println("part1: ", solve_part1(signal_patterns, output_value))
	fmt.Println("part2: ", solve_part2(signal_patterns, output_value))
	
}
