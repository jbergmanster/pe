package main

import (
	"fmt"
)

func main() {
	const max = 1001 * 1001
	sum := 1
	for x, y := 1, 2; x < max; y += 2 {
		for j := 0; j < 4; j++ {
			x += y
			sum += x
		}
	}
	fmt.Println(sum)
}
