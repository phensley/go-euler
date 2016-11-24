package euler067

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler067 *.txt

func init() {
	euler.Register("067", "Maximum path sum 2", solve)
}

func solve(ctx *euler.Context) {
	triangle := readTriangle()
	sum := euler.TriangleMaximumPathSum(triangle)
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func readTriangle() [][]int {
	lines := strings.Split(rawfiles["p067_triangle.txt"], "\n")

	rows := [][]int{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := []int{}
		columns := strings.Split(line, " ")
		for _, col := range columns {
			num, err := strconv.ParseInt(col, 10, 32)
			euler.FatalOnError(err, "ParseInt")
			row = append(row, int(num))
		}
		rows = append(rows, row)
	}
	return rows
}
