package euler037

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("037", "Truncatable primes", solve)
}

func solve(ctx *euler.Context) {
	limit := uint64(1000000)
	sieve := euler.PrimesSieveOfAtkin(limit)
	primes := sieve.Ones()
	sum := uint64(0)

outer:
	for _, p := range primes {
		if p <= 7 {
			continue
		}

		d := euler.UintToDigits(uint(p))

		// Remove left digits and test primality
		for i := 0; i < len(d); i++ {
			n := euler.DigitsToUint(d[i:])
			if !sieve.IsSet(uint64(n)) {
				continue outer
			}
		}

		// Remove right digits and test primality
		for i := len(d) - 1; i > 0; i-- {
			n := euler.DigitsToUint(d[:i])
			if !sieve.IsSet(uint64(n)) {
				continue outer
			}
		}

		sum += p
		if euler.Verbose {
			fmt.Println(p)
		}
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}
