package main

import (
	"fmt"
)

func main() {
	var ans int64
	for i := 1; i <= 20; i++ {
		if i%1000 == 0 {
			fmt.Println("At ", i)
		}
		ans += findIntegralAreas(i)
	}
	fmt.Println(ans)
}

// a is a vertex at (0, 0)
// b is a vertex at (0, area)
// c is a vertex at (2, 0)
// d is the point between a and b
// e is a point between a and c
// f is the intersection of bd and ce
//
// Algorithm is to check if bdf is integral.
// The other areas will be integral if bdf is by construction.
func findIntegralAreas(area int) int64 {
	var sum int64
	found := make(map[[4]int]bool)

	for d := 2; d < area-1; d++ {
		for abe := 2; abe < area-1; abe++ {
			bd := area - d
			numer := bd * bd * abe
			denom := area*area - d*abe
			if numer%denom == 0 {
				bdf := numer / denom
				adfe := abe - bdf
				bfc := bd - bdf
				fec := area - bdf - adfe - bfc
				areas := [4]int{bdf, adfe, bfc, fec}
				// sort areas for the map
				for i := 1; i < 4; i++ {
					for j := i; j > 0; j-- {
						if areas[j] >= areas[j-1] {
							break
						}
						x := areas[j]
						areas[j] = areas[j-1]
						areas[j-1] = x
					}
				}
				if !found[areas] {
					sum += int64(area)
					found[areas] = true
				}

			}
		}
	}
	return sum
}
