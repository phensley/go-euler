package euler027

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("027", "Quadratic primes", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		// Try some known values
		z := countPrimes(1, 41)
		fmt.Println("f40 =", z)
		z = countPrimes(-79, 1601)
		fmt.Println("f80 =", z)
	}

	// Construct several quadratic functions of the form n^2 + an + b
	// and find the values for {a, b} which produce the maximum number
	// of primes for consecutive values of n
	maxa := 0
	maxb := 0
	maxc := 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			c := countPrimes(a, b)
			if c > maxc {
				maxa = a
				maxb = b
				maxc = c
			}
		}
	}

	if euler.Verbose {
		fmt.Println("a=", maxa, "b=", maxb, "produced", maxc, "consecutive primes")
	}
	answer := fmt.Sprintf("%d", maxa*maxb)
	ctx.SetAnswer(answer)
}

// Count the number of consecutive primes produced by the given function
func countPrimes(a, b int) int {
	n := 0
	for {
		// Call the quadratic function for increasing n and stop
		// when we find the first non-prime.
		r := (n * n) + (a * n) + b
		if !euler.IsPrimeForisekJancina32(uint32(r)) {
			break
		}
		n++
	}
	return n
}
