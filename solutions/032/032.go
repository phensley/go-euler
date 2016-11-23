package euler032

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("032", "Pandigital products", solve)
}

func solve(ctx *euler.Context) {
	// Only use each digit once.
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Produce distinct permutations of digits
	perms := euler.NewPermutations(digits)

	// Numbers having the following digits produce results in the
	// valid range:
	//
	//     a x bbbb = cccc
	//     aa x bbb = cccc
	//
	// Example:  28 * 157 == 4396
	//
	slices := [][]int{
		[]int{1, 5},
		[]int{2, 5},
	}

	// Generate the products
	products := make(map[int]*struct{})
	for perms.Next() {
		p := perms.Get()
		for _, s := range slices {
			s1 := s[0]
			s2 := s[1]
			r, ok := calc(p, s1, s2)
			if ok {
				products[r] = &struct{}{}
			}
		}
	}
	sum := 0
	for v := range products {
		sum += v
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

// Multiple numbers represented by the digit slices
func calc(p []int, s1, s2 int) (int, bool) {
	a := euler.DigitsToInt(p[:s1])
	b := euler.DigitsToInt(p[s1:s2])
	c := euler.DigitsToInt(p[s2:])
	if a*b == c {
		return c, true
	}
	return 0, false
}
