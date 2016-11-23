package euler066

import (
	"fmt"
	"math"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("066", "Diophantine equation", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		max := big.NewInt(0)
		maxD := int64(0)
		for d := int64(2); d <= 7; d++ {
			x := chakravala(d)
			if lessthan(max, x) {
				max = x
				maxD = d
			}
		}
		fmt.Printf("2 <= D <= 7  maximum X=%s where D=%d\n", max, maxD)
	}

	max := big.NewInt(0)
	maxD := int64(0)
	for d := int64(2); d <= 1000; d++ {
		x := chakravala(d)
		if lessthan(max, x) {
			max = x
			maxD = d
		}
	}

	if euler.Verbose {
		fmt.Printf("2 <= D <= 7  maximum X=%s where D=%d\n", max, maxD)
	}
	answer := fmt.Sprintf("%d", maxD)
	ctx.SetAnswer(answer)
}

// https://en.wikipedia.org/wiki/Chakravala_method
func chakravala(n int64) *big.Int {
	bign := big.NewInt(n)
	r := euler.RoundFloat64(math.Sqrt(float64(n)))
	m := big.NewInt(int64(r))
	m0 := big.NewInt(int64(r))
	x := m
	y := big.NewInt(1)
	k := big.NewInt(0).Mul(m, m)
	k.Sub(k, big.NewInt(n))

	if k.Int64() == 0 {
		return big.NewInt(0)
	}

	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	four := big.NewInt(4)
	minusOne := big.NewInt(-1)
	for !equal(k, one) {
		evenX := equal(and(x, one), zero)
		evenY := equal(and(y, one), zero)
		if equal(k, minusOne) || equal(abs(k), two) || (equal(abs(k), four) && (evenX || evenY)) {
			q := add(mult(x, x), mult(bign, mult(y, y)))
			p := abs(k)
			return div(q, p)
		}

		d := mod(add(m, m0), abs(k))
		lo := sub(m0, d)
		hi := add(lo, abs(k))
		m = lo
		if lessthan(sub(mult(hi, hi), bign), sub(mult(lo, lo), bign)) {
			m = hi
		}
		x1 := div(add(mult(m, x), mult(bign, y)), abs(k))
		y1 := div(add(x, mult(y, m)), abs(k))
		k1 := div(sub(mult(m, m), bign), k)
		x, y, k = x1, y1, k1
	}
	return x
}

func add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func mult(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return big.NewInt(0).Quo(a, b)
}

func mod(a, b *big.Int) *big.Int {
	r := big.NewInt(0)
	big.NewInt(0).QuoRem(a, b, r)
	return r
}

func and(a, b *big.Int) *big.Int {
	return big.NewInt(0).And(a, b)
}

func abs(n *big.Int) *big.Int {
	return big.NewInt(0).Abs(n)
}

func equal(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func lessthan(a, b *big.Int) bool {
	return a.Cmp(b) < 0
}
