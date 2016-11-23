package euler007

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("007", "10001st prime", solve)
}

func solve(ctx *euler.Context) {
	which := 10001
	i := 1
	n := 3
	for {
		if euler.IsPrimeForisekJancina32(uint32(n)) {
			i++
			if i == which {
				if euler.Verbose {
					fmt.Printf("the %dth prime is %d\n", which, n)
				}
				answer := fmt.Sprintf("%d", n)
				ctx.SetAnswer(answer)
				return
			}
		}
		n += 2
	}
}
