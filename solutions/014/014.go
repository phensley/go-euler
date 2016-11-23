package euler014

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("014", "Longest Collatz sequence", solve)
}

func solve(ctx *euler.Context) {
	limit := 1000000

	if euler.Verbose {
		known := calculate(13)
		fmt.Println("known collatz(13) length: ", known)
	}

	longest := 0
	length := 0
	for n := 2; n < limit; n++ {
		l := calculate(n)
		if length < l {
			longest = n
			length = l
		}
	}

	if euler.Verbose {
		fmt.Println("sequence length is", length)
	}
	answer := fmt.Sprintf("%d", longest)
	ctx.SetAnswer(answer)
}

func calculate(n int) int {
	count := 1
	for {
		if n&1 == 0 {
			n = n >> 1
		} else {
			n = n*3 + 1
		}
		count++
		if n == 1 {
			break
		}
	}
	return count
}
