package euler

import (
	"fmt"
	"math"
)

const allOnes = uint64(0xFFFFFFFFFFFFFFFF)

// Bitstring represents a dense string of at least N bits
type Bitstring struct {
	limit   uint64
	storage uint64
	count   uint64
	bits    []uint64
	mask    uint64
}

// NewBitString constructs a bitstring holding at least N bits
func NewBitString(numbits uint64) *Bitstring {
	storage := uint64(math.Ceil(float64(numbits) / float64(64)))
	if numbits%64 == 0 {
		storage++
	}
	// find index of max-th bit and compute a mask
	mask := ^(allOnes << (numbits % 64))
	limit := (storage * 64) - (64 - CountBitsSet64(mask))
	return &Bitstring{limit, storage, 0, make([]uint64, storage), mask}
}

// Expand the bitstring to hold more bits
func (b *Bitstring) Expand(incr uint64) {
	if incr == 0 {
		return
	}
	b.storage = uint64(math.Ceil(float64(b.limit+incr) / float64(64)))
	if (b.limit+incr)%64 == 0 {
		b.storage++
	}
	b.mask = ^(allOnes << (b.storage % 64))
	b.limit = (b.storage * 64) - (64 - CountBitsSet64(b.mask))
	newbits := make([]uint64, b.storage)
	copy(newbits, b.bits)
	b.bits = newbits
}

func checkSize(a, b, dest *Bitstring) {
	if b.storage != a.storage {
		panic(fmt.Sprintf("storage sizes of bitstrings not equal! %d != %d", b.storage, a.storage))
	}
	if b.storage != dest.storage {
		panic(fmt.Sprintf("storage size of dest bitstring wrong! %d != %d", b.storage, dest.storage))
	}
}

// Copy the bitstring
func (b *Bitstring) Copy() *Bitstring {
	bits := make([]uint64, b.storage)
	copy(bits, b.bits)
	return &Bitstring{b.limit, b.storage, b.count, bits, b.mask}
}

// And two bitstrings together, storing the result in dest
func (b *Bitstring) And(a, dest *Bitstring) {
	checkSize(a, b, dest)
	dest.count = 0
	for i := uint64(0); i < b.storage; i++ {
		v := b.bits[i] & a.bits[i]
		dest.bits[i] = v
		dest.count += CountBitsSet64(v)
	}
}

// Or two bitstrings together, storing the result in dest
func (b *Bitstring) Or(a, dest *Bitstring) {
	checkSize(a, b, dest)
	dest.count = 0
	last := b.storage - 1
	for i := uint64(0); i < b.storage; i++ {
		v := b.bits[i] | a.bits[i]
		if i == last {
			v &= b.mask
		}
		dest.bits[i] = v
		dest.count += CountBitsSet64(v)
	}
}

// Xor two bitstrings together, storing the result in dest
func (b *Bitstring) Xor(a, dest *Bitstring) {
	checkSize(a, b, dest)
	dest.count = 0
	last := b.storage - 1
	for i := uint64(0); i < b.storage; i++ {
		v := b.bits[i] ^ a.bits[i]
		if i == last {
			v &= b.mask
		}
		dest.bits[i] = v
		dest.count += CountBitsSet64(v)
	}
}

// AndNot two bitstrings together, storing the result in dest
func (b *Bitstring) AndNot(a, dest *Bitstring) {
	checkSize(a, b, dest)
	dest.count = 0
	last := b.storage - 1
	for i := uint64(0); i < b.storage; i++ {
		v := b.bits[i] &^ a.bits[i]
		if i == last {
			v &= b.mask
		}
		dest.bits[i] = v
		dest.count += CountBitsSet64(v)
	}
}

// ClearAll sets all bits to zero
func (b *Bitstring) ClearAll() {
	for i := uint64(0); i < b.storage; i++ {
		b.bits[i] = 0
	}
	b.count = 0
}

// SetAll sets all bits to one
func (b *Bitstring) SetAll() {
	b.count = 0
	last := b.storage - 1
	for i := uint64(0); i < b.storage; i++ {
		b.bits[i] = allOnes
		if i == last {
			b.bits[i] &= b.mask
			b.count += CountBitsSet64(b.bits[i])
		}
	}
}

// OneIterator returns a struct that iterates over the one bits
func (b *Bitstring) OneIterator() *BitstringIterator {
	return newBitstringIterator(b, true)
}

// ZeroIterator returns a struct that iterates over the zero bits
func (b *Bitstring) ZeroIterator() *BitstringIterator {
	return newBitstringIterator(b, false)
}

// Ones returns an array of indices of the 1-bits
func (b *Bitstring) Ones() []uint64 {
	r := make([]uint64, b.count)
	i := 0
	for j := uint64(0); j < b.limit; j++ {
		if b.IsSet(j) {
			r[i] = j
			i++
		}
	}
	return r
}

// AppendOnes appends indices of the 1-bits to the slice
func (b *Bitstring) AppendOnes(dest []uint64) []uint64 {
	for j := uint64(0); j < b.limit; j++ {
		if b.IsSet(j) {
			dest = append(dest, j)
		}
	}
	return dest
}

// Limit is the maximum number stored in the bitstring
func (b *Bitstring) Limit() uint64 {
	return b.limit
}

// Count is the number of ones in the bitstring
func (b *Bitstring) Count() uint64 {
	return b.count
}

// Set sets a bit to one
func (b *Bitstring) Set(bit uint64) {
	i := bit / 64
	if i < b.storage {
		mask := uint64(1) << (bit % 64)
		if (b.bits[i] & mask) == 0 {
			b.count++
		}
		b.bits[i] |= mask
	}
}

// Clear sets a bit to zero
func (b *Bitstring) Clear(bit uint64) {
	i := bit / 64
	if i < b.storage {
		mask := uint64(1) << (bit % 64)
		if (b.bits[i] & mask) != 0 {
			b.count--
		}
		b.bits[i] &^= mask
	}
}

func toBinary(n uint64, buf []byte) []byte {
	for i := uint(0); i < 64; i++ {
		switch i {
		case 8, 16, 24, 32, 40, 48, 56:
			buf = append(buf, byte(' '))
		}
		bit := byte('0' + (1 & (n >> i)))
		buf = append(buf, bit)
	}
	return buf
}

// String returns a string representation of the bitstring
func (b *Bitstring) String() string {
	buf := []byte{}
	for _, block := range b.bits {
		buf = toBinary(block, buf)
		buf = append(buf, byte(' '))
	}
	return fmt.Sprintf("bitstring{%d, %s}", b.storage, string(buf))
}

// IsSet returns true if a bit is set, false otherwise
func (b *Bitstring) IsSet(bit uint64) bool {
	i := bit / 64
	if i < b.storage {
		mask := uint64(1) << (bit % 64)
		return (b.bits[i] & mask) != 0
	}
	return false
}

// Flip inverts the value of the given bit
func (b *Bitstring) Flip(bit uint64) {
	i := bit / 64
	if i < b.storage {
		mask := uint64(1) << (bit % 64)
		if (b.bits[i] & mask) == 0 {
			b.count++
		} else {
			b.count--
		}
		b.bits[i] ^= mask
	}
}

// BitstringIterator provides iteration over the zero or one
// values of a bitstring without building an array
type BitstringIterator struct {
	bits *Bitstring
	n    uint64
	flag int64
	ones bool
}

func newBitstringIterator(bits *Bitstring, ones bool) *BitstringIterator {
	return &BitstringIterator{bits, 0, -1, ones}
}

// Reset ...
func (b *BitstringIterator) Reset() {
	b.n = 0
	b.flag = -1
}

// Next ...
func (b *BitstringIterator) Next() bool {
	if b.flag == -2 {
		return false
	}
	n := b.n + 1
	if b.flag == -1 {
		n = 0
	}

	limit := b.bits.Limit()
	for n < limit {
		v := b.bits.IsSet(n)
		if (b.ones && v) || (!b.ones && !v) {
			b.n = n
			b.flag = 0
			return true
		}
		n++
	}

	b.flag = -2
	return false
}

// Get ..
func (b *BitstringIterator) Get() uint64 {
	switch b.flag {
	case -1:
		panic("must call Next() before accessing iterator value")
	case -2:
		panic("iterator has been drained")
	default:
		return b.n
	}
}
