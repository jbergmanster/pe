package main

import (
	"fmt"
	"github.com/jbergmanster/pe/primes"
)

func main() {
	const max = 1000000
	s := primes.Set{}
	s.Sieve(max)
	var aMax, bMax, lMax int64
	for a := int64(-999); a < 1000; a++ {
		for b := int64(-999); b < 1000; b++ {
			if !s.IsPrime(abs(b)) {
				continue
			}
			for n := int64(0); ; n++ {
				x := n*n + a*n + b
				if !s.IsPrime(abs(x)) {
					if n > lMax {
						lMax = n
						aMax = a
						bMax = b
					}
					break
				}
			}
		}
	}
	fmt.Println(aMax, bMax, lMax, aMax*bMax)
}

func abs(x int64) uint64 {
	if x < 0 {
		return uint64(-1 * x)
	}
	return uint64(x)
}
