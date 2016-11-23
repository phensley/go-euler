package euler058

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("058", "Spiral primes", solve)
}

func solve(ctx *euler.Context) {
	// Current corner value in the spiral as we build it.
	c := uint(1)

	// Current side length, starting with the first layer in the spiral
	side := uint(3)

	// Count of corners that are prime
	count := 0

	// Count of total corners, starting with the center '1'
	total := 1

	// Prime/total ratio below which we stop and print the side length
	limit := float64(0.10)

	check := func(n uint) {
		if euler.IsPrimeForisekJancina32(uint32(n)) {
			count++
		}
	}

	for {
		// We're only interested in the numbers at the corners of each layer in the spiral.
		// Skip to the next corner value and check it's primality.
		c += side - 1
		check(c)
		c += side - 1
		check(c)
		c += side - 1
		check(c)

		// The last corner is always a perfect square, so skip the primality check
		c += side - 1

		// Calculate ratio of prime corners to total corners
		total += 4
		ratio := float64(count) / float64(total)
		if ratio < limit {
			answer := fmt.Sprintf("%d", side)
			ctx.SetAnswer(answer)
			return
		}

		// Increase side length
		side += 2
	}

}
