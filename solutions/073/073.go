package euler073

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("073", "Counting fractions in a range", solve)
}

const (
	// Upper and lower bounds for the reduced fractions we are counting
	oneHalf  = 0.5
	oneThird = 1.0 / 3.0
)

func solve(ctx *euler.Context) {
	limit := uint64(12000)
	primes := euler.PrimesSieveOfAtkin(limit)
	if euler.Verbose {
		fmt.Println("Known: ", calculate(primes, uint64(8)))
	}

	result := calculate(primes, limit)
	answer := fmt.Sprintf("%d", result)
	ctx.SetAnswer(answer)
}

func calculate(primes *euler.Bitstring, limit uint64) int {
	count := 0

	// Cover all divisors up to limit
	for d := uint64(2); d <= limit; d++ {

		dPrime := primes.IsSet(d)

		// Reduce the number of comparisons by scanning only the range
		//  1/3 < n < 1/2 for this divisor
		lo := uint64(math.Floor(float64(d) / 3.0))
		hi := uint64(math.Floor(float64(d) / 2.0))
		for n := lo; n <= hi; n++ {

			// Count only reduced fractions
			if dPrime || primes.IsSet(n) || euler.GreatestCommonDivisor(n, d) == 1 {

				// Count only those between 1/3 and 1/2
				f := float64(n) / float64(d)
				if f > oneThird && f < oneHalf {
					count++
				}
			}
		}
	}
	return count
}
