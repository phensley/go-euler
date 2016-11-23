package euler

import (
	"fmt"
	"math/big"
)

// BigSummation caches values for a given summation
type BigSummation struct {
	cache map[string]*big.Int
}

// NewBigSummation builds a new BigSummation object
func NewBigSummation() *BigSummation {
	return &BigSummation{map[string]*big.Int{}}
}

var (
	bigOne = big.NewInt(1)
)

// Compute calculates the ways the given numbers can sum up to the total
func (b *BigSummation) Compute(total int, numbers []int) *big.Int {
	length := len(numbers)
	if length == 0 {
		return nil
	}

	n := numbers[0]
	key := fmt.Sprintf("%d:%d", total, n)
	e := b.cache[key]
	if e != nil {
		return e
	}

	sum := big.NewInt(0)

	times := total / n
	for i := 0; i <= times; i++ {
		r := total - (n * i)
		if r != 0 {
			z := b.Compute(r, numbers[1:])
			if z != nil {
				sum.Add(sum, z)
			}
		} else {
			sum.Add(sum, bigOne)
		}
	}

	e = big.NewInt(0)
	b.cache[key] = e.Add(e, sum)
	return e
}

// Summation caches values for a given summation
type Summation struct {
	cache map[string]int
}

// NewSummation builds a new summation object
func NewSummation() *Summation {
	return &Summation{map[string]int{}}
}

// Compute calculates the ways the given numbers can sum up to the total
func (s *Summation) Compute(total int, numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return 0
	}

	// Sum of the number of ways the numbers can add up to x
	sum := 0

	// Try the first value in the list
	n := numbers[0]

	key := fmt.Sprintf("%d:%d", total, n)
	if s.cache[key] != 0 {
		return s.cache[key]
	}

	// Compute how many times this number can evenly divide into the total
	times := total / n
	for i := 0; i <= times; i++ {
		// Check if there is a remainder for the current number
		r := total - (n * i)
		if r != 0 {
			// Compute the ways the remaining numbers in the list can
			// add up to the remainder. We do this recursively until
			// all the smaller numbers are accounted for.
			sum += s.Compute(r, numbers[1:])
		} else {
			sum++
		}
	}
	if sum < 0 {
		panic(fmt.Sprintf("Summation exceeded int range. Use BigSummation!"))
	}

	s.cache[key] = sum
	return sum
}
