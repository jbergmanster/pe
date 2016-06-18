package main

import (
	"fmt"
)

func main() {
	pows := [10]int{
		0,
		1,
		2 * 2 * 2 * 2 * 2,
		3 * 3 * 3 * 3 * 3,
		4 * 4 * 4 * 4 * 4,
		5 * 5 * 5 * 5 * 5,
		6 * 6 * 6 * 6 * 6,
		7 * 7 * 7 * 7 * 7,
		8 * 8 * 8 * 8 * 8,
		9 * 9 * 9 * 9 * 9}

	ans := -1 // Because we don't count 1.
	for x1 := 0; x1 < 10; x1++ {
		for x2 := 0; x2 < 10; x2++ {
			for x3 := 0; x3 < 10; x3++ {
				for x4 := 0; x4 < 10; x4++ {
					for x5 := 0; x5 < 10; x5++ {
						for x6 := 0; x6 < 10; x6++ {
							sum := pows[x1] + pows[x2] + pows[x3] + pows[x4] + pows[x5] + pows[x6]
							num := 100000*x1 + 10000*x2 + 1000*x3 + 100*x4 + 10*x5 + x6
							if sum == num {
								ans += num
								fmt.Println(sum, num)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
