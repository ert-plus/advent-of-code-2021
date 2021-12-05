package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

type Point struct {
	X, Y int
}

type Segment struct {
	Start, End Point
}

func read_input() []Segment {

	var segments []Segment
	
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")
		p1 := strings.Split(split[0], ",")
		p2 := strings.Split(split[1], ",")
		p1x, _ := strconv.Atoi(p1[0])
		p1y, _ := strconv.Atoi(p1[1])
		p2x, _ := strconv.Atoi(p2[0])
		p2y, _ := strconv.Atoi(p2[1])
		segments = append(segments, Segment{Point{p1x, p1y}, Point{p2x, p2y}})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(nil)
	}

	return segments
}

func solve_part1(segments []Segment) int {
	var grid [1000][1000]int

	for _, segment := range segments {
		// fmt.Println("segment: ", segment.Start, " -> ", segment.End)
		diffx := segment.End.X - segment.Start.X
		diffy := segment.End.Y - segment.Start.Y
		if diffx == 0 {
			signy := 1
			if diffy < 0 {
				signy = -1
			}
			for i := segment.Start.Y; i != segment.End.Y + signy; i += signy {
				grid[segment.Start.X][i] += 1
			}
		}

		if diffy == 0 {
			signx := 1
			if diffx < 0 {
				signx = -1
			}
			for i := segment.Start.X; i != segment.End.X + signx; i += signx {
				grid[i][segment.Start.Y] += 1
			}
		}
		if diffx != 0 && diffy != 0 {
			fmt.Println("skipping...")
		}
	}

	/*
	var count int
	f, _ := os.Create("/tmp/data")
	defer f.Close()
	*/
	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				count ++
			}
		}
	}
	/* 
	for i := 0; i < 1000; i ++ {
		for j := 0; j < 1000; j ++ {
			s := strconv.Itoa(grid[i][j])
			f.WriteString(s + " ")
		}
		f.WriteString("\n")
	}
	*/		
	return count
}

func solve_part2(segments []Segment) int {
	var grid [1000][1000]int

	for _, segment := range segments {
		fmt.Println("segment: ", segment.Start, " -> ", segment.End)
		diffx := segment.End.X - segment.Start.X
		diffy := segment.End.Y - segment.Start.Y

		signy := 1
		if diffy < 0 {
			signy = -1
		}
		signx := 1
		if diffx < 0 {
			signx = -1
		}
		
		if diffx == 0 {
			for i := segment.Start.Y; i != segment.End.Y + signy; i += signy {
				grid[segment.Start.X][i] += 1
			}
			continue
		}
		if diffy == 0 {
			for i := segment.Start.X; i != segment.End.X + signx; i += signx {
				grid[i][segment.Start.Y] += 1
			}
			continue
		}
		x, y := segment.Start.X, segment.Start.Y
		for ; x != segment.End.X + signx && y != segment.End.Y + signy;
		x, y = signx + x, signy + y  {
			grid[x][y] += 1
		}
	}

	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				count ++
			}
		}
	}
		
	return count

}


func main() {
	segments := read_input()
	for i, s := range segments {
		fmt.Println("seg", i, ": ", s.Start.X, ",", s.Start.Y,
			" to ", s.End.X, ",", s.End.Y)
	}
	fmt.Println("part1: ", solve_part1(segments))
	fmt.Println("part2: ", solve_part2(segments))
}
