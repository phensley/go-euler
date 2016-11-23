package euler005

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("005", "Smallest multiple", solve)
}

func solve(ctx *euler.Context) {
	// The final answer must be divisible by the product of primes smaller
	// than 20. To speed up the search, multiply them together to get our
	// increment.
	incr := product(2, 3, 5, 7, 11, 13, 17, 19)

	// Since we know our increment is already divisible by the primes above,
	// no need to check them again. Only check the largest composite numbers
	// that are a multiple of our primes, e.g. 18 = 2 * 3 * 3, 20 = 2 * 2 * 5
	divisors := []uint{12, 14, 15, 16, 18, 20}

	// Find smallest number divisible by 1..20. This will take 24 steps exactly.
	num := smallestDivisible(incr, incr, divisors)

	if euler.Verbose {
		fmt.Println("smallest divisible by 1 to 20 is", num)
	}
	answer := fmt.Sprintf("%d", num)
	ctx.SetAnswer(answer)
}

// Compute a product of the given numbers
func product(nums ...uint) uint {
	res := uint(1)
	for _, n := range nums {
		res *= n
	}
	return res
}

// Find the smallest number divisible by the divisors, in steps of
// size incr.
func smallestDivisible(start, incr uint, divisors []uint) uint {
	num := start
	for {
		for _, i := range divisors {
			if num%i != uint(0) {
				goto CONTINUE
			}
		}
		return num

	CONTINUE:
		// Just in case there is no answer, check for overflow.
		n := num + incr
		if n < num {
			break
		}
		num = n
	}
	panic("failed!")
}
