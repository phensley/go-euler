package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMul64(t *testing.T) {
	cases := [][]uint64{
		{0, 1, 0, 0},
		{1, 1, 0, 1},
		{3, 11, 0, 33},
		{0x5555, 0xAAAA, 0, 0x38E31C72},
		{0, 0xFFFFFFFFFFFFFFFF, 0, 0},
		{1, 0xFFFFFFFFFFFFFFFF, 0, 0xFFFFFFFFFFFFFFFF},
		{4, 0xAAAAAAAAAAAAAAAA, 2, 0xAAAAAAAAAAAAAAA8},
		{0x1000000000000000, 0x1000000000000000, 0x100000000000000, 0},
		{0x1000000000000000, 0x2000000000000000, 0x200000000000000, 0},
		{0x1000000000000000, 0x2000000000000001, 0x200000000000000, 0x1000000000000000},
		{0x1234567800000000, 0x8765432100000000, 0x9A0CD0570B88D78, 0},
		{0x1234567812345678, 0x0000000087654321, 0x000000009A0CD05, 0x7A595A7D70B88D78},
	}

	for _, m := range cases {
		u, v := m[0], m[1]
		rh, rl := Mul64(u, v)
		if rh != m[2] || rl != m[3] {
			t.Errorf("%x * %x: expected {%x, %x} got {%x, %x}", u, v, m[2], m[3], rh, rl)
		}
	}
}

func TestCountBits64(t *testing.T) {
	n := uint64(1)
	for i := uint64(0); i < 64; i++ {
		b := CountBitsSet64(n << i)
		assert.Equal(t, uint64(1), b)
	}
}

func TestRoundFloat64(t *testing.T) {
	for _, n := range []float64{0.0, 0.1, 0.49999, 1.0, 1.1, 1.49999} {
		r := RoundFloat64(n)
		assert.Equal(t, int64(n), r)
	}
	for _, n := range []float64{0.5, 0.500001, 0.6, 1.5, 1.500001, 1.6} {
		r := RoundFloat64(n)
		assert.Equal(t, int64(n)+1, r)
	}
	for _, n := range []float64{-0.1, -0.49999, -1.0, -1.1, -1.49999} {
		r := RoundFloat64(n)
		assert.Equal(t, int64(n), r)
	}
	for _, n := range []float64{-0.5, -0.500001, -0.6, -1.5, -1.500001, -1.6} {
		r := RoundFloat64(n)
		assert.Equal(t, int64(n)-1, r)
	}
}
