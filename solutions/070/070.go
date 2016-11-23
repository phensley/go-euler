package euler070

import (
	"fmt"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("070", "Totient permutation", solve)
}

func solve(ctx *euler.Context) {
	min := uint64(0)
	max := uint64(10000000)
	minratio := float64(0)

	for n, phi := range euler.EulersTotient(max) {
		if n < 2 {
			continue
		}

		// Filter out numbers whose digit sums are not equal.
		if euler.IntDigitSum(int(n)) != euler.IntDigitSum(int(phi)) {
			continue
		}

		dn := euler.UintToDigits(uint(n))
		dphi := euler.UintToDigits(uint(phi))
		sort.Sort(uintlist(dn))
		sort.Sort(uintlist(dphi))
		if equal(dn, dphi) {
			ratio := float64(n) / float64(phi)
			if minratio == 0 || ratio < minratio {
				minratio = ratio
				min = uint64(n)
			}

		}
	}

	if euler.Verbose {
		fmt.Println("ratio n/phi(n): ", minratio)
	}

	answer := fmt.Sprintf("%d", min)
	ctx.SetAnswer(answer)
}

func equal(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type uintlist []uint

func (a uintlist) Len() int           { return len(a) }
func (a uintlist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a uintlist) Less(i, j int) bool { return a[i] < a[j] }
