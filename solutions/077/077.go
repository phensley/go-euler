package euler077

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("077", "Prime summations", solve)
}

func solve(ctx *euler.Context) {
	primes := euler.PrimesSieveOfAtkin(100).Ones()
	n := 2
	for {
		numbers := []int{}
		for _, p := range primes {
			if int(p) > n {
				break
			}
			numbers = append(numbers, int(p))
		}

		s := euler.NewBigSummation()
		count := s.Compute(n, numbers)
		if count.Int64() > 5000 {
			if euler.Verbose {
				fmt.Printf("%d in %d ways\n", n, count)
			}
			answer := fmt.Sprintf("%d", n)
			ctx.SetAnswer(answer)
			return
		}
		n++
	}
}
