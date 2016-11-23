package euler030

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("030", "Digit fifth powers", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		fmt.Println("Confirming 4th powers: ", calculate(4))
	}
	answer := fmt.Sprintf("%d", calculate(5))
	ctx.SetAnswer(answer)
}

func calculate(power int) int {
	// Precompute powers for digits 0-9
	powers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range powers {
		powers[i] = int(math.Pow(float64(powers[i]), float64(power)))
	}

	// Upper limit on sum of powers
	limit := int(power+1) * int(math.Pow(float64(9), float64(power)))
	sum := 0
	for n := 2; n < limit; n++ {
		digits := euler.IntToDigits(n)

		// Produce the sum of the powers of the digits
		r := 0
		for _, d := range digits {
			r += powers[d]
		}

		// If sum of powers == n, collect it
		if r == n {
			sum += r
			if euler.Verbose {
				fmt.Println(n, digits)
			}
		}
	}
	return sum
}
