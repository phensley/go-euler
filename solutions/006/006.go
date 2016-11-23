package euler006

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("006", "Sum square difference", solve)
}

func solve(ctx *euler.Context) {
	n := 100
	sumOf := sumOfSquares(n)
	squareOf := squareOfSum(n)
	answer := fmt.Sprintf("%d", squareOf-sumOf)
	ctx.SetAnswer(answer)
}

func squareOfSum(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i
	}
	return sum * sum
}

func sumOfSquares(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i * i
	}
	return sum
}
