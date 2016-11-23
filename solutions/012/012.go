package euler012

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("012", "Highly divisible triangular number", solve)
}

func solve(ctx *euler.Context) {
	bd := 0
	bn := 0
	t := triangle{}
	for {
		n := t.next()
		d := divisors(n)
		if d >= 500 {
			if euler.Verbose {
				fmt.Println(n, "with", d, "divisors")
			}
			answer := fmt.Sprintf("%d", n)
			ctx.SetAnswer(answer)
			break
		}

		// Show running progress.
		if d > bd {
			bd = d
			bn = n
			if euler.Verbose {
				fmt.Println("passed", bn, "with", bd, "divisors")
			}
		}
	}
}

// Find number of divisors for N.  This exploits
// the sqrt being the limit on divisors.
func divisors(n int) int {
	// Count 1 and the number itself
	count := 2

	// Only consider numbers up to the sqrt(n)
	sqrt := int(math.Sqrt(float64(n)))

	// If i < sqrt(n) and n%i == 0, then also count n/i as a factor
	// Example: n = 36, 2 < sqrt(36), covers divisors 2 and 18

	if n&1 == 0 {
		for i := 2; i <= sqrt; i++ {
			if n%i == 0 {
				count += 2
			}
		}

	} else {
		for i := 3; i <= sqrt; i += 2 {
			if n%i == 0 {
				count += 2
			}
		}
	}

	// Avoid overcounting when N is a perfect square,
	// e.g. sqrt(36) = 6  and  6 * 6 = 36
	if n == sqrt*sqrt {
		count--
	}
	return count
}

// Generate triangular numbers
type triangle struct {
	n int
	i int
}

func (t *triangle) next() int {
	t.n += t.i
	t.i++
	return t.n
}
