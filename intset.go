package euler

import (
	"fmt"
	"math"
)

// IntSet holds a set of integers
type IntSet struct {
	free       int
	capacity   int32
	loadFactor float32
	mask       int32
	keys       []int
	size       int32
	threshold  int32
}

// NewIntSet creates an IntSet
func NewIntSet(free int, capacity int32, loadFactor float32) *IntSet {
	// Ensure capacity is a multiple of 32
	cap := int32(math.Ceil(float64(capacity)/float64(32)) * 32)
	mask := cap - 1
	threshold := int32(math.Floor(float64(cap) * float64(loadFactor)))
	keys := make([]int, cap)
	if free != 0 {
		zero(free, keys)
	}
	return &IntSet{free, cap, loadFactor, mask, keys, 0, threshold}
}

// Clear ...
func (s *IntSet) Clear() {
	s.size = 0
	zero(s.free, s.keys)
}

// Add ...
func (s *IntSet) Add(key int) {
	s.checkKey(key)
	s.grow()
	index := s.locate(key)
	if s.keys[index] == s.free {
		s.size++
		s.keys[index] = key
	}
}

// Contains ...
func (s *IntSet) Contains(key int) bool {
	s.checkKey(key)
	return s.search(key) != -1
}

func zero(free int, arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = free
	}
}

func (s *IntSet) grow() {
	if s.size < s.threshold {
		return
	}

	cap := int32(len(s.keys) * 2)
	mask := cap - 1
	keys := make([]int, cap)
	if s.free != 0 {
		zero(s.free, keys)
	}

	// Rehash the old array into the new
	for i := len(s.keys) - 1; i >= 0; i-- {
		key := s.keys[i]
		if key != s.free {
			index := intmapHash(key) & mask
			for keys[index] != s.free {
				index = (index + 1) & mask
			}
			keys[index] = key
		}
	}

	s.capacity = cap
	s.mask = mask
	s.keys = keys
	s.threshold = int32(math.Floor(float64(cap) * float64(s.loadFactor)))
}

func intmapHash(key int) int32 {
	key ^= key >> 16
	key *= 0x85ebca6b
	key ^= key >> 13
	key *= 0xc2b2ae35
	key ^= key >> 16
	return int32(key)
}

func (s *IntSet) checkKey(key int) {
	if key == s.free {
		panic(fmt.Sprintf("attempt use the 'free slot' indicator %#v as a value", s.free))
	}
}

func (s *IntSet) locate(key int) int32 {
	index := intmapHash(key) & s.mask
	for s.keys[index] != s.free {
		if s.keys[index] == key {
			return index
		}
		index = (index + 1) & s.mask
	}
	return index
}

func (s *IntSet) search(key int) int32 {
	index := intmapHash(key) & s.mask
	for s.keys[index] != s.free {
		if key == s.keys[index] {
			return index
		}
		index = (index + 1) & s.mask
	}
	return -1
}
