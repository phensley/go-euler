package euler015

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("015", "Lattice paths", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		bruteForceConfirmation()
	}

	f := &euler.BigFactorial{}
	for n := uint32(2); n < 21; n++ {
		r := interleavings(f, n)
		if n == 20 {
			answer := fmt.Sprintf("%d", r)
			ctx.SetAnswer(answer)
		}
		if euler.Verbose {
			fmt.Printf("%d x %d grid: %s paths\n", n, n, r)
		}
	}
}

// Pattern is central coefficient in even rows of Pascal's triangle:
//
// 0 x 0 grid
//    1
//
// 1 x 1 grid  (2 choose 1) = 2
//    2  1
//    1  2
//
// 2 x 2 grid  (4 choose 2) = 6
//    6  3  1
//    3  4  3
//    1  3  6
//
// 3 x 3 grid  (6 choose 3) = 20
//   20 10  4  1
//   10 12  9  4
//    4  9 12 10
//    1  4 10 20
//
// 4 x 4 grid  (8 choose 4) = 70
//   70 35 15  5  1
//   35 40 30 16  5
//   15 30 36 30 15
//    5 16 30 40 35
//    1  5 15 35 70
//
func interleavings(f *euler.BigFactorial, n uint32) *big.Int {
	fn := f.Calculate(n)
	fn.Mul(fn, fn)
	f2n := f.Calculate(2 * n)
	return f2n.Div(f2n, fn)
}

func bruteForceConfirmation() {
	// Count of integers up to 2N who have exactly N bits set
	// We only count up to 10 since this takes quite a while
	limit := uint64(10)
	fmt.Printf("brute force up to %d\n", limit-1)
	for n := uint64(2); n < 10; n++ {
		board := 2 * n
		limit := uint64((1 << board))
		count := uint64(0)
		for i := uint64(0); i < limit; i++ {
			if euler.CountBitsSet64(i) == n {
				count++
			}
		}
		fmt.Printf("%d x %d == %d paths\n", n, n, count)
	}
	fmt.Println()
}
