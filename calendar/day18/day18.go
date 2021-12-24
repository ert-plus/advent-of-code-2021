package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

type snailNum struct {
	left, right *snailNum
	val int 
}

func (sn *snailNum) print() {
	if sn.left == nil || sn.right == nil {
		fmt.Print(sn.val)
	} else {
		fmt.Print("[")
		sn.left.print()
		fmt.Print(",")
		sn.right.print()
		fmt.Print("]")
	}	
}

func parseSnailNum(s string) (*snailNum, string) {
	out := new(snailNum)
	for s != "" {
		c := s[0]
		switch c {
		case '[':
			out.left, s = parseSnailNum(s[1:])
		case ']':
			return out, s[1:]
		case ',':
			out.right, s = parseSnailNum(s[1:])
		default: // it's a number
			out.val = int(c-'0')
			return out, s[1:]
		}
	}
	return out, ""
}

var parents = make(map[*snailNum]*snailNum)

// return the first snailnum that should be exploded
func (sn *snailNum) explodeDFS(depth int) *snailNum {
	var out *snailNum
	if sn.left != nil {
		parents[sn.left] = sn
		out = sn.left.explodeDFS(depth+1)
		if out != nil {
			return out
		}
	}
	
	if depth == 4 && sn.left != nil && sn.right != nil {
		return sn
	}
	
	if sn.right != nil {
		parents[sn.right] = sn
		out = sn.right.explodeDFS(depth+1)
		if out != nil {
			return out
		}
	}
	return nil
}

func (sn *snailNum) isLeaf() bool {
	return sn.left == nil && sn.right == nil
}

func (sn *snailNum) addLeft(toAdd int) {
	curr :=  sn

	for {
		if parents[curr] == nil {
			return
		} else if parents[curr].left != curr {
			curr = parents[curr].left
			break
		}
		curr = parents[curr]
	}
	for curr.right != nil {
		curr = curr.right
	}
	curr.val = curr.val + toAdd
}

func (sn *snailNum) addRight(toAdd int) {
	curr :=  sn

	for {
		if parents[curr] == nil {
			return
		} else if parents[curr].right != curr {
			curr = parents[curr].right
			break
		}
		curr = parents[curr]
	}
	for curr.left != nil {
		curr = curr.left
	}
	curr.val = curr.val + toAdd
}

// check and explode if need be, returning if an explode happened
func (sn *snailNum) explode() bool {
	toExplode := sn.explodeDFS(0)

	if toExplode != nil {

		toExplode.addLeft(toExplode.left.val)
		toExplode.addRight(toExplode.right.val)
		

		toExplode.val = 0
		toExplode.left = nil
		toExplode.right = nil

		return true
	}
	return false
}

// check and split if need be, returning if a split happened
func (sn *snailNum) split() bool {
	out := false
	if sn.left != nil {
		out = sn.left.split()
		if out {
			return out
		}
	}

	if sn.val >= 10 {
		bigger, littler := sn.val/2, sn.val/2
		if bigger + littler != sn.val {
			bigger += 1
		}
		
		sn.val = 0
		sn.left = new(snailNum)
		sn.left.val = littler
		sn.right = new(snailNum)
		sn.right.val = bigger
		return true
	}
	
	if sn.right != nil {
		out = sn.right.split()
		if out {
			return out
		}
	}
	
	return false
}

func (sn *snailNum) reduce() {
	new := true
	for new {
		new = sn.explode()
		new = new || sn.split()

		// sn.print()
		// fmt.Println()
	}
}

func (sn *snailNum) magnitude() int {
	if sn.isLeaf() {
		return sn.val
	}
	return sn.left.magnitude()*3 + sn.right.magnitude()*2
}

func addSnailNums(s1 *snailNum, s2 *snailNum) *snailNum{
	out := new(snailNum)
	out.left = s1
	out.right = s2
	out.reduce()
	return out	
}

func readInput() []*snailNum {
	var output []*snailNum
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sn, _ := parseSnailNum(scanner.Text())
		output = append(output, sn)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func (sn *snailNum) copy() *snailNum {
	out := new(snailNum)
	out.val = sn.val
	if sn.left != nil {
		out.left = sn.left.copy()
	}
	if sn.right != nil {
		out.right = sn.right.copy()
	}
	return out	
}

func solvePart2(nums []*snailNum) int {
	max := 0
	for i, _ := range(nums) {
		for j, _ := range(nums) {
			if i == j {
				continue
			}
			one, two := nums[i].copy() , nums[j].copy()
			mag := addSnailNums(one,two).magnitude()
			if mag > max {
				max = mag
			}
		}
	}
	return max
}

func solvePart1(nums []*snailNum) int {
	var curr *snailNum
	if len(nums) != 0 {
		curr = nums[0]
	}
	for _, n := range nums[1:] {
		curr = addSnailNums(curr, n)
		//	fmt.Println()
	}
	return curr.magnitude()
}

func main() {
	nums := readInput()
	fmt.Println("part1: ", solvePart1(nums))
	nums = readInput()
	fmt.Println("part2: ", solvePart2(nums))
}
