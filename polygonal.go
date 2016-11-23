package euler

// Generate a sequence of numbers from min to max using the generator function
func polygonal(min, max uint64, generator func(uint64) uint64) *Bitstring {
	r := NewBitString(max)
	i := uint64(1)
	for {
		n := generator(i)
		if n >= max {
			break
		}
		if n > min {
			r.Set(n)
		}
		i++
	}
	return r
}

// Square numbers up to N
func Square(min, max uint64) *Bitstring {
	return polygonal(min, max, func(n uint64) uint64 {
		return n * n
	})
}

// Triangle numbers up to N
func Triangle(min, max uint64) *Bitstring {
	return polygonal(min, max, func(n uint64) uint64 {
		return (n * (n + 1)) / 2
	})
}

// Pentagonal numbers up to N
func Pentagonal(min, max uint64) *Bitstring {
	return polygonal(min, max, func(n uint64) uint64 {
		return (n * ((3 * n) - 1)) / 2
	})
}

// Hexagonal numbers up to N
func Hexagonal(min, max uint64) *Bitstring {
	return polygonal(min, max, func(n uint64) uint64 {
		return n * ((2 * n) - 1)
	})
}

// Heptagonal numbers up to N
func Heptagonal(min, max uint64) *Bitstring {
	return polygonal(min, max, func(n uint64) uint64 {
		return (n * ((5 * n) - 3)) / 2
	})
}

// Octagonal numbers up to N
func Octagonal(min, max uint64) *Bitstring {
	return polygonal(min, max, func(n uint64) uint64 {
		return n * ((3 * n) - 2)
	})
}
