package euler087

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("087", "Prime power triples", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		nums := compute(50)
		fmt.Println("known", nums.Ones())
	}

	nums := compute(50000000)
	answer := fmt.Sprintf("%d", nums.Count())
	ctx.SetAnswer(answer)
}

func compute(lim uint64) *euler.Bitstring {
	// Generate all primes up to sqrt(limit)
	root := euler.NthRootNewtonFloat64(float64(lim), 2, 20)
	primes := euler.PrimesSieveOfAtkin(uint64(root)).Ones()

	// Generate powers of primes for exponents 2, 3, 4
	powers2 := primePowers(primes, 2, lim)
	powers3 := primePowers(primes, 3, lim)
	powers4 := primePowers(primes, 4, lim)

	// Mark all numbers that are a sum of one from the p^2, p^3, p^4 sequences
	nums := euler.NewBitString(lim)

	// Scan p^2 sequence
	for _, a := range powers2 {

		// Sum with p^3 sequence
		for _, b := range powers3 {
			ab := a + b
			// Skip ranges that exceed our limit
			if ab > lim {
				continue
			}

			// Sum with p^4 sequence
			for _, c := range powers4 {
				abc := ab + c
				if abc > lim {
					break
				}
				// Mark all sums that are < limit
				nums.Set(uint64(abc))
			}
		}
	}

	// Return the marked numbers
	return nums
}

// Generate a sequence containing powers of prime numbers up to the limit
func primePowers(primes []uint64, power, lim uint64) []uint64 {
	nroot := euler.NthRootNewtonFloat64(float64(lim), float64(power), 50)
	res := []uint64{}
	for _, p := range primes {
		// Break when we find a prime larger than the nth root of the limit
		if p > uint64(nroot) {
			break
		}
		a := p
		for i := uint64(0); i < power-1; i++ {
			a *= p
		}
		if a > uint64(lim) {
			break
		}
		res = append(res, a)
	}
	return res
}
