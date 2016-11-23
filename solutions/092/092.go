package euler092

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("092", "Square digit chains", solve)
}

func solve(ctx *euler.Context) {
	limit := 10000000
	count := 0
	for n := 1; n < limit; n++ {
		r := chain(n)
		if r == 89 {
			count++
		}
	}

	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

func chain(n int) int {
	for {
		s := 0
		for n >= 10 {
			d := n % 10
			n /= 10
			s += (d * d)
		}
		if n > 0 {
			s += (n * n)
		}
		switch s {
		case 1, 89:
			return s
		default:
			n = s
		}
	}
}
