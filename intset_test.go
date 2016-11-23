package euler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
	set := NewIntSet(0, 100, 0.85)
	assert.False(t, set.Contains(123))
	set.Add(123)
	assert.True(t, set.Contains(123))
	assert.False(t, set.Contains(122))

	for i := 1; i <= 200; i++ {
		set.Add(i)
	}

	assert.Equal(t, int32(200), set.size, "set size")
	assert.Equal(t, int32(256), set.capacity, "set capacity")

	set.Clear()
	assert.Equal(t, int32(0), set.size, "set size")
	assert.Equal(t, int32(256), set.capacity, "set capacity")
}

func TestIntSetNegative(t *testing.T) {
	set := NewIntSet(-256, 100, 0.85)
	set.Add(-5)
	assert.True(t, set.Contains(-5))
	set.Add(-10)
	assert.True(t, set.Contains(-5))

	for i := -200; i < 200; i++ {
		set.Add(i)
	}

	assert.Equal(t, int32(400), set.size, "size")
	assert.Equal(t, int32(512), set.capacity, "capacity")

	assert.Panics(t, func() { set.Contains(-256) }, "contains(free)")
	assert.Panics(t, func() { set.Add(-256) }, "add(free)")
}
