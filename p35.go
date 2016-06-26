package main

import (
	"fmt"
	"github.com/jbergmanster/pe/mathco"
	"github.com/jbergmanster/pe/primes"
)

func main() {
	ans := 4 // 2,3,5,7 are the one digit primes
	p := primes.CreateSet(1000000)
	var multisets [][]int
	candidates := make(map[int][]int)

	addMultiset := func(set[] int) bool {
		s := make([]int, len(set), len(set))
		copy(s,set)
		multisets = append(multisets, s)
		return true
	}

	addPrimePermutation := func(set[] int) bool {
		x := 0
		for k, j := len(set)-1, 1; k > -1; k, j = k-1, j*10 {
			x += j * set[k]
		}
		if !p.IsPrime(uint64(x)) {
			return true
		}
		s := make([]int, len(set), len(set))
		copy(s,set)
		candidates[x] = s
		return true
	}

	checkRotations := func(set []int) {
		m := make(map[int]bool)
		l := len(set)
		for i := 0; i < l; i++ {
			x := 0
			for k, j := l-1, 1; k > -1; k, j = k-1, j*10 {
				x += j * set[(k+i)%l]
			}
			candidates[x]=nil
			if !p.IsPrime(uint64(x)) {
				return
			}
			m[x] = true
		}
		ans += len(m)
	}

	// Circular primes can only be made up of the following digits
	e := []int{1, 3, 7, 9}
	for i := 2; i < 7; i++ {
		mathco.Multichoose(e, i, addMultiset)
	}
	for _,set := range multisets {
		mathco.LexicographicPermutations(set, addPrimePermutation)
	}
	for key,_ := range candidates {
		if candidates[key] == nil {
			continue
		}
		checkRotations(candidates[key])
	}
	fmt.Println(ans)
}


