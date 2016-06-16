package main

import (
	"fmt"
	"github.com/jbergmanster/pe/primes"
	"sort"
)

func main() {
	const max = 100
	const min = 2
	s := primes.CreateSet(max + 1)
	m := make(map[[2][3]uint64]bool)
	for a := uint64(min); a <= max; a++ {
		f := s.Factor(a)
		k := mapToKey(f)
		for b := uint64(min); b <= max; b++ {
			c := k
			for i, x := range k[1] {
				c[1][i] = x * b
			}
			m[c] = true
		}
	}
	fmt.Println(len(m))
}

func mapToKey(m map[uint64]uint) [2][3]uint64 {
	var key [2][3]uint64
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for i, k := range keys {
		key[0][i] = uint64(k)
		key[1][i] = uint64(m[uint64(k)])
	}
	return key
}
