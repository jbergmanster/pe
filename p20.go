package main

import (
	"fmt"
)

func main() {
	var p []int
	p = append(p, 1)

	for i := 2; i <= 100; i++ {
		p = mul(i, p)
	}

	ans := 0
	for _, x := range p {
		ans += x
	}
	fmt.Println(ans)
}

func mul(x int, y []int) []int {
	var ans []int
	for offset, digit := range getDigits(x) {
		for j, c := 0, 0; j < len(y) || c != 0; j++ {
			x := c
			if j < len(y) {
				x = y[j]*digit + c
			}
			c = x / 10
			if offset+j < len(ans) {
				ans[offset+j] += x % 10
			} else {
				ans = append(ans, x%10)
			}
		}
	}
	return ans
}

func getDigits(i int) []int {
	var digits []int
	for {
		digits = append(digits, i%10)
		if i/10 == 0 {
			break
		}
		i /= 10
	}
	return digits
}
