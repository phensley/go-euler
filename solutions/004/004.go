package euler004

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("004", "Largest palindrone product", solve)
}

func solve(ctx *euler.Context) {
	largest := 0
	limit := 999
	i := limit
	for i >= 100 {
		// Finds the first palindrome for this value of i
		for j := limit; j >= i; j-- {
			n := i * j
			if euler.IsPalindrome(uint(n)) {
				if largest < n {
					largest = n
				}
				// Break, since j decreases for this value of i, and
				// won't produce a larger candidate.
				break
			}
		}
		i--
	}
	answer := fmt.Sprintf("%d", largest)
	ctx.SetAnswer(answer)
}
