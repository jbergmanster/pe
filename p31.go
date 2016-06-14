package main

import (
	"fmt"
)

func main() {
	coins := []int{200, 100, 50, 20, 10, 5, 2, 1}
	fmt.Println(count(coins, 200))
}

func count(coins []int, amt int) int {
	if len(coins) == 1 {
		return 1
	} else if amt == 0 {
		return 1
	}
	ans := 0
	x := amt / coins[0]
	if x == 0 {
		return count(coins[1:], amt)
	}
	for i := 0; i <= x; i++ {
		ans += count(coins[1:], amt-i*coins[0])
	}
	return ans
}
