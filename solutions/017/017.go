package euler017

import (
	"fmt"

	"github.com/phensley/go-euler"
)

// Problem 17 - Number letter counts

func init() {
	euler.Register("017", "Number letter counts", solve)
}

func solve(ctx *euler.Context) {
	count := 0
	for n := 1; n <= 1000; n++ {
		words := makeWord(n)
		wlen := 0
		for _, w := range words {
			wlen += len(w)
		}
		count += wlen
		if euler.Verbose {
			fmt.Println(n, " ", words)
		}
	}
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

var (
	oneToNine = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	tenToNineteen = []string{
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	twentyToNinety = []string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
)

func makeWord(n int) []string {
	var w []string
	th := n / 1000
	n -= th * 1000
	hd := n / 100
	n -= hd * 100

	flag := false
	if th > 0 {
		w = append(w, oneToNine[th-1])
		w = append(w, "thousand")
		flag = true
	}
	if hd > 0 {
		w = append(w, oneToNine[hd-1])
		w = append(w, "hundred")
		flag = true
	}

	if flag && n > 0 {
		w = append(w, "and")
	}

	switch {
	case n == 0:
	case n < 10:
		w = append(w, oneToNine[n-1])
	case n < 20:
		w = append(w, tenToNineteen[n-10])
	default:
		r := n % 10
		n /= 10
		if n > 0 {
			w = append(w, twentyToNinety[n-2])
		}
		if r > 0 {
			w = append(w, oneToNine[r-1])
		}
	}
	return w
}
