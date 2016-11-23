package euler043

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("043", "Sub-string divisibility", solve)
}

func solve(ctx *euler.Context) {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := uint(0)
	perms := euler.NewPermutations(digits)
	for perms.Next() {
		d := perms.Get()

		// Skip numbers with a leading zero
		if d[0] == 0 {
			continue
		}

		if (euler.DigitsToInt(d[1:4]) % 2) != 0 {
			continue
		}
		if (euler.DigitsToInt(d[2:5]) % 3) != 0 {
			continue
		}
		if (euler.DigitsToInt(d[3:6]) % 5) != 0 {
			continue
		}
		if (euler.DigitsToInt(d[4:7]) % 7) != 0 {
			continue
		}
		if (euler.DigitsToInt(d[5:8]) % 11) != 0 {
			continue
		}
		if (euler.DigitsToInt(d[6:9]) % 13) != 0 {
			continue
		}
		if (euler.DigitsToInt(d[7:10]) % 17) != 0 {
			continue
		}

		// Number has all properties. Convert from digits back to integer form
		n := euler.DigitsToInt(d)
		if euler.Verbose {
			fmt.Println("Found: ", n)
		}
		sum += uint(n)
	}
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}
