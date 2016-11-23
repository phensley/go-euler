package euler099

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler099 *.txt

func init() {
	euler.Register("099", "Largest exponential", solve)
}

const precision = 50

func solve(ctx *euler.Context) {
	answer := fmt.Sprintf("%d", calculate())
	ctx.SetAnswer(answer)
}

func calculate() int {
	data := string(rawfiles["p099_base_exp.txt"])
	lines := strings.Split(data, "\n")
	nums := make(numberlist, len(lines))
	for i, line := range lines {
		b, e := splitLine(line)
		n := big.NewFloat(math.Log10(b)).SetPrec(precision)
		n.Mul(n, big.NewFloat(e).SetPrec(precision))
		nums[i] = &number{i + 1, n}
	}
	sort.Sort(sort.Reverse(nums))
	return nums[0].line
}

func splitLine(line string) (float64, float64) {
	row := strings.Split(line, ",")
	if len(row) != 2 {
		panic("bad row length!")
	}
	b := parseFloat(row[0])
	e := parseFloat(row[1])
	return b, e
}

func parseFloat(r string) float64 {
	n, err := strconv.ParseFloat(r, 64)
	if err != nil {
		panic(err)
	}
	return n
}

type number struct {
	line int
	n    *big.Float
}

func (n number) String() string {
	return fmt.Sprintf("{%d} %.30f", n.line, n.n)
}

type numberlist []*number

func (n numberlist) Len() int { return len(n) }

func (n numberlist) Less(i, j int) bool { return n[i].n.Cmp(n[j].n) < 0 }

func (n numberlist) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
