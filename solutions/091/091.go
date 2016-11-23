package euler091

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("091", "Right triangles with integer coordinates", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		fmt.Println("known grid 2x2: ", compute(2))
	}

	count := compute(50)
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

func compute(limit int) int {
	count := 0

	// Count all triangles and their corresponding reflection about the Y=X diagonal
	// This covers triangles whose 90-degree angle is not at the origin.
	for py := 0; py <= limit; py++ {
		for px := 0; px <= limit; px++ {
			if px == 0 && py == 0 {
				continue
			}

			// P is at some point not the origin. We want to turn left and right
			// from this point and go a certain distance determined by the
			// reduced X/Y slope.

			// Find the reduced slope at point P by dividing by the GCD.  Then
			// turn both left and right 90-degrees and advance Q one or more increments
			// in that direction.
			gcd := int(euler.GreatestCommonDivisor(uint64(px), uint64(py)))
			dx := py / gcd
			dy := -px / gcd

			for _, direction := range []int{-1, 1} {
				incr := direction
				for {
					// Compute Q using the slope
					qx := px + (dx * incr)
					qy := py + (dy * incr)

					// Stop scanning this direction for Q when we hit a bound
					if (qx == 0 && qy == 0) || qx < 0 || qy < 0 || qx > limit || qy > limit {
						break
					}

					incr += direction
					count++
				}
			}
		}
	}

	// Count all triangles where the 90-degree angle is at the origin
	for y := 1; y <= limit; y++ {
		for x := 1; x <= limit; x++ {
			count++
		}
	}
	return count
}
