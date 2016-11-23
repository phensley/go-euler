package euler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	for _, n := range []int{2, 3, 5, 7, 11, 13} {
		assert.True(t, IsPrimeTrialDivision(n))
	}
	for _, n := range []int{4, 6, 9, 12} {
		assert.False(t, IsPrimeTrialDivision(n), fmt.Sprintf("%d is not prime", n))
	}
}
