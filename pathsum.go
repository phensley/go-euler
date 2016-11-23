package euler

// TriangleMaximumPathSum calculates the maximum sum for a path
// from the top to bottom of the given triangle.
// We actually compute the sum bottom-up.
func TriangleMaximumPathSum(input [][]int) int {
	// Create a mirror of the triangle to hold our incremental sum
	rows := make([][]int, len(input))
	for i, row := range input {
		rows[i] = make([]int, len(row))
		copy(rows[i], row)
	}

	// Sum the N-1 row against the N row and work up until we
	// sum the 0th row against the first.
	start := len(rows) - 2
	for i := start; i >= 0; i-- {
		rlen := len(rows[i])

		// Move across the columns and store in the [j, N-1] cell
		// the larger sum of the [j, N] and [j, N+1] cells.
		for j := 0; j < rlen; j++ {
			if rows[i+1][j] < rows[i+1][j+1] {
				rows[i][j] += rows[i+1][j+1]
			} else {
				rows[i][j] += rows[i+1][j]
			}
		}
	}

	// Top of the pyramid now contains the maximum sum
	return rows[0][0]
}
