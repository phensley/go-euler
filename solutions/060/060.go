package euler060

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("060", "Prime pair sets", solve)
}

func solve(ctx *euler.Context) {
	n := uint64(9000)
	s := newSolution(n, n*100)
	ctx.SetAnswer(s.calculate(5))
}

type prime struct {
	n uint
	d []uint
}

type solution struct {
	primeset *euler.Bitstring
	primes   []uint

	// Larger set of primes for fast checking
	checkset *euler.Bitstring
}

func newSolution(limit, checkLimit uint64) *solution {
	primeset := euler.PrimesSieveOfAtkin(limit)
	checkset := euler.PrimesSieveOfAtkin(checkLimit)
	primes := make([]uint, primeset.Count())
	for i, p := range primeset.Ones() {
		primes[i] = uint(p)
	}
	return &solution{primeset, primes, checkset}
}

func (s *solution) isPrime(n uint) bool {
	if uint64(n) < s.checkset.Limit() {
		return s.checkset.IsSet(uint64(n))
	}
	return euler.IsPrimeForisekJancina32(uint32(n))
}

// Concatenate the digits of A and B to produce a new number.
// Risk of overflow here when number of digits in the sum exceeds
// size of uint
func concat(a, b uint) uint {
	// Number of base-10 digits in b
	d := euler.NumDigitsBase10(uint64(b))
	// Multiplier shift d digits to the left in base 10
	m := uint(math.Pow(float64(10), d))
	return (a * m) + b
}

func (s *solution) check(a, b uint) bool {
	return a != b && s.isPrime(concat(a, b)) && s.isPrime(concat(b, a))
}

func (s *solution) calculate(size int) string {
	// First, create lists of sets of 2 primes that check out.
	// The next phase will be to grow these until the set contains 5
	sets := [][]uint{}
	for i, p1 := range s.primes {
		for _, p2 := range s.primes[i+1:] {
			if !s.check(p1, p2) {
				continue
			}
			sets = append(sets, []uint{p1, p2})
		}
	}

	// Track the smallest sum we've found for a 5-set
	smallest := uint(0)
	smallestSet := []uint{}

	// Grow each set until we find the first 5-set
	for _, set := range sets {
		// Larger of the two primes
		larger := set[1]

		for _, p1 := range s.primes {
			// Skip primes smaller than the largest prime already in the set
			if p1 <= larger {
				continue
			}

			// Flag indicating whether this prime checked against all
			// primes already in the set.
			add := true
			for _, p2 := range set {
				if !s.check(p1, p2) {
					add = false
					break
				}
			}

			// Add this prime to the set
			if add {
				set = append(set, p1)
			}
		}

		// Once we hit 5 this is the smallest sum, since we checked
		// primes in increasing order
		if len(set) == size {
			sum := uint(0)
			for _, n := range set {
				sum += n
			}
			smallest = sum
			smallestSet = set
			break
		}

		// Indicates we should break out
		if smallest > 0 {
			break
		}
	}

	if euler.Verbose {
		fmt.Println(smallest, "  set: ", smallestSet)
	}

	return fmt.Sprintf("%d", smallest)
}
