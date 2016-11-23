package euler082

import (
	"fmt"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler082 *.txt

func init() {
	euler.Register("082", "Path sum: three ways", solve)
}

var (
	known = [][]int{
		[]int{131, 673, 234, 103, 18},
		[]int{201, 96, 342, 965, 150},
		[]int{630, 803, 746, 422, 111},
		[]int{537, 699, 497, 121, 956},
		[]int{805, 732, 524, 37, 331},
	}
)

func solve(ctx *euler.Context) {

	if euler.Verbose {
		fmt.Println("known: ", compute(known))
	}

	unknown := euler.ReadMatrix(string(rawfiles["p082_matrix.txt"]))
	answer := fmt.Sprintf("%d", compute(unknown))
	ctx.SetAnswer(answer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func compute(m [][]int) int {
	rows, cols := len(m), len(m[0])
	lastRow := rows - 1

	// Running minimum sum for each row
	sums := make([]int, cols)

	// Start sums using values in the first column
	for r := 0; r < rows; r++ {
		sums[r] = m[r][0]
	}

	// Scan columns to the right, starting with the second column
	for c := 1; c < cols; c++ {
		// Row 0: add RIGHT
		sums[0] += m[0][c]

		for r := 1; r < rows; r++ {
			// Row 1 to N: minimum(ABOVE, RIGHT) + CURRENT
			above := sums[r-1]
			sums[r] = min(sums[r], above) + m[r][c]
		}

		for r := lastRow - 1; r >= 0; r-- {
			// Row N-1 to 0: minimum(CURRENT, BELOW)
			below := sums[r+1]
			sums[r] = min(sums[r], below+m[r][c])
		}
	}

	// Answer is the minimum sum
	res := 0
	for _, n := range sums {
		if res == 0 {
			res = n
		} else {
			res = min(res, n)
		}
	}

	return res
}
