package euler097

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("097", "Large non-Mersenne prime", solve)
}

func solve(ctx *euler.Context) {
	answer := fmt.Sprintf("%d", calculate())
	ctx.SetAnswer(answer)
}

func calculate() uint64 {
	scale := uint64(math.Pow(10, 10))
	power := 7830457

	// Compute the running last 10 digits while raising
	// 2 to the power of 7830457. We keep scaling down
	// the number by erasing all digits 10-billion or larger.
	n := uint64(2)
	for i := 1; i < power; i++ {
		if n >= scale {
			n -= (n / scale) * scale
		}

		n = n * 2
	}
	n *= 28433
	n++
	if n >= scale {
		n -= (n / scale) * scale
	}
	return n
}
