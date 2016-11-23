package euler026

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("026", "Reciprocal cycles", solve)
}

func solve(ctx *euler.Context) {
	// d, where 1/d produces the longest recurrence for d < 1000
	longest := uint(0)

	// Length of the longest recurrence for 1/d
	length := 0

	// Consider only d < 1000
	limit := uint(1000)

	n := uint(1)
	for d := uint(2); d < limit; d++ {
		integ, fract, recur := division(n, d)
		if len(recur) > length {
			longest = d
			length = len(recur)
		}

		if euler.Verbose {
			fmt.Printf("%d / %d == %s\n", n, d, convert(integ, fract, recur))
			if len(recur) > 0 {
				fmt.Println("recurrence length=", len(recur))
			}
			fmt.Println()
		}
	}
	answer := fmt.Sprintf("%d", longest)
	ctx.SetAnswer(answer)
}

func digitsToString(digits []uint) string {
	r := []byte{}
	for _, d := range digits {
		r = append(r, byte('0')+byte(d))
	}
	return string(r)
}

func convert(integ, fract, recur []uint) string {
	r := fmt.Sprintf("%s.%s", digitsToString(integ), digitsToString(fract))
	if len(recur) == 0 {
		return r
	}
	return fmt.Sprintf("%s(%s)", r, digitsToString(recur))
}

// Generalized integer division. This will return 3 slices
// containing digits corresponding to:
//
//       <integ> . <fract> ( <recur> )
//
func division(x, y uint) ([]uint, []uint, []uint) {
	// Digits before decimal point
	integ := []uint{0}

	// Digits after decimal point including any recurrence
	fract := []uint{}

	// Track state of caculation to detect recurrence
	state := make(map[uint]*struct{})
	index := make(map[uint]int)

	// First, see if we have digits before the decimal point.
	if x/y != 0 {
		integ = euler.UintToDigits(x / y)
		x = x % y
	}

	// Compute the fractional part. We will accumulate digits
	// in a slice and keep track of the leading digits.

	// Increasing here sets up to calculate the first digit, so
	// does not add a leading zero.
	if x > 0 && x/y == 0 {
		x *= 10
	}

	// Index of the current digit in the calculation
	i := 0

	// If further increases are needed, add more leading zeros
	for x > 0 && x/y == 0 {
		x *= 10
		fract = append(fract, 0)
		i++
	}

	for {
		// Increase x as needed and add zero digits
		for x > 0 && x/y == 0 {
			x *= 10
			fract = append(fract, 0)
		}

		// Check if we've returned to a previous state, where the
		// current divisor has been seen before. Return the point
		// where the recurrence began.  If we don't break out here
		// this would recur forever.
		if state[x] != nil {
			// Calculate the point the recurrence starts and
			// return slices for each part.
			r := index[x]
			return integ, fract[:r], fract[r:]
		}

		// Record the current state
		state[x] = &struct{}{}
		index[x] = i

		// Div mod
		n := x / y
		m := x % y

		// Save digit
		fract = append(fract, n)

		// If no remainder, we're done
		if m == 0 {
			break
		}

		// Continue the calculation with the remainder * 10
		x = m * 10

		i++
	}

	// No recurrence was found, return digits.
	return integ, fract, []uint{}
}
