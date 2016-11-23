package euler031

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("031", "Coin sums", solve)
}

func solve(ctx *euler.Context) {
	coins := []int{200, 100, 50, 20, 10, 5, 2, 1}
	s := euler.NewBigSummation()
	count := s.Compute(200, coins)
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}
