package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const max = 1000
	longest := 0
	ans := 0
	ans2 := ""
	for i := 1; i < max; i++ {
		s := divide(1, i)
		if x := strings.Index(s, "("); x != -1 {
			l := len(s) - x - 1
			if l > longest {
				longest = l
				ans = i
				ans2 = s
			}
		}
	}
	fmt.Println(longest)
	fmt.Println(ans)
	fmt.Println(ans2)
}

func divide(num, den int) string {
	x := num
	y := den
	ans := bytes.NewBufferString("")
	dec := false
	seen := make(map[int]int)
	for {
		r := 10 * (x % y)
		if _, ok := seen[r]; ok {
			s := ans.String()
			return s[0:seen[r]-1] + "(" + s[seen[r]-1:] + ")"
		}
		ans.WriteString(strconv.Itoa(x / y))
		if r == 0 {
			break
		}
		if !dec {
			ans.WriteString(".")
			dec = true
		} else {
			seen[r] = ans.Len()
		}
		x = r
	}
	return ans.String()
}
