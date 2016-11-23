package euler024

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("024", "Lexicographic permutations", solve)
}

func solve(ctx *euler.Context) {
	count := 1
	perms := euler.NewPermutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for perms.Next() {
		p := perms.Get()
		if count == 1000000 {
			answer := fmt.Sprintf("%d", euler.DigitsToInt(p))
			ctx.SetAnswer(answer)
			return
		}
		count++
	}
}
