package euler009

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("009", "Pythagorean triplet", solve)
}

func solve(ctx *euler.Context) {
	answer := triples(1000)
	if answer != "" {
		ctx.SetAnswer(answer)
	}
}

func triples(n uint) string {
	// Generate the set of distinct values of (a,b,c) that
	// sum exactly to N.  We're only interested in which of these
	// is a Pythagorean triple.

	// Start off with (1, 1, 998)
	a := uint(1)
	b := uint(1)
	c := n - (a + b)

	for {
		// Check if this is a Pythagorean triplet. The first one we find
		// is the answer.
		z := a*a + b*b
		if z == c*c {
			if euler.Verbose {
				fmt.Printf("Found: a=%d b=%d c=%d\n", a, b, c)
			}
			return fmt.Sprintf("%d", a*b*c)
		}

		// Generate the sequence:
		//    (1,     1, 998)
		//    (1,     2, 997)
		//    ..
		//    (1,   499, 500)
		//    (2,     2, 996)
		//    ..
		//    (333, 333, 334)
		b++
		c = n - (a + b)
		if c < b {
			// Increment a and test again
			a++
			b = a
			c = n - (a + b)

			// Once both a and b reach c we've finished.
			if c < b {
				break
			}
		}
	}
	return ""
}
