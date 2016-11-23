package euler047

import (
	"fmt"
	"os"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("047", "Distinct prime factors", solve)
}

func solve(ctx *euler.Context) {
	s := construct(150000)
	if euler.Verbose {
		s.compute(2, 2)
		s.compute(3, 3)
	}

	// Solve the main problem
	n := s.compute(4, 4)
	if n == 0 {
		os.Exit(0)
	}

	answer := fmt.Sprintf("%d", n)
	ctx.SetAnswer(answer)
}

type state struct {
	primes   []uint64
	primeset *euler.Bitstring
}

func construct(maxPrime uint64) *state {
	primeset := euler.PrimesSieveOfAtkin(maxPrime)
	s := &state{
		primes:   primeset.Ones(),
		primeset: primeset,
	}
	return s
}

// Return true if the number N has M distinct prime factors (target)
func (s *state) factors(n uint64, target int) bool {
	if s.primeset.IsSet(n) {
		return false
	}

	count := 0
	for _, p := range s.primes {
		if n%p == 0 {
			count++
			if count > target {
				return false
			}

			// Divide out this prime completely
			n /= p
			for n%p == 0 {
				n /= p
			}
		}
		if n == 1 {
			break
		}
	}

	return count == target
}

func (s *state) compute(factors, consecutive int) uint64 {
	// Starting value is smallest possible result
	i := uint64(1)
	for j, p := range s.primes {
		if j == consecutive {
			break
		}
		i *= p
	}
	limit := s.primeset.Limit()
	for i < limit {
		// If (i + consecutive) does not satisfy our goal, i cannot be the first number
		// in a sequence of consecutive integers. This lets us skip ahead several places.
		r := true
		for j := i + uint64(consecutive) - 1; j >= i; j-- {
			if s.primeset.IsSet(uint64(j)) || !s.factors(j, int(factors)) {
				// Start our next check at j + 1
				i = j + 1
				r = false
				break
			}
		}

		// If not found, i is already set, so continue
		if !r {
			continue
		}

		// We found an answer.
		if euler.Verbose {
			fmt.Println(consecutive, "consecutive integers having", factors, "distinct prime factors")
			for j := uint64(0); j < uint64(consecutive); j++ {
				fmt.Println("  ", i+j)
			}
		}
		return i
	}
	fmt.Println("Error: scanned up to maximum prime", limit, ". no answer found. bailing out.")
	return 0
}
