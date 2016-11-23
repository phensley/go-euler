package euler076

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("076", "Counting summations", solve)
}

func solve(ctx *euler.Context) {
	answer := fmt.Sprintf("%d", calculate(100))
	ctx.SetAnswer(answer)
}

func calculate(n int) *big.Int {
	numbers := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		numbers[i] = i + 1
	}
	s := euler.NewBigSummation()
	return s.Compute(n, numbers)
}
