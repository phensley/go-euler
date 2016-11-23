package euler071

import (
	"fmt"
	"math"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("071", "Ordered fractions", solve)
}

func solve(ctx *euler.Context) {
	limit := 1000000

	numer := int64(3)
	denom := int64(7)

	// Find the factor by which to increase the N/D
	factor := int64(math.Floor(float64(limit) / float64(denom)))

	// Decrement the scaled-up 3/7 to find the fraction just to its left
	// in a sorted list
	n := (numer * factor) - 1
	d := (denom * factor) - 2

	// Construct a Rat which will auto-reduce
	r := big.NewRat(n, d)

	if euler.Verbose {
		fmt.Printf("... %d/%d  3/7 ...\n", r.Num().Int64(), r.Denom().Int64())
	}
	answer := fmt.Sprintf("%d", n)
	ctx.SetAnswer(answer)

}
