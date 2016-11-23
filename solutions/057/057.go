package euler057

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("057", "Square root convergents", solve)
}

func solve(ctx *euler.Context) {
	c := conv{make(map[int]*big.Rat)}
	c.calculate(1000)

	count := 0
	for i := 0; i < 1000; i++ {
		r := c.cache[i]
		if len(r.Num().String()) > len(r.Denom().String()) {
			count++
		}
	}

	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

type conv struct {
	cache map[int]*big.Rat
}

// Calculate all continued fraction expansions for sqrt(2) up to the limit
func (c conv) calculate(limit int) {
	n := big.NewRat(1, 2)
	one := big.NewRat(1, 1)
	two := big.NewRat(2, 1)
	iters := 0

	for iters <= limit {

		// Cache current result
		r := &big.Rat{}
		r.Set(n)
		r.Add(one, r)
		c.cache[iters] = r

		s := n.Add(two, n)
		n = n.Quo(one, s)

		iters++
	}

	// Cache last result
	n.Add(one, n)
	c.cache[iters] = n
}
