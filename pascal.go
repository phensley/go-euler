package euler

import "math/big"

// PascalTriangleRow computes the elements of the Nth row
// of Pascal's Triangle
func PascalTriangleRow(n int) []*big.Int {
	r := make([]*big.Int, n+1)
	r[0] = big.NewInt(int64(1))
	j := n
	for k := 1; k <= j; k++ {
		x := big.NewInt(int64(n))
		x.Mul(x, r[k-1])
		x.Div(x, big.NewInt(int64(k)))
		r[k] = x
		n--
	}
	return r
}
