package main

import (
	"fmt"
)

func main() {
	const max = 2540160
	f := [10]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	ans := 0
	for i := 10; i < max; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			d := j % 10
			sum += f[d]
		}
		if sum == i {
			fmt.Println(i)
			ans += sum
		}
	}
	fmt.Println(ans)
}
