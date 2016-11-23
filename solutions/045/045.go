package euler045

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("045", "Triangular, pentagonal, and hexagonal", solve)
}

func solve(ctx *euler.Context) {
	tn := 285
	pn := uint64(1)
	hn := uint64(1)
	p := uint64(0)
	h := uint64(0)
	for {
		tn++
		t := uint64((tn * (tn + 1)) / 2)

		for p < t {
			p = uint64(pn*((3*pn)-1)) / 2
			if p > t {
				break
			}
			pn++
		}

		for h < t {
			h = uint64(hn * ((2 * hn) - 1))
			if h > t {
				break
			}
			hn++
		}

		if t == p && t == h {
			answer := fmt.Sprintf("%d", t)
			ctx.SetAnswer(answer)
			return
		}
	}
}
