package main

import (
	"fmt"
	"github.com/jbergmanster/pe/primes"
)

func main() {
	const max = 10000
	s := primes.CreateSet(max)
	m := make(map[uint64]uint64)
	for i := uint64(2); i < max; i++ {
		m[i] = sumProperDivisors(i, s.Divisors(i))
	}
	sum := uint64(0)
	for k, v := range m {
		if m[v] == k && k != v {
			fmt.Println(k, v)
			sum += k
		}
	}
	fmt.Println(sum)
}

func sumProperDivisors(x uint64, divisors []uint64) uint64 {
	sum := uint64(0)
	for _, y := range divisors {
		if y != x {
			sum += y
		}
	}
	return sum
}
