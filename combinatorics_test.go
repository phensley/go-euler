package euler

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	msg := "%d-digit counter, %d digits"
	for _, n := range []int{3, 4, 5} {
		for _, m := range []int{3, 4, 5} {
			r := NewCounter(n, m).All()
			expected := int(math.Pow(float64(m), float64(n)))
			assert.Equal(t, expected, len(r), msg, n, m)
		}
	}
}

func TestMixedRadixCounter(t *testing.T) {
	r := NewMixedRadixCounter([]int{2, 3}).All()
	assert.Equal(t, 2*3, len(r), "mixed radix {2, 3} sequence length")

	r = NewMixedRadixCounter([]int{2, 4}).All()
	assert.Equal(t, 2*4, len(r), "mixed radix {2, 4} sequence length")

	r = NewMixedRadixCounter([]int{2, 3, 4, 5}).All()
	assert.Equal(t, 2*3*4*5, len(r), "mixed radix {2, 3, 4, 5} sequence length")
}

func TestRotations(t *testing.T) {
	msg := "sequence of rotations of %#v"
	for _, size := range []int{3, 4, 5, 6} {
		n := IntRange(size)
		r := NewRotations(n).All()
		assert.Equal(t, size, len(r), msg, n)
	}
}

func TestLexicographicPermutations(t *testing.T) {
	msg := "permutations of %#v in lexicographic order"
	for _, size := range []int{2, 3, 4, 6, 8} {
		n := IntRange(size)
		r := NewPermutations(n).All()
		expected := SmallFactorial(uint64(size))
		assert.Equal(t, expected, uint64(len(r)), msg, n)
	}
}

func TestPartialPermutations(t *testing.T) {
	set := []int{2, 3, 4}
	perms := NewPartialPermutations(set, 2)
	for perms.Next() {
		fmt.Println(perms.Get())
	}
}
