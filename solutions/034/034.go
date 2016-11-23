package euler034

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("034", "Digit factorials", solve)
}

func solve(ctx *euler.Context) {
	// Pre-calculate factorials
	factorials := euler.DigitFactorials()

	sumFactorial := func(n int) int {
		sum := 0
		for n >= 10 {
			sum += factorials[n%10]
			n /= 10
		}
		if n > 0 {
			sum += factorials[n]
		}
		return sum
	}

	// Find the upper bound
	limit := 9
	for sumFactorial(limit) > limit {
		limit = (limit * 10) + 9
	}

	// Find numbers which equal the sum of their digit factorials
	sum := 0
	for n := 3; n <= limit; n++ {
		r := sumFactorial(n)
		if n == r {
			sum += r
		}
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}
