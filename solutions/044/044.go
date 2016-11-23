package euler044

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("044", "Pentagon numbers", solve)
}

func solve(ctx *euler.Context) {
	set := euler.NewIntSet(0, 256*64, 0.85)
	nums := []int{}
	for i := 1; i < 10000; i++ {
		p := pentagon(i)
		set.Add(p)
		nums = append(nums, p)
	}

	for i, a := range nums {
		if i < 1 {
			continue
		}
		for _, b := range nums[:i] {
			d := a - b
			if set.Contains(d) && set.Contains(a+b) {
				answer := fmt.Sprintf("%d", d)
				ctx.SetAnswer(answer)
				return
			}
		}
	}
}

func pentagon(n int) int {
	return (n * ((3 * n) - 1)) / 2
}
