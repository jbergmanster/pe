package main

import (
	"fmt"
)

// Just prints out the fractions. Doesn't find lowest common terms form.
func main() {
	for x1 := 1; x1 < 10; x1++ {
		for x2 := 1; x2 < 10; x2++ {
			// Eliminates trivial cases
			if x1 == x2 {
				continue
			}
			y1 := x2
			for y2 := 1; y2 < 10; y2++ {
				a := float64(x1) / float64(y2)
				b := float64(x1*10+x2) / float64(y1*10+y2)
				if a == b {
					fmt.Printf("%d%d/%d%d\n", x1, x2, y1, y2)
				}
			}
		}
	}
}
