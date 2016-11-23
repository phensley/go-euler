package euler036

import (
	"fmt"
	"math"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("036", "Double-base palindromes", solve)
}

func solve(ctx *euler.Context) {
	p := palindrome{make(map[int][]uint)}
	limit := 1000000
	maxPower := int(math.Log2(float64(limit)) + 1)
	binary := []uint{}
	result := []uint{}
	for i := 0; i < maxPower; i++ {
		for _, r := range p.calc(limit, i) {
			if r < 1000000 {
				binary = append(binary, r)
				if euler.IsPalindrome(uint(r)) {
					result = append(result, r)
				}
			}
		}
	}

	if euler.Verbose {
		fmt.Printf("count base 2 palindromes < %d = %d\n", limit, len(binary))
		fmt.Printf("count that are also base 10 palindromes = %d\n", len(result))
	}

	sum := uint(0)
	for _, r := range result {
		if euler.Verbose {
			fmt.Printf("%b  %d\n", r, r)
		}
		sum += r
	}
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

// Exploit the fact that a binary palindrome must start with a 1. Thus it must also
// end in a 1.  Generate a series of numbers for each power of 2 up to our limit.
// For power == 4, we have a range:

//   1 << 4) + 1 == 10001
//  (1 << 5) - 1 == 11111
//
// In order to fill in the 000 in the middle with palindromes, calculate the
// palindrome numbers for power == 2 and their bitwise inverse.  Then multiply
// by 2, shifting the digits 1 place to the left, and bitwise OR.
//
// This produces the sequence:
//
//   10001
//   11011
//   10101
//   11111
//
// So power = 9 will recursively use the palindromes for power = 7, which uses
// those of power 5, and so on.  We cache the palindromes for each power for efficiency.

type uintlist []uint

func (u uintlist) Len() int { return len(u) }

func (u uintlist) Less(i, j int) bool { return u[i] < u[j] }

func (u uintlist) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

type palindrome struct {
	cache map[int][]uint
}

// Generate a sequence of palindromic binary numbers
func (p *palindrome) calc(limit, power int) []uint {
	if power <= 0 {
		return []uint{1}
	}
	if r := p.cache[power]; r != nil {
		return r
	}

	// calculate this segment's upper and lower bounds
	lo := (1 << uint(power)) + 1
	up := (1 << (uint(power) + 1)) - 1

	flags := euler.NewBitString(uint64(limit + 1))

	res := uintlist{}
	if lo != 0 {
		res = append(res, uint(lo))
		flags.Set(uint64(lo))
	}

	// Collect the palindromes for current power - 2.
	// This loop populates the string of 0 digits in the center
	// of `lo` with a prior sequence of palindromes for
	// power - 2.
	for _, n := range p.calc(limit, power-2) {
		// Bitwise OR the shifted prior palindrome
		r := uint(lo) | (n << uint(1))
		if r != 0 && !flags.IsSet(uint64(r)) {
			flags.Set(uint64(r))
			res = append(res, r)
		}

		// The bitwise negation of a palindrome is also a palindrome
		r = (uint(lo) | ^(n << uint(1))) & uint(up)
		if r != 0 && !flags.IsSet(uint64(r)) {
			flags.Set(uint64(r))
			res = append(res, r)
		}
	}
	if !flags.IsSet(uint64(up)) {
		flags.Set(uint64(up))
		res = append(res, uint(up))
	}

	sort.Sort(sort.Interface(res))
	p.cache[power] = res
	return res
}
