package euler025

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("025", "1000-digit Fibonacci number", solve)
}

func solve(ctx *euler.Context) {
	ch := euler.FibonacciBigSequence(1, 1)
	i := 1
	for n := range ch {
		s := n.String()
		if len(s) == 1000 {
			if euler.Verbose {
				fmt.Println(s)
			}
			answer := fmt.Sprintf("%d", i)
			ctx.SetAnswer(answer)
			return
		}
		i++
	}
}
