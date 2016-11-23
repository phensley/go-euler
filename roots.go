package euler

import (
	"math"
	"math/big"
)

// NthRootNewtonFloat64 uses Newton's method of finding the Nth root
// of a number, converging towards it with increasing iterations.
func NthRootNewtonFloat64(n, p float64, iters int) float64 {
	x := float64(float64(n) / float64(p))
	p1 := float64(p - 1)
	for i := 0; i < iters; i++ {
		a := 1.0 / float64(p)
		b := p1 * x
		c := math.Pow(x, p1)
		d := float64(n) / c
		xn := a * (b + d)
		x = xn
	}
	return x

}

// NthRootNewtonBigFloat uses Newton's method of finding the Nth root of a number
// to a given precision, converging towards it with increasing iterations.
func NthRootNewtonBigFloat(n, p float64, precision uint, iters int) *big.Float {
	x := big.NewFloat(n / p).SetPrec(precision)
	p1 := big.NewFloat(p - 1.0).SetPrec(precision)
	for i := 0; i < iters; i++ {
		a := big.NewFloat(1.0).SetPrec(precision)
		a = a.Quo(a, big.NewFloat(p).SetPrec(precision))
		b := big.NewFloat(1.0).SetPrec(precision)
		b.Mul(p1, x)
		c := big.NewFloat(1.0).SetPrec(precision)
		c.Mul(c, x)
		for j := 0; j < int(p-2); j++ {
			c.Mul(c, x)
		}
		d := big.NewFloat(n).SetPrec(precision)
		d.Quo(d, c)
		e := big.NewFloat(0.0).SetPrec(precision).Add(b, d)
		xn := big.NewFloat(1.0).SetPrec(precision).Mul(a, e)
		x = xn
	}
	return x
}
