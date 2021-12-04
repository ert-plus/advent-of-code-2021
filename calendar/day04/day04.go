package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

type Bingo struct {
	Board [5][5]int
	Marks [5][5]bool
}

func new_bingo() *Bingo {
	var b [5][5]int
	var m [5][5]bool
	out := Bingo{b, m}
	return &out
}

func read_input(filename string) ([]int, []*Bingo) {
	var drawn_nums []int
	var bingos []*Bingo
	var next_bingo *Bingo
	
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan() ; i++ {
		// if it's the first line, fill out drawn_nums
		if drawn_nums == nil && i == 0 {
			draw_line := strings.Split(scanner.Text(), ",")
			for _, draw := range draw_line {
				n, err := strconv.Atoi(draw)
				if err != nil {
					log.Fatal(err)
				}
				drawn_nums = append(drawn_nums, n)
			}
			continue
		}

		line := scanner.Text()
		if line == "" {
			next_bingo = new_bingo()
			i = -1 // so the next row is zero indexed
			continue
		}

		row := strings.Fields(line)
		for j, r := range row {
			n, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			next_bingo.Board[i][j] = n
		}

		if (i == 4) {
			bingos = append(bingos, next_bingo)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return drawn_nums, bingos
}


func (b *Bingo) mark_bingo(num int) {
	for i := 0; i < 5; i ++ {
		for j :=0; j < 5; j ++ {
			if b.Board[i][j] == num {
				b.Marks[i][j] = true
			}
		}
	}
}

// returns true if bingo has a bingo
func (b *Bingo) check_bingo() bool {
	var across, down, out bool
	// check every row and column, if we don't get 5 in a row, across + down are false
	for i := 0; i < 5; i++ {
		across, down = true, true
		for j := 0; j < 5; j ++ {
			across = across && b.Marks[i][j]
			down = down && b.Marks[j][i]
		}
		out = out || across || down
	}
	return out
}

func (b *Bingo) sum_unmarked() int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0 ; j < 5; j++ {
			if ! b.Marks[i][j] {
				sum += b.Board[i][j]
			}
		}
	}
	return sum
}

func solve_part1(drawn_nums []int, bingos []*Bingo) int {
	for _, n := range drawn_nums {
		for _, b := range bingos {
			b.mark_bingo(n)
			if b.check_bingo() {
				fmt.Println("at drawn number: ", n)
				fmt.Println("with a winning board")
				b.print()
				return n * b.sum_unmarked()
			}
		}
	}
	return -1
}

func filter_bingos(bingos []*Bingo) []*Bingo {
	var out []*Bingo
	for _, b := range bingos {
		if ! b.check_bingo() {
			out = append(out, b)
		}
	}
	return out
}

func solve_part2(drawn_nums []int, bingos []*Bingo) int {
	for _, n := range drawn_nums {
		for _, b := range bingos {
			b.mark_bingo(n)
			if len(bingos) == 1 && b.check_bingo() {
				fmt.Println("at drawn number: ", n)
				fmt.Println("with a winning board")
				b.print()
				return n * b.sum_unmarked()
			}
		}
		bingos = filter_bingos(bingos)
	}
	return -1
}

func (b *Bingo) print() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Marks[i][j] {
				fmt.Printf("[%2d] ", b.Board[i][j])
			} else {
				fmt.Printf("%4d ", b.Board[i][j])
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	
	drawn_nums, bingos := read_input("input.txt")
	/*
	for _, n := range drawn_nums {
		fmt.Println(n)
	}
	for i, bingo := range bingos {
		fmt.Println("bingo #", i)
		for j := 0; j < 5; j++ {
			fmt.Println(bingo.Board[j][0], bingo.Board[j][1], 
				bingo.Board[j][2], bingo.Board[j][3], bingo.Board[j][4])
		}
	}
	*/
	fmt.Println("part1: ", solve_part1(drawn_nums, bingos))
	/*
	for i, b := range bingos {
		fmt.Println("Board #", i)
		b.print()
	}
	*/
	fmt.Println("part2: ", solve_part2(drawn_nums, bingos))
}
