package euler056

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("056", "Powerful digit sum", solve)
}

func solve(ctx *euler.Context) {
	maxsum := 0
	na := &big.Int{}
	nb := &big.Int{}

	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {

			// Compute a**b
			na.SetInt64(int64(a))
			nb.SetInt64(int64(b))
			na.Exp(na, nb, nil)

			// Sum up the digits of the bigint
			sum := 0
			for _, c := range na.String() {
				sum += int(c - '0')
			}

			if sum > maxsum {
				maxsum = sum
			}

		}
	}

	answer := fmt.Sprintf("%d", maxsum)
	ctx.SetAnswer(answer)

}
