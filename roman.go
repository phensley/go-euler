package euler

import "fmt"

// RomanNumeralEncode ...
func RomanNumeralEncode(n int) string {
	s := ""
	for _, r := range romanNumerals {
		for n >= r.n {
			s += r.s
			n -= r.n
		}
	}
	return s
}

// RomanNumeralDecode ...
func RomanNumeralDecode(roman string) int {
	last := len(roman) - 1
	s := 0
	i := 0
	for i <= last {
		ch := roman[i]
		p := romanDenominations[ch]
		if p == 0 {
			panic(fmt.Sprintf("invalid value: %v", ch))
		}

		if i == last {
			s += p
			break
		}

		// Peek at next character.
		ch = roman[i+1]
		n := romanDenominations[ch]
		if p < n {
			s -= p
		} else {
			s += p
		}
		i++
	}
	return s
}

type roman struct {
	n int
	s string
}

var (
	romanNumerals = []roman{
		roman{1000, "M"},
		roman{900, "CM"},
		roman{500, "D"},
		roman{400, "CD"},
		roman{100, "C"},
		roman{90, "XC"},
		roman{50, "L"},
		roman{40, "XL"},
		roman{10, "X"},
		roman{9, "IX"},
		roman{5, "V"},
		roman{4, "IV"},
		roman{1, "I"},
	}

	romanDenominations = map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
)
