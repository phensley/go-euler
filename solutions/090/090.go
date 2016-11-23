package euler090

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("090", "Cube digit pairs", solve)
}

func solve(ctx *euler.Context) {
	count := calculate()
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

func calculate() int {
	squares := [][]int{
		[]int{0, 1},
		[]int{0, 4},
		[]int{0, 6}, // 09 with 6 swapped
		[]int{1, 6},
		[]int{2, 5},
		[]int{3, 6},
		[]int{4, 6}, // 49 with 6 swapped
		[]int{6, 4},
		[]int{8, 1},
	}

	// Digits from 0-9 with 6 replacing nine.
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 6}

	// Find all combinations of the digits of length 6 and build a list
	// of sets containing the digits. This is (10 choose 6), or 210
	// distinct combinations.
	sets := []map[int]*struct{}{}
	for _, c := range euler.Combinations(digits, 6) {
		m := make(map[int]*struct{})
		for _, n := range c {
			m[n] = &struct{}{}
		}
		sets = append(sets, m)
	}

	// Iterate over the 2-digit squares and return true if all squares are satisfied
	// by digits in the two given sets.
	check := func(a, b map[int]*struct{}) bool {
		for _, s := range squares {
			x, y := s[0], s[1]
			v := (a[x] != nil && b[y] != nil) || (a[y] != nil && b[x] != nil)
			if !v {
				return false
			}
		}
		return true
	}

	// Count the distinct pairs of sets whose digits can produce all of the
	// 2-digit squares.
	count := 0
	for i := 0; i < len(sets); i++ {
		for j := i + 1; j < len(sets); j++ {
			if check(sets[i], sets[j]) {
				count++
			}
		}
	}
	return count
}
