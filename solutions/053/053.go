package euler053

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("053", "Combinatoric selections", solve)
}

func solve(ctx *euler.Context) {
	limit := big.NewInt(int64(1000000))
	count := 0

	p := &problem{&euler.BigFactorial{}}

	for n := uint32(1); n <= 100; n++ {
		for r := uint32(1); r <= n; r++ {
			r := p.selection(n, r)
			if r.Cmp(limit) >= 0 {
				count++
			}
		}
	}
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

type problem struct {
	fac *euler.BigFactorial
}

func (p *problem) selection(n, r uint32) *big.Int {
	a := big.NewInt(int64(1))
	a.Set(p.fac.Calculate(n))

	b := &big.Int{}
	b.Set(p.fac.Calculate(r))

	b.Mul(b, p.fac.Calculate(n-r))
	a.Div(a, b)
	return a
}
