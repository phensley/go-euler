package euler041

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("041", "Pandigital prime", solve)
}

func solve(ctx *euler.Context) {
	for _, length := range []int{9, 8, 7, 6, 5, 4, 3, 2} {
		r := permutations(length)
		if r != 0 {
			answer := fmt.Sprintf("%d", r)
			ctx.SetAnswer(answer)
			return
		}
	}
}

func permutations(length int) int {
	indices := make([]int, length)
	digits := make([]int, length)
	for i := 0; i < length; i++ {
		indices[i] = i
		digits[i] = length - i
	}

	d := make([]int, length)
	perms := euler.NewPermutations(indices)
	for perms.Next() {
		p := perms.Get()
		for i, idx := range p {
			d[i] = digits[idx]
		}
		n := euler.DigitsToInt(d)
		if euler.IsPrimeForisekJancina32(uint32(n)) {
			return n
		}
	}
	return 0
}
