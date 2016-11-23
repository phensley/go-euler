package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalTriangleRow(t *testing.T) {
	assert.Equal(t, "1", renderPascalTriangleRow(0))
	assert.Equal(t, "1, 1", renderPascalTriangleRow(1))
	assert.Equal(t, "1, 2, 1", renderPascalTriangleRow(2))
	assert.Equal(t, "1, 3, 3, 1", renderPascalTriangleRow(3))
	assert.Equal(t, "1, 4, 6, 4, 1", renderPascalTriangleRow(4))
	assert.Equal(t, "1, 5, 10, 10, 5, 1", renderPascalTriangleRow(5))
	assert.Equal(t, "1, 6, 15, 20, 15, 6, 1", renderPascalTriangleRow(6))
	assert.Equal(t, "1, 7, 21, 35, 35, 21, 7, 1", renderPascalTriangleRow(7))
	assert.Equal(t, "1, 8, 28, 56, 70, 56, 28, 8, 1", renderPascalTriangleRow(8))

	assert.Equal(t,
		"1, 16, 120, 560, 1820, 4368, 8008, 11440, 12870, 11440, 8008, 4368, 1820, 560, 120, 16, 1",
		renderPascalTriangleRow(16))

	assert.Equal(t,
		"1, 32, 496, 4960, 35960, 201376, 906192, 3365856, 10518300, 28048800, 64512240, "+
			"129024480, 225792840, 347373600, 471435600, 565722720, 601080390, 565722720, 471435600, "+
			"347373600, 225792840, 129024480, 64512240, 28048800, 10518300, 3365856, 906192, 201376, "+
			"35960, 4960, 496, 32, 1",
		renderPascalTriangleRow(32))
}

func renderPascalTriangleRow(n int) string {
	a := PascalTriangleRow(n)
	s := ""
	for i := 0; i < len(a); i++ {
		if i > 0 {
			s += ", "
		}
		s += a[i].String()
	}
	return s
}
