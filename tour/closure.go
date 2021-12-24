package main

import "fmt"

func fibonacci() func() int {
	prev := 0
	curr := 1
	return func() int {
		next := curr + prev
		prev, curr = curr, next
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0 ;  i < 10 ; i ++ {
		fmt.Println(f())
	}
}
