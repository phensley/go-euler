package euler010

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("010", "Summation of primes", solve)
}

func solve(ctx *euler.Context) {
	iter := euler.PrimesSieveOfAtkin(2000000).OneIterator()
	sum := uint64(0)
	for iter.Next() {
		sum += iter.Get()
	}
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}
