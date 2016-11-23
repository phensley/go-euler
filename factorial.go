package euler

import "math/big"

// DigitFactorials produces the factorials for digits 0-9
func DigitFactorials() []int {
	r := make([]int, 10)
	r[0] = 1
	fac := 1
	for i := 1; i < 10; i++ {
		fac *= i
		r[i] = fac
	}
	return r
}

// SmallFactorial produces the factorial for N
func SmallFactorial(n uint64) uint64 {
	if n > 20 {
		panic("factorials n > 20 exceed capacity of uint64")
	}
	f := uint64(1)
	for n > 1 {
		f *= n
		n--
	}
	return f
}

// BigFactorial computes big.Int factorials of N with memoization
type BigFactorial struct {
	cache map[uint32]*big.Int
}

// Calculate computes big.Int factorials of N with memoization
func (f *BigFactorial) Calculate(n uint32) *big.Int {
	if f.cache == nil {
		f.cache = make(map[uint32]*big.Int)
	}

	if n == 0 {
		n = 1
	}

	r := f.cache[n]
	if r != nil {
		return big.NewInt(0).Set(r)
	}

	r = big.NewInt(int64(n))
	if n > 2 {
		t := f.Calculate(n - 1)
		r.Mul(r, t)
	}

	f.cache[n] = r
	return big.NewInt(0).Set(r)
}
