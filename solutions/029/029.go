package euler029

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("029", "Distinct powers", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		count := compute(5)
		fmt.Println("Known: # terms where 2 <= (a,b) <= 5: ", count)
	}

	count := compute(100)
	answer := fmt.Sprintf("%v", count)
	ctx.SetAnswer(answer)
}

func compute(limit int) int {
	// Find all numbers that can be represented as a power of a smaller number
	// For example, 16^53 == 2^212
	powers := make(map[int]*power)
	for i := 2; i <= limit; i++ {
		n := i * i
		f := 2
		for n <= limit {
			if powers[n] == nil {
				powers[n] = &power{i, f}
			}
			n *= i
			f++
		}
	}

	// Count all numbers in terms of powers
	seen := make(map[power]bool)
	count := 0
	x := power{0, 0}
	for n := 2; n <= limit; n++ {
		f := powers[n]
		for p := 2; p <= limit; p++ {

			// If the number can be represented in terms of a power of
			// a smaller number, do so.
			if f == nil {
				x.n = n
				x.p = p
			} else {
				x.n = f.n
				x.p = f.p * p
			}

			// Check if this power has been seen; if not, count it
			if !seen[x] {
				count++
				seen[x] = true
			}
		}
	}
	return count
}

type power struct {
	n int
	p int
}
