package main

import (
	"os"
	"bufio"
	"fmt"
	//	"strconv"
	"log"
)

type Vertex struct {
	weight int
	dist int
	adj []*Vertex
	prev *Vertex
	visited bool
}

func read_input() []*Vertex {
	var input []string
	var output []*Vertex
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
		for i := 0; i < len(line); i ++ {
			output = append(output, new(Vertex))
		}
	}

	for i, line := range input {
		for j, c := range line {
			v := output[j+i*len(line)]
			v.weight = int(c - '0')
			if i != 0 {
				a := output[j+(i-1)*len(line)]
				v.adj = append(v.adj, a)
			}
			if i != len(line)-1 {
				a := output[j+(i+1)*len(line)]
				v.adj = append(v.adj, a)
			}
			if j != 0 {
				a := output[(j-1)+i*len(line)]
				v.adj = append(v.adj, a)
			}
			if j != len(line)-1 {
				a := output[(j+1)+i*len(line)]
				v.adj = append(v.adj, a)
			}
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}


func read_input2() []*Vertex {
	var input []string
	var output []*Vertex
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
		for i := 0; i < len(line)*25; i ++ {
			output = append(output, new(Vertex))
		}
	}

	for i, line := range input {
		for j, c := range line {
			for m, _ := range []int{0,1,2,3,4} {  // copy graph x val
				for n, _ := range []int{0,1,2,3,4} { // copy graph y val
					// lines are 5 times as long now
					//         [ x value        , y value                    ]
					v := output[j + len(line)*m + 5*len(line)*(i + len(line)*n)]
					weight :=  (((int(c - '0') + m + n - 1) % 9) + 1)
					v.weight = weight
					if j + len(line)*m != 0 {
						a := output[(j + len(line)*m - 1) +  5*len(line)*(i + len(line)*n)]
						v.adj = append(v.adj, a)
					}
					if j + len(line)*m != 5*len(line) - 1 {
						a := output[(j + len(line)*m + 1) +  5*len(line)*(i + len(line)*n)]
						v.adj = append(v.adj, a)
					}
					if  5*len(line)*(i + len(line)*n) != 0 {
						a := output[j + len(line)*m +  5*len(line)*(i + len(line)*n - 1)]
						v.adj = append(v.adj, a)
					}
					if  (i + len(line)*n) != 5*len(line)-1 {
						a := output[j + len(line)*m +  5*len(line)*(i + len(line)*n + 1)]
						v.adj = append(v.adj, a)
					}
				}
			}
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

// broke: copying and pasting from stack overflow
// woke: copying and pasting from wikipedia
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode
func (s *Vertex) dijkstra(graph []*Vertex) {
	for _, v := range graph {
		v.dist = -1
	}
	s.dist = 0
	
	for {
		var min *Vertex
		// damn they weren't kidding about making a set for this
		// realizing in part two that this is slow as I iterate here
		// over existing vertices.

		// welp who cares! it runs on my laptop in about a minute and I need to sleep
		for _, v := range graph {
			switch {
			case v.dist == -1:
				continue
			case min == nil && ! v.visited:
				min = v
			case ! v.visited && v.dist < min.dist:
				min = v
			default:
				continue
			}
		}
		if min == nil { // no more vertices to visit!
			break
		}
		min.visited = true
		
		for _, neighbor := range min.adj {
			if neighbor.visited {
				continue
			}
			alt := min.dist + neighbor.weight
			if alt < neighbor.dist || neighbor.dist == -1 {
				neighbor.dist = alt
				neighbor.prev = min
			}
		}
	}
}

func solve_part2(graph []*Vertex) int {
	graph[0].dijkstra(graph)
	dst := graph[249999]
	return dst.dist
}

func solve_part1(graph []*Vertex) int {
	graph[0].dijkstra(graph)
	dst := graph[9999]
	return dst.dist
}

func main() {
	graph := read_input()

	/*
	for i, v := range s {
		// fmt.Print(v.weight)
		if (i % 100) == 0 {
			fmt.Printf("\n")
		}
		fmt.Println(v)
	}
	*/
	fmt.Println("part1: ", solve_part1(graph))

	big_graph := read_input2()
	/*
	for i, v := range big_graph {
		fmt.Print(v.weight)
		if (i % 50) == 49 {
			fmt.Printf("\n")
		}
		//fmt.Println(v)
	}
	*/
	//	fmt.Println("len:", len(big_graph))
	fmt.Println("part2: ", solve_part2(big_graph))
}
