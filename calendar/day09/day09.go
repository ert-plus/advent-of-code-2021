package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"sort"
)

type Depth struct {
	X, Y, Value int
}

func read_input() [][]int {
	output := make([][]int, 100)
	for i := 0 ; i < 100; i ++ {
		output[i] = make([]int, 100)
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
			output[i][j] = n
			// output = append(output, Depth{j, i, n})
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func cord_contains(cords [][2]int, x,y int) bool {
	cmp := [2]int{x,y}
	for _, c := range cords {
		if cmp == c {
			return true
		}
	}
	return false
}

// given an xy cord of a low point, finds the size of basin
// 
func find_basin_size(depths [][]int, x,y int) int {
	var basin_cords, edge_cords [][2]int
	// basin_cords := make([][2]int, 100)
	// edge_cords := make([][2]int, 100)
	edge_cords = append(edge_cords, [2]int{x,y})
	lenx, leny := len(depths[0]), len(depths)
	var p [2]int
	
	for len(edge_cords) > 0 {
		p, edge_cords = edge_cords[0], edge_cords[1:]
		x, y = p[0], p[1]
		basin_cords = append(basin_cords, [2]int{x,y})
		/*
		   fmt.Println("x", x, "y", y)
		fmt.Println("edge", edge_cords)
		fmt.Println("basin", basin_cords)
		*/
		if x != 0 && depths[y][x-1] != 9 &&
			! cord_contains(basin_cords, x-1, y) &&
			! cord_contains(edge_cords, x-1, y)	{
			edge_cords = append(edge_cords, [2]int{x-1, y})
		}
		if x < lenx-1 && depths[y][x+1] != 9 &&
			! cord_contains(basin_cords, x+1, y) &&
			! cord_contains(edge_cords, x+1, y) {
			edge_cords = append(edge_cords, [2]int{x+1, y})
		}
		if y != 0 && depths[y-1][x] != 9 &&
			! cord_contains(basin_cords, x, y-1) &&
			! cord_contains(edge_cords, x, y-1) {
			edge_cords = append(edge_cords, [2]int{x, y-1})
		}
		if y < leny-1 && depths[y+1][x] != 9 &&
			! cord_contains(basin_cords, x, y+1) &&
			! cord_contains(edge_cords, x, y+1) {
			edge_cords = append(edge_cords, [2]int{x, y+1})
		}
	}
	//	fmt.Println(basin_cords)
	return len(basin_cords)
}

func solve_part2(depths [][]int) int {
	leny := len(depths)
	lenx := len(depths[0])
	var sizes []int
	for i := 0; i < leny; i++ {
		for j := 0; j < lenx; j ++ {
			low_point := true
			low_point = low_point && (i == 0 || depths[i-1][j] > depths[i][j])
			low_point = low_point && (i == leny-1 || depths[i+1][j] > depths[i][j])
			low_point = low_point && (j == 0 || depths[i][j-1] > depths[i][j])
			low_point = low_point && (j == lenx-1 || depths[i][j+1] > depths[i][j])
			if low_point {
				// fmt.Println("low point", j, i, depths[i][j])
				sizes = append(sizes, find_basin_size(depths, j, i))
			}
		}
	}
	//	fmt.Println(sizes)
	sort.Ints(sizes)
	sl := len(sizes)
	return sizes[sl-1] * sizes[sl-2] * sizes[sl-3]
}

func solve_part1(depths [][]int) int {
	var risk_sum int
	leny := len(depths)
	lenx := len(depths[0])
	for i := 0; i < leny; i++ {
		for j := 0; j < lenx; j ++ {
			low_point := true
			low_point = low_point && (i == 0 || depths[i-1][j] > depths[i][j])
			low_point = low_point && (i == leny-1 || depths[i+1][j] > depths[i][j])
			low_point = low_point && (j == 0 || depths[i][j-1] > depths[i][j])
			low_point = low_point && (j == lenx-1 || depths[i][j+1] > depths[i][j])
			if low_point {
				// fmt.Println("low point", j, i, depths[i][j])
				risk_sum += depths[i][j] + 1
			}
		}
	}
	return risk_sum 
}

func main() {
	depths := read_input()
	// 	fmt.Println(depths)
	fmt.Println("part1: ", solve_part1(depths))
	fmt.Println("part2: ", solve_part2(depths))
}
