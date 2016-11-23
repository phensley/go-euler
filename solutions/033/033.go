package euler033

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("033", "Digit cancelling fractions", solve)
}

func solve(ctx *euler.Context) {
	n := numbers()
	numerator := uint(1)
	denominator := uint(1)
	for _, i := range n {
		for _, j := range n {

			// Only examine results < 1.0
			if j <= i {
				continue
			}

			di := euler.UintToDigits(i)
			dj := euler.UintToDigits(j)

			// Non-trivial cases only
			if di[0] != dj[1] && di[1] != dj[0] {
				continue
			}

			r := float64(i) / float64(j)
			if float64(di[0])/float64(dj[1]) == r {
				if euler.Verbose {
					fmt.Println(i, "/", j, " == ", di[0], "/", dj[1])
				}
				numerator *= di[0]
				denominator *= dj[1]
			}
			if float64(di[1])/float64(dj[0]) == r {
				if euler.Verbose {
					fmt.Println(i, "/", j, " == ", di[1], "/", dj[0])
				}
				numerator *= di[1]
				denominator *= dj[0]
			}
		}
	}

	// Find greatest common factor
	fac := intersect(factors(numerator), factors(denominator))
	f := uint(1)
	if len(fac) >= 0 {
		f = fac[len(fac)-1]
	}

	answer := fmt.Sprintf("%d", denominator/f)
	ctx.SetAnswer(answer)
}

func factors(n uint) []uint {
	r := []uint{1}
	for i := uint(2); i <= n/2; i++ {
		if n%i == 0 {
			r = append(r, i)
		}
	}
	r = append(r, n)
	return r
}

func intersect(a, b []uint) []uint {
	r := []uint{}
	for _, i := range a {
		for _, j := range b {
			if i == j {
				r = append(r, i)
			}
		}
	}
	return r
}

func numbers() []uint {
	r := []uint{}
	for i := uint(1); i < 10; i++ {
		for j := uint(1); j < 10; j++ {
			// Equivalent numbers can't share an alternate digit
			if i == j {
				continue
			}
			r = append(r, euler.DigitsToUint([]uint{i, j}))
		}
	}
	return r
}
