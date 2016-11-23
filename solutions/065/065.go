package euler065

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("065", "Convergents of e", solve)
}

func solve(ctx *euler.Context) {
	e := partialDigitsOfE(100)
	n, _ := continuedFraction(100, e)

	sum := 0
	str := n.String()
	for i := 0; i < len(str); i++ {
		digit := int(str[i] - '0')
		sum += digit
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func partialDigitsOfE(num int) []int {
	r := make([]int, num)
	n := 1
	j := 1
	for i := 0; i < num; i++ {
		if i == 0 {
			r[i] = 2
		} else if j == 3 {
			r[i] = 2 * n
			j = 0
			n++
		} else {
			r[i] = 1
		}
		j++
	}
	return r
}

func continuedFraction(convergent int, digits []int) (*big.Int, *big.Int) {
	n := big.NewInt(1)
	d := big.NewInt(int64(digits[convergent-1]))
	for i := convergent - 2; i >= 0; i-- {
		x := big.NewInt(int64(digits[i]))
		x.Mul(x, d)
		n.Add(n, x)
		d, n = n, d
	}
	return d, n
}
