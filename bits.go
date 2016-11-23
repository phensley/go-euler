package euler

// CountBitsSet32 counts the number of bits set in a 32-bit unsigned integer
func CountBitsSet32(x uint32) uint32 {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0F0F0F0F
	x = x + (x >> 8)
	x = x + (x >> 16)
	return x & 0x0000003F
}

// CountBitsSet64 counts the number of bits set in a 64-bit unsigned integer
func CountBitsSet64(x uint64) uint64 {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	return (((x + (x >> 4)) & 0xF0F0F0F0F0F0F0F) * 0x101010101010101) >> 56
}
