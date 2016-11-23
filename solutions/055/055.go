package euler055

import (
	"fmt"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("055", "Lychrel numbers", solve)
}

func solve(ctx *euler.Context) {
	count := 0
	for n := 0; n < 10000; n++ {
		if lychrel(n) {
			count++
		}
	}

	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

func lychrel(n int) bool {
	for i := 0; i < 50; i++ {
		d := euler.IntToDigits(n)
		euler.ReverseSortable(sort.IntSlice(d))
		s := n + euler.DigitsToInt(d)
		if euler.IsPalindrome(uint(s)) {
			return false
		}
		n = s
	}
	return true
}
