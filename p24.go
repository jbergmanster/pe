package main

import (
	"fmt"
)

func main() {
	head := make([]int, 0, 10)
	tail := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	permute(0, head, tail)
}

var count = 0

func permute(pos int, head, tail []int) {
	if len(tail) == 1 {
		count++
		if count == 1000000 {
			fmt.Println(head, tail)
		}
		return
	}
	head = head[:len(head)+1]
	for _, d := range tail {
		head[pos] = d
		permute(pos+1, head, remove(d, tail))
	}
	head = head[:len(head)-1]
}

func remove(d int, digits []int) []int {
	r := make([]int, 0, len(digits)-1)
	for _, c := range digits {
		if c != d {
			r = append(r, c)
		}
	}
	return r
}
