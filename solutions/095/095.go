package euler095

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("095", "Amicable chains", solve)
}

func solve(ctx *euler.Context) {
	answer := fmt.Sprintf("%d", calculate())
	ctx.SetAnswer(answer)
}

func calculate() int {
	limit := 1000000
	sums := euler.ProperDivisorSumSieve(limit)
	seen := newIntArray(1024)
	longest := 0
	longestMin := 0

	for n := 2; n < limit; n++ {
		count := 1
		seen.truncate()
		seen.append(n)
		p := n
		for {
			p = sums[p]

			// Don't count chains of length 1 or those that include elements
			// above the limit
			if p == 1 || p > limit {
				break
			}

			// Check if we've seen this number before.
			idx := seen.find(p)
			if idx != -1 {
				// Slice out the amicable chain and compare
				chain := seen.n[idx:seen.i]
				length := len(chain)
				if longest < length {
					longest = length
					longestMin = minimum(chain)
				}
				break
			}

			seen.append(p)
			count++
		}

	}
	return longestMin
}

func minimum(n []int) int {
	m := math.MaxInt32
	for i := 0; i < len(n); i++ {
		if n[i] < m {
			m = n[i]
		}
	}
	return m
}

// Dynamic int array
type intarray struct {
	i int
	n []int
}

func newIntArray(initial int) *intarray {
	return &intarray{0, make([]int, initial)}
}

func (a *intarray) Len() int {
	return a.i
}

func (a *intarray) append(n int) {
	a.grow()
	a.n[a.i] = n
	a.i++
}

func (a *intarray) find(n int) int {
	for i := 0; i < a.i; i++ {
		if n == a.n[i] {
			return i
		}
	}
	return -1
}

func (a *intarray) grow() {
	length := len(a.n)
	if a.i < length {
		return
	}

	n := make([]int, length*2)
	for i := 0; i < a.i; i++ {
		n[i] = a.n[i]
	}
	a.n = n
}

func (a *intarray) truncate() {
	a.i = 0
}
