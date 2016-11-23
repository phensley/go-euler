package euler063

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("063", "Powerful digit counts", solve)
}

func solve(ctx *euler.Context) {
	count := 0
	for p := int64(1); p <= 30; p++ {
		for n := int64(1); n < 500; n++ {
			x := power(n, p)
			if len(x) == int(p) {
				if euler.Verbose {
					fmt.Printf("%d ** %d == %s\n", n, p, x)
				}
				count++
			}
		}
	}
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

func power(n, m int64) string {
	x := big.NewInt(n)
	y := big.NewInt(m)
	x.Exp(x, y, nil)
	return x.String()
}
