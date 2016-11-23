package euler046

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("046", "Goldbach's other conjecture", solve)
}

func solve(ctx *euler.Context) {
	// Precalculate a quick way to check primality and a slice containing
	// the first N primes to use for calculations

	primeset := euler.PrimesSieveOfAtkin(10000)
	primes := primeset.Ones()

	// Start with the first composite odd number
	i := uint64(9)
	for {
		// Stop when we've checked all odd numbers up to our maximum prime
		ii := i
		if ii > primeset.Limit() {
			break
		}

		// Skip all primes
		if primeset.IsSet(ii) {
			i += 2
			continue
		}

		found := false
		for _, p := range primes {
			// Prime must be less than the odd number
			if p > i {
				break
			}

			// Subtract the prime
			r := i - p

			// Difference evenly divisible by 2?
			q := float64(r) / float64(2)
			if q != float64(int64(q)) {
				continue
			}

			// Is a perfect square?
			sq := math.Sqrt(q)
			if sq != float64(int64(sq)) {
				continue
			}

			// This one satisfies Goldbach's other conjecture. On to the next.
			found = true
			break
		}

		// We checked each prime up to the number, and none of them left a remainder
		// that was evenly divisible by 2, or when divided by 2 left a perfect square.
		if !found {
			answer := fmt.Sprintf("%d", i)
			ctx.SetAnswer(answer)
			return
		}
		i += 2
	}
}
