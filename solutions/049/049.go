package euler049

import (
	"fmt"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("049", "Prime permutations", solve)
}

func solve(ctx *euler.Context) {
	// Ignore the known value
	known := []int{1487, 4817, 8147}

	m := make(map[string][]int)
	primes := euler.PrimesSieveOfAtkin(10000).OneIterator()

	// Collect all 4-digit primes and group by their common digits
	for primes.Next() {
		p := primes.Get()
		if p < 1000 {
			continue
		}

		// Grouping primes by shared digits
		d := euler.IntToDigits(int(p))
		sort.Ints(d)
		k := fmt.Sprintf("%d", euler.DigitsToInt(d))

		m[k] = append(m[k], int(p))
	}

	// Check each group of primes and print those that return a valid result.
	for _, values := range m {
		if r := check(values); r != nil && !equal(r, known) {
			answer := ""
			for _, e := range r {
				answer += fmt.Sprintf("%d", e)
			}
			ctx.SetAnswer(answer)
			return
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func contains(n int, values []int) bool {
	for _, v := range values {
		if n == v {
			return true
		}
	}
	return false
}

func check(values []int) []int {
	if len(values) < 3 {
		return nil
	}

	sort.Ints(values)

	// Compute the deltas between the primes
	deltas := []int{}
	for i := 1; i < len(values); i++ {
		deltas = append(deltas, values[i]-values[i-1])
	}

	// Starting with each value in order, add the deltas and check
	// if the resulting number is one of the target primes.  If
	// both (n + d) and (n + d + d) are members of the set, we've
	// found a result.
	for i := 0; i < len(values); i++ {
		v := values[i]
		d := 0
		for j := i; j < len(deltas); j++ {
			d += deltas[j]
			if contains(v+d, values) && contains(v+d+d, values) {
				return []int{v, v + d, v + d + d}
			}
		}
	}
	return nil
}
