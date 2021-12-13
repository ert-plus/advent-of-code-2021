package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Fold struct {
	along string
	val int
}

func read_input() ([][]bool, []Fold) {
	grid := make([][]bool, 1500)
	for y, _ := range grid {
		grid[y] = make([]bool, 1500)
	}
	folds := make([]Fold, 0)
	
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cords := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			cords = false
			continue
		}

		if cords {
			cord := strings.Split(line, ",")
			x, _ := strconv.Atoi(cord[0])
			y, _ := strconv.Atoi(cord[1])
			grid[y][x] = true
		} else {
			fold := strings.Split(strings.Split(line," ")[2], "=")
			val, _ := strconv.Atoi(fold[1])
			folds = append(folds, Fold{fold[0], val})
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return grid, folds
}

func apply_fold(grid [][]bool, fold Fold) [][]bool {

	for y, row := range grid {
		for x, _ := range row {
			switch fold.along {
			case "y":
				if y >= fold.val {
					grid[y][x] = false
				} else {
					yreflect := fold.val - y
					if grid[fold.val + yreflect][x] {
						grid[fold.val - yreflect][x] = true
					}
				}
			case "x":
				if x >= fold.val {
					grid[y][x] = false
				} else {
					xreflect := fold.val - x
					if grid[y][fold.val + xreflect] {
						grid[y][fold.val - xreflect] = true
					}
				}
			default:
				log.Fatal("invalid fold along val:", fold.along)
			}
		}
	}
	return grid
}

func solve_part2(grid [][]bool, folds []Fold) int {

	for _, fold := range folds {
		grid = apply_fold(grid, fold)
	}

	for y, row := range grid {
		if y > 6 {
			break
		}
		for x, val := range row {
			if x > 40 {
				break
			}
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}	
		}
		fmt.Print("\n")
	}
	
	return 0
}

func solve_part1(grid [][]bool, folds []Fold) int {
	grid = apply_fold(grid, folds[0])

	var count int
	for y, row := range grid {
		for x, _ := range row {
			/*
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			*/
			if grid[y][x] {
				count ++
			}
		}
		//	fmt.Print("\n")
	}
	return count
}

func main() {
	grid, folds := read_input()
	//	fmt.Println(folds)
	fmt.Println("part1: ", solve_part1(grid, folds))
	fmt.Println("part2: ", solve_part2(grid, folds))
}
