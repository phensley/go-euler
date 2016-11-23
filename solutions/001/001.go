package euler001

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("001", "Multiples of 3 and 5", solve)
}

func solve(ctx *euler.Context) {
	sum := sumMultiplesBrute(1000, []int{3, 5})
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func sumMultiplesBrute(limit int, nums []int) int {
	sum := 0
	for n := 1; n < limit; n++ {
		for _, m := range nums {
			if n%m == 0 {
				sum += n
				break
			}
		}
	}
	return sum
}
