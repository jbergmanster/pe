package main

import (
	"fmt"
	"math/big"
)

func main() {
	max := new(big.Int)
	max.Exp(big.NewInt(10), big.NewInt(999), nil)
	x := big.NewInt(1)
	y := big.NewInt(1)
	z := new(big.Int)
	i := 2
	for ; y.Cmp(max) < 0; i++ {
		z.Set(y)
		y.Add(x, y)
		x.Set(z)
	}
	fmt.Println(y.String())
	fmt.Println(i)
}
