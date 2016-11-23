package euler

// IsPalindrome returns true of the unsigned int is a palindrome,
// e.g. 12321
func IsPalindrome(n uint) bool {
	d := UintToDigits(n)
	length := len(d)
	for i := 0; i < length/2; i++ {
		if d[i] != d[length-1-i] {
			return false
		}
	}
	return true
}

// IntDigitSum returns the sum of the digits of n
func IntDigitSum(n int) int {
	sum := 0
	for n >= 10 {
		sum += n % 10
		n /= 10
	}
	if n > 0 {
		sum += n
	}
	return sum
}

// SumInts ...
func SumInts(n []int) uint64 {
	r := uint64(0)
	for _, i := range n {
		r += uint64(i)
	}
	return r
}

// UintToDigits converts unsigned int into a slice of its digits
func UintToDigits(n uint) []uint {
	size := uint(NumDigitsBase10(uint64(n)))
	res := make([]uint, size)
	i := size - 1
	for n >= 10 {
		res[i] = n % 10
		n /= 10
		i--
	}
	if n > 0 {
		res[i] = n
	}
	return res
}

// IntToDigits converts unsigned int into a slice of its digits
func IntToDigits(n int) []int {
	size := uint(NumDigitsBase10(uint64(n)))
	res := make([]int, size)
	i := size - 1
	nn := uint(n)
	for nn >= 10 {
		res[i] = int(nn % 10)
		nn /= 10
		i--
	}
	if nn > 0 {
		res[i] = int(nn)
	}
	return res
}

// DigitsToUint convert digits to unsigned int
func DigitsToUint(n []uint) uint {
	f := uint(1)
	r := uint(0)
	for i := len(n) - 1; i >= 0; i-- {
		r += f * n[i]
		f *= 10
	}
	return r
}

// DigitsToInt convert digits to int
func DigitsToInt(n []int) int {
	f := 1
	r := 0
	for i := len(n) - 1; i >= 0; i-- {
		r += f * n[i]
		f *= 10
	}
	return r
}
