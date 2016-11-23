package euler078

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("078", "Coin partitions", solve)
}

func solve(ctx *euler.Context) {
	k := pentagonal(250)
	p := []int{1}
	signs := []int{1, 1, -1, -1}
	n := 0
	million := int64(1000000)
	for p[n] > 0 {
		n++
		// Generating function
		px := int64(0)
		i := 0
		for k[i] <= n {
			px += int64(p[n-k[i]] * signs[i%4])
			i++
		}
		p = append(p, int(euler.EuclideanMod(px, million)))
	}
	answer := fmt.Sprintf("%d", n)
	ctx.SetAnswer(answer)
}

// Exponents for the Euler function which calculates the number of
// partitions for N
func pentagonal(max int) []int {
	r := []int{}
	for k := 1; k < max; k++ {
		n := k * (3*k - 1) / 2
		r = append(r, n)
		r = append(r, n+k)
	}
	return r
}
