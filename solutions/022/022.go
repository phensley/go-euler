package euler022

import (
	"fmt"
	"sort"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler022 *.txt

func init() {
	euler.Register("022", "Names scores", solve)
}

func solve(ctx *euler.Context) {
	text := string(rawfiles["p022_names.txt"])
	lines := strings.Split(text, ",")
	names := []string{}
	for _, line := range lines {
		end := len(line) - 1
		names = append(names, line[1:end])
	}
	sort.Strings(names)

	total := 0
	for pos, name := range names {
		score := (pos + 1) * nameScore(name)
		total += score
	}
	answer := fmt.Sprintf("%d", total)
	ctx.SetAnswer(answer)
}

func nameScore(name string) int {
	start := byte('A') - 1
	score := 0
	for _, ch := range name {
		val := int(byte(ch) - start)
		score += val
	}
	return score
}
