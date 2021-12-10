package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strings"
	"sort"
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

func solve_part2(lines []string) int {
	var scores []int
	var score_table = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4 }	
	
	for _, line := range lines {
		var opener rune
		illegal := false
		var chunks []rune
		for _, c := range line {
			if strings.ContainsRune("({<[", c) {
				chunks = append(chunks, c)
			} else {
				if len(chunks) == 0 {
					illegal = true
					break
				} else if len(chunks) == 1 {
					opener, chunks = chunks[0], make([]rune, 0)
				} else {
					opener, chunks = chunks[len(chunks)-1], chunks[:len(chunks)-1]
				}
				switch opener {
				case '(':
					if c != ')' {
						illegal = true
						break
					}
				case '{':
					if c != '}' {
						illegal = true
						break
					}
				case '[':
					if c != ']' {
						illegal = true
						break
					}
				case '<':
					if c != '>' {
						illegal = true
						break
					}
				}
			}	
		}
		if ! illegal {
			var score int
			for i, _ := range chunks {
				score *= 5
				score += score_table[chunks[len(chunks)-(i+1)]]
			}
			scores = append(scores, score)
			/*
			fmt.Println("got incomplete:")
			fmt.Println(line)
			fmt.Println("got opening chunks:")
			fmt.Println(string(chunks))
			fmt.Println("with score:", score)
		} else {
			fmt.Println("got corrupted:")
			fmt.Println(line)
			*/
		}
	}

	sort.Ints(scores)
	
	return scores[len(scores)/2]
}




func solve_part1(lines []string) int {
	var chunks []rune
	var score int
	var score_table = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137 }	
	
	for _, line := range lines {
		var illegal, opener rune
		for _, c := range line {
			if strings.ContainsRune("({<[", c) {
				chunks = append(chunks, c)
			} else {
				if len(chunks) == 0 {
					opener = 0
				} else if len(chunks) == 1 {
					opener, chunks = chunks[0], make([]rune, 0)
				} else {
					opener, chunks = chunks[len(chunks)-1], chunks[:len(chunks)-1]
				}
				switch opener {
				case '(':
					if c != ')' {
						illegal = c
						break
					}
				case '{':
					if c != '}' {
						illegal = c
						break
					}
				case '[':
					if c != ']' {
						illegal = c
						break
					}
				case '<':
					if c != '>' {
						illegal = c
						break
					}
				case 0:
					illegal = c
					break
				}
			}	
		}
		if illegal != 0 {
			//			fmt.Println("found illegal char:", string(illegal), "for line")
			//			fmt.Println(line)
			score += score_table[illegal]
		}
	}
	return score
}

func main() {
	lines := read_input()
	fmt.Println("part1: ", solve_part1(lines))
	fmt.Println("part2: ", solve_part2(lines))
}
