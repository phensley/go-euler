package euler038

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("038", "Pandigital multiples", solve)
}

func solve(ctx *euler.Context) {
	// Produce a sequence of 1-9 pandigital numbers descending from the max
	indices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	digit := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	perms := euler.NewPermutations(indices)
	for perms.Next() {
		idx := perms.Get()
		d := []int{}
		for _, i := range idx {
			d = append(d, digit[i])
		}
		n := euler.DigitsToInt(d)
		if n%2 != 0 {
			continue
		}
		r := calculate(uint(n))
		if r > 0 {
			if euler.Verbose {
				fmt.Printf("(%d * 1) : (%d * 2) = %d:%d = %d\n", r, r, r, r*2, n)
			}
			answer := fmt.Sprintf("%d", n)
			ctx.SetAnswer(answer)
			return
		}
	}
}

// Largest must be a 4:5 slice, since our first 4 digits can (in theory)
// start with 9xxx and produce 5 digits when multiplied by 2. So we need
// to find the largest 4-digit number where the remaining 5-digit pandigital
// segment is twice as large.
func calculate(n uint) uint {
	d := euler.UintToDigits(n)
	s1 := euler.DigitsToUint(d[:4])
	s2 := euler.DigitsToUint(d[4:])
	if s1 == s2/2 {
		return s1
	}
	return 0
}
