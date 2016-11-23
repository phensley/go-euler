package euler023

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("023", "Non-abundant sums", solve)
}

func solve(ctx *euler.Context) {
	limit := 28123
	nums := euler.NewBitString(uint64(limit + 1))
	abundant := euler.NewBitString(uint64(limit + 1))
	for n := 1; n <= limit; n++ {
		nums.Set(uint64(n))
		if n%2 != 0 && euler.IsPrimeForisekJancina32(uint32(n)) {
			continue
		}
		if n < euler.SumOfProperDivisors(n) {
			abundant.Set(uint64(n))
		}
	}

	abundantNums := abundant.Ones()
	for _, a := range abundantNums {
		for _, b := range abundantNums {
			s := uint64(a + b)
			if s <= nums.Limit() {
				nums.Clear(s)
			}
		}
	}

	sum := 0
	iter := nums.OneIterator()
	for iter.Next() {
		sum += int(iter.Get())
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}
