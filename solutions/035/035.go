package euler035

import (
	"fmt"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("035", "Circular primes", solve)
}

func solve(ctx *euler.Context) {
	// Compute the first million primes
	sieve := euler.PrimesSieveOfAtkin(1000000)
	primes := sieve.Ones()

	res := []int{}
	for _, n := range primes {
		// Get digits for the number so we can produce rotations
		d := euler.IntToDigits(int(n))

		// Check that all rotations of the number are prime
		f := true
		rotations := euler.NewRotations(d)

		// Skip over the first rotation which is identical
		rotations.Get()
		for rotations.Next() {
			// Get the number corresponding to the rotation
			r := rotations.Get()
			x := euler.DigitsToInt(r)

			// If not prime, this number fails
			if !sieve.IsSet(uint64(x)) {
				f = false
				break
			}
		}

		// Collect all primes whose rotations are also prime
		if f {
			res = append(res, int(n))
		}
	}

	if euler.Verbose {
		sort.Ints(res)
		fmt.Println("PRIMES: ", res)
	}

	answer := fmt.Sprintf("%d", len(res))
	ctx.SetAnswer(answer)
}
