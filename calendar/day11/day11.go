package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

type Point struct {
	X,Y int
}

type Octopus struct {
	Energy int
	Flash bool
	Loc Point
}

func get_adjacent(p Point) []Point {
	var adj []Point	
	left, right, top, bottom :=
		p.X == 0,
		p.X == 9,
		p.Y == 0,
		p.Y == 9

	if ! left {
		adj = append(adj, Point{p.X-1, p.Y})
	}
	if ! right {
		adj = append(adj, Point{p.X+1, p.Y})
	}
	if ! top {
		adj = append(adj, Point{p.X, p.Y-1})
	}
	if ! bottom {
		adj = append(adj, Point{p.X, p.Y+1})
	}
	if ! left && ! top {
		adj = append(adj, Point{p.X-1, p.Y-1})
	}
	if ! left && ! bottom {
		adj = append(adj, Point{p.X-1, p.Y+1})
	}
	if ! right && ! top {
		adj = append(adj, Point{p.X+1, p.Y-1})
	}
	if ! right && ! bottom {
		adj = append(adj, Point{p.X+1, p.Y+1})
	}
	
	return adj
}

func read_input() [][]Octopus {
	output := make([][]Octopus, 10)
	for i := 0 ; i < 10; i ++ {
		output[i] = make([]Octopus, 10)
	}
	
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i ++ {
		line := scanner.Text()
		for j, c := range line {
			n, _ := strconv.Atoi(string(c))
			output[i][j] = Octopus{n, false, Point{j, i}}
			// output = append(output, Depth{j, i, n})
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    return output
}


func solve_part2(octopi [][]Octopus) int {
	var count int
	
	for step := 0; step < 1000; step ++ {
		for x := 0; x < 10; x ++ {
			for y := 0; y < 10; y ++ {
				octopi[y][x].Energy ++
			}
		}
		count = 0
		for new_flashes := true ; new_flashes ; {
			new_flashes = false
			for x := 0; x < 10; x ++ {
				for y := 0; y < 10; y ++ {
					if octopi[y][x].Energy > 9 && ! octopi[y][x].Flash {
						new_flashes = true
						octopi[y][x].Flash = true
						count ++ 
						adj := get_adjacent(octopi[y][x].Loc)
						for _, a := range adj {
							octopi[a.Y][a.X].Energy ++
						}
					}
				}
			}
		}
		if count == 100 {
			return step + 1
		}
		for x := 0; x < 10; x ++ {
			for y := 0; y < 10; y ++ {
				if octopi[y][x].Flash {
					octopi[y][x].Energy = 0
					octopi[y][x].Flash = false
				}
			}
		}
	}
	return -1
}

func solve_part1(octopi [][]Octopus) int {
	var count int
	
	for step := 0; step < 100; step ++ {
		for x := 0; x < 10; x ++ {
			for y := 0; y < 10; y ++ {
				octopi[y][x].Energy ++
			}
		}
		for no_new_flashes := true ; no_new_flashes ; {
			no_new_flashes = false
			for x := 0; x < 10; x ++ {
				for y := 0; y < 10; y ++ {
					if octopi[y][x].Energy > 9 && ! octopi[y][x].Flash {
						no_new_flashes = true
						octopi[y][x].Flash = true
						count ++ 
						adj := get_adjacent(octopi[y][x].Loc)
						for _, a := range adj {
							octopi[a.Y][a.X].Energy ++
						}
					}
				}
			}
		}

		for x := 0; x < 10; x ++ {
			for y := 0; y < 10; y ++ {
				if octopi[y][x].Flash {
					octopi[y][x].Energy = 0
					octopi[y][x].Flash = false
				}
			}
		}
	}
	
	return count
}

func main() {
	octopi := read_input()
	fmt.Println("part1: ", solve_part1(octopi))
	octopi = read_input()
	fmt.Println("part2: ", solve_part2(octopi))
}
