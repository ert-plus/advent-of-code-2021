package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strings"
	"sort"
)

func read_input() (string, map[string]string) {
	var template string
	rules := make(map[string]string)
	
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if template == "" {
			template = scanner.Text()
			continue
		}  
		line := scanner.Text()
		if line == "" {
			continue
		}
		rule := strings.Split(line," -> ")
		rules[rule[0]] = rule[1]
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return template, rules
}

func apply_rule(input string, rules map[string]string) string {
	var output string
	for i := 0; i < len(input)-1; i ++ {
		pair := input[i:i+2]
		if i == 0  {
			output += input[i:i+1] + rules[pair] + input[i+1:i+2] 
		} else {
			output += rules[pair] + input[i+1:i+2] 
		}
	}
	return output
}

func apply_rule_2(adj map[string]int, rules map[string]string) map[string]int {
	output := make(map[string]int)
	for k, _ := range adj {
		output[k] = 0
	}
	for k, v := range adj {
		left := k[:1] + rules[k]
		right := rules[k] + k[1:]
		//	fmt.Println("key:", k, "left:", left, "right:", right, "val:", v)
		// output[k] -= v
		output[left] += v
		output[right] += v
	}
	return output
}

func solve_part2(template string, rules map[string]string) int {
	//	fmt.Println(template)
	adj := make(map[string]int)
	elem := make(map[byte]int)
	
	for k, _ := range rules {
		adj[k] = 0
		elem[k[0]] = 0
		elem[k[1]] = 0
	}
	
	elem[template[0]] += 1
	elem[template[len(template)-1]] += 1
	
	for i := 0 ; i < len(template) - 1; i ++ {
		adj[template[i:i+2]] += 1
	}

	for step := 0 ; step < 40 ; step ++ {
		adj = apply_rule_2(adj, rules)
		//	fmt.Println(adj)
	}
	
	for k,v := range adj {
		elem[k[0]] += v
		elem[k[1]] += v
	}

	var counts []int 
	for _,v := range elem {
		counts = append(counts, v/2)
	}
	sort.Ints(counts)
	return counts[len(counts)-1] - counts[0]
}

	
func solve_part1(template string, rules map[string]string) int {
	//	fmt.Println(rules)
	// fmt.Println(template)
	polymer := template
	for step := 0; step < 10; step ++ {
		polymer = apply_rule(polymer, rules)
		// fmt.Print(step, " ")
		//	fmt.Println(polymer)
	}

	elem := make(map[byte]int)
	for i := 0; i < len(polymer); i ++ {
		elem[polymer[i]]++
	}

	lengths := make([]int, 0)
	for _, v := range elem {
		lengths = append(lengths,v)
	}
	sort.Ints(lengths)
	return lengths[len(lengths)-1] - lengths[0]
}

func main() {
	template, rules := read_input()
	fmt.Println("part1: ", solve_part1(template, rules))
	fmt.Println("part2: ", solve_part2(template, rules))
}
