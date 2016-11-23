package euler081

import (
	"fmt"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler081 *.txt

func init() {
	euler.Register("081", "Path sum: two ways", solve)
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
		sum := compute(known)
		fmt.Println("known: ", sum)
	}

	unknown := euler.ReadMatrix(string(rawfiles["p081_matrix.txt"]))
	sum := compute(unknown)
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func compute(m [][]int) int {
	rows := len(m)
	cols := len(m[0])

	// Scan the matrix in a single pass, left-to-right and up-to-down,
	// and update each cell with the minimum sum using it's up/left
	// neighbors.
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if c > 0 && r > 0 {
				// Add to the current cell the smaller of up or left cell
				up := m[r-1][c]
				lf := m[r][c-1]
				if lf < up {
					m[r][c] += lf
				} else {
					m[r][c] += up
				}

			} else if c > 0 {
				// Add to the current cell the value in the left cell
				m[r][c] += m[r][c-1]

			} else if r > 0 {
				// Add to the current cell the value of the up cell
				m[r][c] += m[r-1][c]
			}
		}
	}
	return m[rows-1][cols-1]
}
