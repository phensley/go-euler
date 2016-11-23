package euler020

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("020", "Factorial digit sum", solve)
}

func solve(ctx *euler.Context) {
	f := euler.BigFactorial{}
	digits := f.Calculate(uint32(100)).String()
	if euler.Verbose {
		fmt.Println(digits)
	}
	ch := byte('0')
	sum := 0
	for _, d := range digits {
		sum += int(byte(d) - ch)
	}
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func factorial(n int64) string {
	a := big.Int{}
	a.SetInt64(n)
	b := big.Int{}
	for n > 1 {
		n--
		b.SetInt64(n)
		a.Mul(&a, &b)
	}
	return a.String()
}
