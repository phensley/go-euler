package euler094

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("094", "Almost equilateral triangles", solve)
}

func solve(ctx *euler.Context) {
	limit := uint64(1000000000)

	// Since the area of a primitive Pythagorean triangle is an integer,
	// a semi-equilateral triangle formed from two identical Pythagorean
	// triangles will have area that is an integer.

	// Start with the initial primitive triangle: (3, 4, 5). Placing 2 of
	// these together back-to-back creates a semi-equilateral (5, 5, 6) where
	// the hypotenuse forms the two equal sides a == 5 and the base b == 6 == 3 + 3

	t := euler.NewTriple(3, 4, 5)
	sum := uint64(2 * (t.A + t.C))
	for {
		// Use the Berggren matrix B to generate a sequence of primitive
		// triangles.
		t = euler.BerggrenMatrixB(t.A, t.B, t.C)
		perim := 2 * (t.A + t.C)
		if perim >= limit {
			break
		}
		sum += perim
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}
