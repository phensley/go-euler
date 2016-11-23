package euler002

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("002", "Even Fibonacci numbers", solve)
}

func solve(ctx *euler.Context) {
	sum := uint64(0)
	for n := range euler.FibonacciSequence(1, 1) {
		if n >= 4000000 {
			break
		}
		if n%2 == 0 {
			sum += n
		}
	}
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}
