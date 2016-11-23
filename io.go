package euler

import (
	"strconv"
	"strings"
)

// ReadMatrix reads a matrix from the ASCII string
func ReadMatrix(raw string) [][]int {
	matrix := [][]int{}
	for _, line := range strings.Split(string(raw), "\n") {
		if len(line) == 0 {
			continue
		}
		row := []int{}
		for _, cell := range strings.Split(line, ",") {
			if len(cell) == 0 {
				continue
			}
			num, err := strconv.ParseInt(cell, 10, 64)
			FatalOnError(err, "strconv.ParseInt")
			row = append(row, int(num))
		}
		matrix = append(matrix, row)
	}
	return matrix
}
