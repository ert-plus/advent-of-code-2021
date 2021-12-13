package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strings"
)

type Vertex struct {
	//	idx int
	name string
	adj []*Vertex
	smol bool
}

func print_path(p []*Vertex) {
	if len(p) == 0 {
		fmt.Print("Path empty!")
		return
	}
	fmt.Print(p[0].name)
	for _, v := range p[1:] {
		fmt.Print(" -> " + v.name)
	}
	fmt.Print("\n")
}

func (v *Vertex) print() {
	if v.smol {
		fmt.Print("Small ")
	} else {
		fmt.Print("Large ")
	}
	fmt.Print("vertex ", v.name, ":\nAdjacent to: ")
	for _, u := range v.adj {
		fmt.Print(u.name, " ")
	}
	fmt.Print("\n")		
}

func (v *Vertex) check_adjacent(u *Vertex) bool {
	for _, w := range v.adj {
		if w == u {
			return true
		}
	}
	return false
}

func lookup_vertex(name string, g []*Vertex) *Vertex {
	for _, v := range g {
		if v.name == name {
			return v
		}
	}
	return nil
}

func add_edge(name1 string, name2 string, g []*Vertex) []*Vertex {
	/*
	fmt.Println("\nadding edge", name1, name2, "Before graph:")
	for _, u := range g {
		u.print()
	} */
	
	v1 := lookup_vertex(name1, g)
	v2 := lookup_vertex(name2, g)
	if v1 == nil {
		smol := strings.ToLower(name1) == name1
		v1 = &Vertex{name1, make([]*Vertex, 0), smol}
		g = append(g, v1)
	}
	if v2 == nil {
		smol := strings.ToLower(name2) == name2
		v2 = &Vertex{name2, make([]*Vertex, 0), smol}
		g = append(g, v2)
	}
	v1.adj = append(v1.adj, v2)
	v2.adj = append(v2.adj, v1)

	/*
	fmt.Println("returning graph")
	for _, u := range g {
		u.print()
	}
	*/
	return g
}

func read_input() []*Vertex {
	var output []*Vertex
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		verts := strings.Split(line, "-")
		//	fmt.Println("trying to get vert", verts)
		output = add_edge(verts[0], verts[1], output)
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func find_paths(start *Vertex, end *Vertex, graph []*Vertex, current []*Vertex,
	paths [][]*Vertex, check func(*Vertex, []*Vertex) bool) [][]*Vertex {

	if current == nil {
		current = make([]*Vertex, 0)
	}

	current = append(current, start)

	fmt.Print("start: ", start.name, " end: ", end.name)
	fmt.Println("\ncurrent path:")
	print_path(current)
	/*
	fmt.Println("paths:")
	for _, p := range paths {
		print_path(p)
	}
	fmt.Println("")
	*/
	
	if start == end {
		path := make([]*Vertex, len(current))
		copy(path, current)
		paths = append(paths, path)
		return paths
	} else {
		for _, u := range start.adj {
			if check(u, current) {
				paths = find_paths(u, end, graph, current, paths, check)
			}
		}
	}

	current = current[:len(current)-1]

	return paths
}

func part1_check(u *Vertex, p []*Vertex) bool {
	return ! u.smol || lookup_vertex(u.name, p) == nil
}

// this is slow without a path datatype to express a double cave visit
func part2_check(u *Vertex, p []*Vertex) bool {
	if ! u.smol {
		return true
	}
	
	if u.name == "start" {
		return false
	}
	repeats := make(map[*Vertex]int)
	double := false
	for _, v := range p {
		if _, ok := repeats[v]; ! ok{
			repeats[v] = 1
		} else {
			repeats[v] += 1
		}
	}
	for k, v := range repeats {
		if k.smol && v > 1 {
			double = true
		}
	}
	_, ok := repeats[u]
	return ! double || ! ok
}

func solve_part1(graph []*Vertex) int {
	start := lookup_vertex("start", graph)
	end := lookup_vertex("end", graph)
	current := make([]*Vertex, 0)
	paths := make([][]*Vertex, 0)
	paths = find_paths(start, end, graph, current, paths, part1_check)
	return len(paths)
}

func solve_part2(graph []*Vertex) int {
	start := lookup_vertex("start", graph)
	end := lookup_vertex("end", graph)
	current := make([]*Vertex, 0)
	paths := make([][]*Vertex, 0)
	paths = find_paths(start, end, graph, current, paths, part2_check)
	return len(paths)
}

func main() {
	graph := read_input()
	for _, v := range graph {
		v.print()
	}
	fmt.Println("part1: ", solve_part1(graph))
	fmt.Println("part2: ", solve_part2(graph))
}
