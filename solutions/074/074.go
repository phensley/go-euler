package euler074

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("074", "Digit factorial chains", solve)
}

func solve(ctx *euler.Context) {
	limit := 1000000
	count := 0
	s := newSolution()
	for n := 2; n < limit; n++ {
		if s.countTerms(n) == 60 {
			count++
		}
	}

	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

type solution struct {
	set             *euler.IntSet
	digitFactorials []int
}

func newSolution() *solution {
	// Precalculate the factorials for digits 0-9
	factorials := euler.DigitFactorials()
	return &solution{
		euler.NewIntSet(0, 256, 0.5),
		factorials,
	}
}

// Sum the factorials of the digits of N
func (s *solution) sumFactorial(n int) int {
	sum := 0
	for n >= 10 {
		sum += s.digitFactorials[n%10]
		n /= 10
	}
	if n > 0 {
		sum += s.digitFactorials[n]
	}
	return sum
}

// Count the non-repeating terms of the digit factorial chains
func (s *solution) countTerms(n int) int {
	p := n
	count := 1
	s.set.Clear()

	for {
		sum := s.sumFactorial(p)

		// Check if we've reached an earlier state
		if s.set.Contains(sum) {
			break
		}

		s.set.Add(sum)
		count++
		p = sum
	}
	return count
}
