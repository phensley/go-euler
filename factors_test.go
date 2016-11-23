package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, uint64(3), EuclidGCD(9, 12))
	assert.Equal(t, uint64(3), EuclidGCD(12, 9))
	assert.Equal(t, uint64(4), EuclidGCD(12, 20))
	assert.Equal(t, uint64(4), EuclidGCD(20, 12))
	assert.Equal(t, uint64(7), EuclidGCD(7, 49))
	assert.Equal(t, uint64(7), EuclidGCD(49, 7))

	assert.Equal(t, uint64(3), GreatestCommonDivisor(9, 12))
	assert.Equal(t, uint64(3), GreatestCommonDivisor(12, 9))
	assert.Equal(t, uint64(4), GreatestCommonDivisor(12, 20))
	assert.Equal(t, uint64(4), GreatestCommonDivisor(20, 12))
	assert.Equal(t, uint64(7), GreatestCommonDivisor(7, 49))
	assert.Equal(t, uint64(7), GreatestCommonDivisor(49, 7))
}
