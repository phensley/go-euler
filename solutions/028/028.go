package euler028

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("028", "Number spiral diagonals", solve)
}

func solve(ctx *euler.Context) {
	answer := fmt.Sprintf("%d", spiral(1001))
	ctx.SetAnswer(answer)
}

func spiral(limit int) int {
	// Current incremental value of the spiral as we build it
	n := 1

	// Length of the current side
	side := 3

	// Sum of the corners starting with the center
	sum := 1

	for {
		// Add the values at the 4 corners of the spiral
		for i := 0; i < 4; i++ {
			n += side - 1
			sum += n
		}
		// Increase side length until it exceeds our maximum spiral side
		side += 2
		if side > limit {
			break
		}
	}
	return sum
}
