package euler016

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("016", "Power digit sum", solve)
}

func solve(ctx *euler.Context) {
	b := big.Int{}
	b.SetBit(&b, 1000, 1)
	sum := sumDigits(&b)
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func sumDigits(n *big.Int) uint64 {
	sum := uint64(0)
	s := n.String()
	ch := '0'
	for i := 0; i < len(s); i++ {
		sum += uint64(s[i] - byte(ch))
	}
	return sum
}
