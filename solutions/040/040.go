package euler040

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("040", "Champernowne's constant", solve)
}

func solve(ctx *euler.Context) {
	limit := 1000000

	n := 1
	digits := []int{}
	for {
		for _, d := range euler.IntToDigits(n) {
			digits = append(digits, d)
		}
		if len(digits) > limit {
			break
		}
		n++
	}

	p := 1
	for _, i := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		p *= digits[i-1]
	}

	answer := fmt.Sprintf("%d", p)
	ctx.SetAnswer(answer)
}
