package euler080

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("080", "Square root digital expansion", solve)
}

func solve(ctx *euler.Context) {
	sum := 0
	for n := 2; n <= 101; n++ {
		d := squareRoot(n, 100)
		if len(d) == 100 {
			sum += sumDigits(d)
		}
	}
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func sumDigits(d []int) int {
	r := 0
	for i := 0; i < len(d); i++ {
		r += d[i]
	}
	return r
}

// Calculates the exact digits of the square root of N using
// the shifting-digits algorithm.
func squareRoot(n, limit int) []int {
	rad := radicand(n, 2)
	p := big.NewInt(0)
	i := 0
	r := big.NewInt(0)
	d := []int{}
	f := big.NewInt(100)
	for {
		// Bring down the next 2 digits from the radicand, or 0 if no
		// digits are left.
		q := 0
		if i < len(rad) {
			q = rad[i]
		}

		// Shift the remainder 2 digits to the left and add the 2 digits
		// from the radicand.
		c := big.NewInt(1).Mul(r, f)
		c.Add(c, big.NewInt(int64(q)))

		// Find the value of X that produces the value Y closest to
		// our block of digits without going over.
		y, x := findX(2, p, c)

		// Subtract to get a new remainder
		r.Sub(c, y)

		// X becomes the next digit in the square root
		d = append(d, int(x.Int64()))

		// If our remainder is 0 we're finished.
		if r.Cmp(big.NewInt(0)) == 0 {
			break
		}

		// P holds the digits in the intermediate answer, used to find the next X
		p.Mul(p, big.NewInt(10))
		p.Add(p, x)
		i++
		if i >= limit {
			break
		}
	}
	return d
}

func findX(power int, p, c *big.Int) (*big.Int, *big.Int) {
	x := big.NewInt(0)
	pp := big.NewInt(1).Mul(big.NewInt(int64(power)), big.NewInt(10))
	pp.Mul(pp, p)
	for {
		x.Add(x, big.NewInt(1))
		b := big.NewInt(0).Add(pp, x)
		y := big.NewInt(0).Mul(x, b)
		if y.Cmp(c) == 1 {
			x.Sub(x, big.NewInt(1))
			b = big.NewInt(0).Add(pp, x)
			r := big.NewInt(1)
			r.Mul(x, b)
			return r, x
		}
	}
}

func radicand(n int, power int) []int {
	d := euler.IntToDigits(n)
	length := len(d)
	x := length / power
	r := length % power
	if x == 0 {
		return []int{euler.DigitsToInt(d)}
	}
	res := []int{}
	if r != 0 {
		res = append(res, euler.DigitsToInt(d[:r]))
		d = d[r:]
	}
	for x > 0 {
		res = append(res, euler.DigitsToInt(d[:power]))
		d = d[power:]
		x--
	}
	return res
}
