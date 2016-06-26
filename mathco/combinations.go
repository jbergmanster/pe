// Package matchco contains various functions for combinatorics.
//
// For example it includes functions for generating permutations and
// combinations.
package mathco

import (
	//"fmt"
	"sort"
)

// Callback is the type of the callbacks from the various functions that iterate
// combinations and permutations.
type Callback func(set []int) bool

// Mulitchoose iterates all the multisets of size k of elements.
// It calls f for each multiset.
func Multichoose(elements []int, k int, f Callback) {
	idx := make([]int, k, k)
	set := make([]int, k, k)
	for {
		c := k - 1
		for i := 0; i < k; i++ {
			set[i] = elements[idx[i]]
		}
		if !f(set) {
			return
		}
		if idx[0] == len(elements)-1 {
			break
		} else if idx[c] < len(elements)-1 {
			idx[c] += 1
		} else if c != 0 {
			for ; idx[c] == len(elements)-1; c-- {
			}
			idx[c] += 1
			for i := c + 1; i < k; i++ {
				idx[i] = idx[c]
			}
			c = k - 1
		}
	}
}

// LexicographicPermutations implements the algorithm described in
// https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order.
// It calls the Callback f for each permutation.
func LexicographicPermutations(a sort.IntSlice, f Callback) {
	sort.Ints(a)
	f(a)
	for {
		// 1. Find largest k such that a[k] < a[k+1]
		k := -1
		for i := 0; i < len(a)-1; i++ {
			if a[i] < a[i+1] {
				k = i
			}
		}
		if k == -1 {
			return
		}
		// 2. Find largest l > k such that a[k] < a[l]
		l := -1
		for i := k + 1; i < len(a); i++ {
			if a[k] < a[i] {
				l = i
			}
		}
		// 3. Swap a[k] and a[l]
		a.Swap(k, l)
		// 4. Reverse a[k+1] up to final element a[n]
		for i, j := k+1, len(a)-1; i < j; i, j = i+1, j-1 {
			a.Swap(i, j)
		}
		f(a)
	}
}
