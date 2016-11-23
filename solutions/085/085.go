package euler085

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("085", "Counting rectangles", solve)
}

func solve(ctx *euler.Context) {
	target := 2000000
	limit := 100
	nearest := -1
	area := 0

	for w := 2; w < limit; w++ {
		for h := 2; h < limit; h++ {
			// Calculate number of rectangles that can fill W x H
			count := rectangles(w, h)

			// Find the smallest difference between our count and the target
			diff := euler.AbsInt(target - count)
			if nearest == -1 || diff < nearest {
				nearest = diff
				area = w * h
			}
		}
	}

	answer := fmt.Sprintf("%d", area)
	ctx.SetAnswer(answer)
}

// Count number of rectangles that can fill an area of W x H
func rectangles(w, h int) int {
	count := 0
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			a := w - j + 1
			b := h - i + 1
			count += a * b
		}
	}
	return count
}
