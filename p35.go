package main

import (
	"fmt"
	"github.com/jbergmanster/pe/primes"
)

func main() {
	ans := 4 // 2,3,5,7 are the one digit primes
	p := primes.CreateSet(1000000)
	visited := make(map[int]bool)

	checkRotations := func(set []int) {
		m := make(map[int]bool)
		l := len(set)
		allPrime := true
		for i := 0; i < l; i++ {
			x := 0
			for k, j := l-1, 1; k > -1; k, j = k-1, j*10 {
				x += j * set[(k+i)%l]
			}
			if visited[x] {
				return
			}
			if !p.IsPrime(uint64(x)) {
				allPrime = false
			}
			m[x] = true
		}
		if allPrime {
			ans += len(m)
		}
		for k := range m {
			visited[k]=true
		}
	}

	// Circular primes can only be made up of the following digits
	a := []int{1, 3, 7, 9}
	// Digit indexes array
	b := [6]int{0, 0, 0, 0, 0, 0}
	// Digits array
	c := [6]int{0, 0, 0, 0, 1, 1}
	last := [6]int{3, 3, 3, 3, 3, 3}
	blen := len(b)
	numdigits := 2
	for {
		checkRotations(c[blen-numdigits:])
		if b == last {
			break
		}
		for i := 1; i <= numdigits && numdigits <= blen; i++ {
			idx := blen - i
			b[idx] = (b[idx] + 1) % 4
			c[idx] = a[b[idx]]
			if b[idx] != 0 {
				break
			}
			if i == numdigits && b[idx] == 0 {
				numdigits++
			}
		}
	}
	fmt.Println(ans)
}
