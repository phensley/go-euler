package euler069

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("069", "Totient maximum", solve)
}

func solve(ctx *euler.Context) {
	max := uint64(0)
	maxratio := float64(0)
	for i, n := range euler.EulersTotient(1000000) {
		if i == 0 {
			continue
		}
		ratio := float64(i) / float64(n)
		if maxratio < ratio {
			maxratio = ratio
			max = uint64(i)
		}
	}
	if euler.Verbose {
		fmt.Println("maximum n/phi(n): ", maxratio)
	}
	answer := fmt.Sprintf("%d", max)
	ctx.SetAnswer(answer)
}
