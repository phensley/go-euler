package euler

import "math"

// AbsInt returns the absolute value of integer a
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// AbsInt64 returns absolute value of a
func AbsInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

// AbsFloat64 returns absolute value of a
func AbsFloat64(a float64) float64 {
	if a < 0.0 {
		return -a
	}
	return a
}

// RoundFloat64 ..
func RoundFloat64(n float64) int64 {
	i := int64(n)
	d := AbsFloat64(n - float64(i))
	incr := int64(1)
	if n < 0.0 {
		incr = -1
	}
	if d < 0.5 {
		return i
	}
	return i + incr
}

// EuclideanMod implements Euclidean modulus x mod y
func EuclideanMod(x, y int64) int64 {
	m := x % y
	if m < 0 {
		return m + y
	}
	return m
}

// NumDigitsBase10 returns the number of digits in N in base-10
func NumDigitsBase10(n uint64) float64 {
	if n == 0 {
		return 1
	}
	return math.Floor(math.Log10(float64(n))) + 1
}

// Mul64 multiplies two 64-bit numbers together and returns
// a 128-bit result as two 64-bit numbers representing the high and low bits
// of the result.  Derived from examples in Hacker's Delight.
// http://www.hackersdelight.org/hdcodetxt/mont64.c.txt
func Mul64(u, v uint64) (hi, lo uint64) {
	u1 := u >> 32
	u0 := u & 0xFFFFFFFF
	v1 := v >> 32
	v0 := v & 0xFFFFFFFF

	t := u0 * v0
	w0 := t & 0xFFFFFFFF
	k := t >> 32

	t = u1*v0 + k
	w1 := t & 0xFFFFFFFF
	w2 := t >> 32

	t = u0*v1 + w1
	k = t >> 32

	lo = (t << 32) + w0
	hi = u1*v1 + w2 + k
	return
}
