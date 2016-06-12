package main

import (
	"fmt"
	"github.com/jbergmanster/pe/primes"
)

func main() {
	const max = 28123
	s := primes.CreateSet(max)

	var a []uint64
	m := make(map[uint64]bool)

	for i := uint64(12); i < max; i++ {
		if sumDivisors(s.Divisors(i))-i > i {
			a = append(a, i)
			for _, x := range a {
				if x != i {
					m[x+i] = true
				}
			}
		}
	}
	sum := uint64(0)
	for i := uint64(1); i < max; i++ {
		if !m[i] {
			sum += i
		}
	}
	fmt.Println(sum)
}

func sumDivisors(d []uint64) uint64 {
	sum := 0
	for _, x := range d {
		sum += x
	}
	return sum
}
