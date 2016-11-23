package euler048

import (
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("048", "Self powers", solve)
}

func solve(ctx *euler.Context) {
	// Holds way more than 10 digits, more than enough
	s := big.NewInt(1)
	p := big.Int{}
	for n := int64(2); n <= 1000; n++ {
		p.SetInt64(n)
		p.Exp(&p, &p, nil)
		s.Add(s, &p)
	}
	r := s.String()

	answer := r[len(r)-10:]
	ctx.SetAnswer(answer)
}
