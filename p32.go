package main

import (
	"fmt"
)

func main() {
	sum := 0
	ps := make(map[int]bool)
	for a := 1; a < 10; a++ {
		for b := 1; b < 10; b++ {
			for c := 1; c < 10; c++ {
				for d := 1; d < 10; d++ {
					for e := 1; e < 10; e++ {
						if b == a || c == a || d == a || e == a ||
							c == b || d == b || e == b ||
							d == c || e == c || e == d {
							continue
						}
						p23 := (a*10 + b) * (c*100 + d*10 + e)
						p14 := a * (b*1000 + c*100 + d*10 + e)
						if p23 < 10000 {
							if ps[p23] == false && isPandigital(p23, a, b, c, d, e) {
								fmt.Printf("%d%d * %d%d%d = %d\n", a, b, c, d, e, p23)
								sum += p23
								ps[p23] = true

							}
						}
						if p14 < 10000 {
							if ps[p14] == false && isPandigital(p14, a, b, c, d, e) {
								fmt.Printf("%d * %d%d%d%d = %d\n", a, b, c, d, e, p14)
								sum += p14
								ps[p14] = true

							}
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func isPandigital(p, a, b, c, d, e int) bool {
	m := make(map[int]bool)
	m[a] = true
	m[b] = true
	m[c] = true
	m[d] = true
	m[e] = true
	for i := 0; i < 5; i++ {
		digit := p % 10
		if m[digit] == true {
			break
		}
		m[digit] = true
		p /= 10

	}
	if len(m) == 10 {
		return true
	}
	return false
}
