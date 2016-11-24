package euler089

import (
	"fmt"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler089 *.txt

func init() {
	euler.Register("089", "Roman numerals", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		// Sanity check with knowns
		known := []string{
			"IIIIIIIIIIIIIIII",
			"VIIIIIIIIIII",
			"VVIIIIII",
			"XIIIIII",
			"VVVI",
			"XVI",
		}
		for _, num := range known {
			fmt.Println(num, canonicalize(num))
		}
	}

	data := rawfiles["p089_roman.txt"]
	saved := 0
	for _, num := range strings.Split(data, "\n") {
		// Find canonical form
		valid := canonicalize(num)

		// Track the total characters we saved by canonicalizing
		saved += len(num) - len(valid)
	}

	answer := fmt.Sprintf("%d", saved)
	ctx.SetAnswer(answer)
}

// Canonicalize a Roman number by parsing and then encoding it
func canonicalize(raw string) string {
	n := euler.RomanNumeralDecode(raw)
	return euler.RomanNumeralEncode(n)
}
