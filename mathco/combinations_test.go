package mathco

import (
	"fmt"
	"testing"
)

func TestMultichoose(t *testing.T) {
	print := func(set []int) bool {
		fmt.Println(set)
		return true
	}
	Multichoose([]int{1, 3, 7, 9}, 4, print)
}

func TestLexicographicPermutations(t *testing.T) {
	print := func(set []int) bool {
		fmt.Println(set)
		return true
	}
	LexicographicPermutations([]int{1, 3, 3, 9}, print)
}
