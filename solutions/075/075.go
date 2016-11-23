package euler075

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("075", "Singular integer right triangles", solve)
}

func solve(ctx *euler.Context) {
	count := calculate(uint64(1500000))
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

func calculate(limit uint64) int {
	counts := make([]int, limit+1)

	limiter := func(t *euler.Triple) bool {
		return t.Sum > limit
	}
	for t := range euler.PythagoreanTriplesBerggren(limiter) {
		counts[t.Sum]++

		// Scale up to get other triples until we exceed the limit
		f := uint64(2)
		for {
			x := t.Multiply(f)
			if x.Sum > limit {
				break
			}
			counts[x.Sum]++
			f++
		}
	}

	// Finally, count which N produced only 1 distinct Pythagorean triple
	count := 0
	for _, c := range counts {
		if c == 1 {
			count++
		}
	}
	return count
}
