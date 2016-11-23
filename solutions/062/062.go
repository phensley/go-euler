package euler062

import (
	"fmt"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("062", "Cubic permutations", solve)
}

func solve(ctx *euler.Context) {
	limit := uint64(10000)
	nums := map[string][]uint64{}
	ordered := []uint64{}
	for i := uint64(2); i < limit; i++ {
		n := cube(i)
		ordered = append(ordered, n)
		d := euler.UintToDigits(uint(n))
		sort.Sort(sort.Interface(uintlist(d)))
		key := fmt.Sprintf("%v", d)
		nums[key] = append(nums[key], n)
	}

	for _, n := range ordered {
		d := euler.UintToDigits(uint(n))
		sort.Sort(sort.Interface(uintlist(d)))
		key := fmt.Sprintf("%v", d)

		if len(nums[key]) == 5 {
			if euler.Verbose {
				fmt.Println(nums[key])
			}
			answer := fmt.Sprintf("%d", n)
			ctx.SetAnswer(answer)
			return
		}
	}
}

func cube(n uint64) uint64 {
	return n * n * n
}

type uintlist []uint

func (u uintlist) Len() int {
	return len(u)
}

func (u uintlist) Less(i, j int) bool {
	return u[i] < u[j]
}

func (u uintlist) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
