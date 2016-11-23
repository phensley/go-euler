package euler100

import (
	"fmt"
	"math"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("100", "Arranged probability", solve)
}

var (
	oneHalf = big.NewRat(1, 2)
)

func solve(ctx *euler.Context) {
	if euler.Verbose {
		// Brute force search for small inputs to determine possible series
		for n := int64(10); n < 3000; n++ {
			if r, found := search(n); found {
				fmt.Print(r)
			}
		}
		fmt.Println()
	}

	sqrt2 := math.Sqrt(2.0)
	limit := math.Pow(10, 12)
	i := int64(0)
	for {
		n := computeA011900(i)
		sqrt := sqrt2 * math.Sqrt(float64(n)) * math.Sqrt(float64(n-1))
		if sqrt > limit {
			answer := fmt.Sprintf("%d", n)
			ctx.SetAnswer(answer)
			return
		}
		i++
	}
}

// Confirmed series is OEIS A011900
// Formula:  a(n) = 6*a(n-1) - a(n-2) - 2
//   where:  a(0) = 1, a(1) = 3
func computeA011900(n int64) int64 {
	switch n {
	case 0:
		return 1
	case 1:
		return 3
	default:
		return 6*computeA011900(n-1) - computeA011900(n-2) - 2
	}
}

// Brute-force search small range to determine series
func search(n int64) (int64, bool) {
	t := n * (n - 1)
	b := float64(t) * 0.5
	e := int64(math.Ceil(math.Sqrt(b)))

	// Search down from the square root until we're < 0.5
	for {
		p := big.NewRat(e, n)
		q := big.NewRat(e-1, n-1)
		q.Mul(q, p)
		v := q.Cmp(oneHalf)
		if v == 0 {
			return e, true
		}
		if v == -1 {
			break
		}
		e--
	}
	return -1, false
}
