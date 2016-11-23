package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitFactorials(t *testing.T) {
	fac := DigitFactorials()
	assert.Equal(t, 1, fac[0], "factorial(0)")
	assert.Equal(t, 1, fac[1], "factorial(1)")
	assert.Equal(t, 2, fac[2], "factorial(2)")
	assert.Equal(t, 6, fac[3], "factorial(3)")
	assert.Equal(t, 24, fac[4], "factorial(4)")
	assert.Equal(t, 120, fac[5], "factorial(5)")
	assert.Equal(t, 720, fac[6], "factorial(6)")
	assert.Equal(t, 5040, fac[7], "factorial(7)")
	assert.Equal(t, 40320, fac[8], "factorial(8)")
	assert.Equal(t, 362880, fac[9], "factorial(9)")
}

func TestSmallFactorial(t *testing.T) {
	assert.Equal(t, uint64(1), SmallFactorial(0), "factorial(0)")
	assert.Equal(t, uint64(1), SmallFactorial(1), "factorial(1)")
	assert.Equal(t, uint64(6), SmallFactorial(3), "factorial(3)")
	assert.Equal(t, uint64(24), SmallFactorial(4), "factorial(4)")
	assert.Equal(t, uint64(120), SmallFactorial(5), "factorial(5)")
	assert.Equal(t, uint64(720), SmallFactorial(6), "factorial(6)")
	assert.Equal(t, uint64(2432902008176640000), SmallFactorial(20), "factorial(20)")
	assert.Panics(t, func() { SmallFactorial(21) }, "factorial(21)")
}

func TestBigFactorial(t *testing.T) {
	fac := BigFactorial{}
	assert.Equal(t, "1", fac.Calculate(0).String(), "factorial(0)")
	assert.Equal(t, "1", fac.Calculate(1).String(), "factorial(1)")
	assert.Equal(t, "6", fac.Calculate(3).String(), "factorial(6)")
	assert.Equal(t, "2432902008176640000", fac.Calculate(20).String(), "factorial(20)")
	assert.Equal(t, "51090942171709440000", fac.Calculate(21).String(), "factorial(21)")
	assert.Equal(t,
		"30414093201713378043612608166064768844377641568960512000000000000",
		fac.Calculate(50).String(),
		"factorial(50)")
}
