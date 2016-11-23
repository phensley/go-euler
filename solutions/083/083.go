package euler083

import (
	"fmt"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler083 *.txt

func init() {
	euler.Register("083", "Path sum: four ways", solve)
}

var (
	// Known case sum: 2297
	known = [][]int{
		[]int{131, 673, 234, 103, 18},
		[]int{201, 96, 342, 965, 150},
		[]int{630, 803, 746, 422, 111},
		[]int{537, 699, 497, 121, 956},
		[]int{805, 732, 524, 37, 331},
	}

	// Edge case sum: 33 + 33 == 66
	edgeCase = [][]int{
		[]int{33, 100, 100, 100, 100, 100},
		[]int{0, 100, 100, 100, 100, 100},
		[]int{0, 0, 0, 100, 100, 100},
		[]int{100, 0, 100, 0, 0, 0},
		[]int{100, 0, 100, 0, 100, 0},
		[]int{100, 0, 0, 0, 100, 33},
	}
)

func solve(ctx *euler.Context) {
	if euler.Verbose {
		fmt.Println("    known: ", compute(known))
		fmt.Println("edge case: ", compute(edgeCase))
	}

	unknown := euler.ReadMatrix(string(rawfiles["p083_matrix.txt"]))
	answer := fmt.Sprintf("%d", compute(unknown))
	ctx.SetAnswer(answer)
}

func compute(m [][]int) int {
	rows, cols := len(m), len(m[0])

	startID := int64(0)
	endID := int64(((rows - 1) * cols) + (cols - 1))

	start := euler.NewIntNode(startID, m[0][0])
	end := euler.NewIntNode(endID, m[rows-1][cols-1])

	// Links in our graph can only be up, down, left and right
	directions := [][]int{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, -1},
		[]int{0, 1},
	}

	// Convert the matrix into a graph and use Dijkstra's shortest path
	// to solve it.
	g := euler.IntMatrixToGraph(m, directions)
	path := g.DijkstraShortestPath(start, end)

	// Sum the weights of the nodes in the shortest path
	sum := 0
	for _, n := range path {
		sum += n.(*euler.IntNode).Value()
	}
	return sum
}
