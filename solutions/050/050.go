package euler050

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("050", "Consecutive prime sum", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		fmt.Print("prime < 100: ")
		fmt.Println(construct(100).compute())

		fmt.Print("prime < 1000: ")
		fmt.Println(construct(1000).compute())
	}

	sum, primes := construct(1000000).compute()
	if euler.Verbose {
		fmt.Printf("prime < 1000000: %v\n\n", primes)
	}
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

type state struct {
	primes   []uint64
	primeset *euler.Bitstring
}

func construct(maxPrime uint64) *state {
	primeset := euler.PrimesSieveOfAtkin(maxPrime)
	primes := primeset.Ones()
	return &state{primes, primeset}
}

func (s *state) dump(index, length int) (uint64, []uint64) {
	sum := uint64(0)
	r := []uint64{}
	for i := 0; i < length; i++ {
		p := s.primes[i+index]
		sum += p
		r = append(r, p)
	}
	return sum, r
}

// Finds the prime whose sum is the longest consecutive sequence of primes.
// Will be biased towards smaller primes since they produce the longest sequence.
// So we start with the smallest (2), find the longest sequence, then shift 1-position
// rightward looking for ever longer sequences. We also keep a running sum as we
// shift the window right, to avoid re-summing primes we've already covered.
func (s *state) compute() (uint64, []uint64) {
	// Starting index of the prime whose sum is the longest consecutive sequence of primes
	index := 0

	// Width of the sequence of consecutive primes for the largest sum
	width := 0

	// current index as we scan the primes from left to right
	curr := 0

	// maximum sum we've found that == a prime
	sum := uint64(0)
	for {
		// The temporary sum is used to find the next candidate.
		tempsum := sum
		for i := curr + width; i < len(s.primes); i++ {
			tempsum += s.primes[i]

			// Stop when the sum exceeds the maximum prime. Can't possibly be
			// the answer
			if tempsum > s.primeset.Limit() {
				break
			}

			// Measure length of the window.
			w := i - curr + 1

			// If the current sum is a prime and the window is wider than our
			// previous candidate, update the answer.
			if s.primeset.IsSet(tempsum) && w > width {
				if euler.Verbose {
					s.dump(index, w)
				}
				index = curr
				width = w
				sum = tempsum
			}
		}

		// If the current index + window is wider than the list of primes, we're done.
		if curr+width >= len(s.primes) {
			break
		}

		// Subtract the first prime in the window and add the next prime after
		// the window, then shift the window 1 position left.
		sum -= s.primes[curr]
		sum += s.primes[curr+width]
		curr++
	}

	return s.dump(index, width)
}
