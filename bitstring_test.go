package euler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitstringStorageAndCount(t *testing.T) {
	b := NewBitString(129)
	for i := uint64(0); i < 64; i += 2 {
		b.Set(i)
	}
	assert.Equal(t, uint64(32), b.Count())

	// Already set
	b.Set(12)
	b.Set(14)
	assert.Equal(t, uint64(32), b.Count())

	b.Set(120)
	assert.Equal(t, uint64(33), b.Count())

	// Clears 12 and 14
	b.Flip(12)
	b.Flip(14)
	assert.Equal(t, uint64(31), b.Count())

	// Sets 12
	b.Flip(12)
	b.Clear(120)
	assert.Equal(t, uint64(31), b.Count())

	assert.Equal(t, uint64(129), b.Limit())
}

func TestBitstringExpand(t *testing.T) {
	b := NewBitString(33)
	b.Set(15)
	b.Set(32)
	b.Set(35)
	b.Expand(33)
	fmt.Println(b)
	fmt.Println(b.Count())
	fmt.Println(b.Limit())
}

func TestBitstringMask(t *testing.T) {
	b := NewBitString(13)
	b.SetAll()
	fmt.Println("count ", b.Count())
}

func TestBitstringAndOR(t *testing.T) {
	n := uint64(100)
	a := NewBitString(n)
	b := NewBitString(n)
	d := NewBitString(n)
	for i := uint64(0); i < n; i++ {
		if i%2 == 0 {
			a.Set(i)
		} else {
			b.Set(i)
		}
	}

	a.Or(b, d)
	assert.Equal(t, uint64(100), d.Count())

	d.ClearAll()
	a.And(b, d)
	assert.Equal(t, uint64(0), d.Count())
}

func TestBitstringIterator(t *testing.T) {
	n := uint64(16)
	b := NewBitString(n)
	for i := uint64(0); i < n; i += 2 {
		b.Set(i)
	}

	evens := []uint64{0, 2, 4, 6, 8, 10, 12, 14}
	odds := []uint64{1, 3, 5, 7, 9, 11, 13, 15}
	all := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	none := []uint64{}

	assert.Equal(t, evens, readFromIterator(b.OneIterator()))
	assert.Equal(t, odds, readFromIterator(b.ZeroIterator()))

	b.ClearAll()
	assert.Equal(t, none, readFromIterator(b.OneIterator()))
	assert.Equal(t, all, readFromIterator(b.ZeroIterator()))

	b.SetAll()
	assert.Equal(t, all, readFromIterator(b.OneIterator()))
	assert.Equal(t, none, readFromIterator(b.ZeroIterator()))
}

func readFromIterator(iter *BitstringIterator) []uint64 {
	r := []uint64{}
	for iter.Next() {
		r = append(r, iter.Get())
	}
	return r
}
