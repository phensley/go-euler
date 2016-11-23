package euler003

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("003", "Largest prime factor", solve)
}

func solve(ctx *euler.Context) {
	n := 600851475143

	// Collect all prime factors
	var factors []int
	largest := 0

	v := n
	for i := 3; v > 1; i += 2 {
		if v%i == 0 && euler.IsPrimeForisekJancina32(uint32(i)) {
			factors = append(factors, i)
			largest = i
			// divide to reduce remaining checks
			v /= i
		}
	}

	if euler.Verbose {
		fmt.Println("prime factors of", n, "are", factors)
		fmt.Println("largest prime factor of", n, "is", largest)
	}
	answer := fmt.Sprintf("%d", largest)
	ctx.SetAnswer(answer)
}
