package euler019

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("019", "Counting Sundays", solve)
}

func solve(ctx *euler.Context) {
	dayOfWeek := monday
	count := 0

	for year := uint(1900); year <= 2000; year++ {
		for month := january; month <= december; month++ {
			if year > uint(1900) && dayOfWeek == sunday {
				if euler.Verbose {
					fmt.Printf("%d %s\n", year, months[month])
				}
				count++
			}
			delta := daysInMonth(month, year)
			dayOfWeek = (dayOfWeek + delta) % uint(7)
		}
	}
	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

const (
	january = uint(iota)
	february
	march
	april
	may
	june
	july
	august
	september
	october
	november
	december
)

const (
	sunday = uint(iota)
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

var (
	months = map[uint]string{
		january:   "jan",
		february:  "feb",
		march:     "mar",
		april:     "apr",
		may:       "may",
		june:      "jun",
		july:      "jul",
		august:    "aug",
		september: "sep",
		october:   "oct",
		november:  "nov",
		december:  "dec",
	}
)

func isLeapYear(year uint) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}

func daysInMonth(month, year uint) uint {
	switch month {
	case september, april, june, november:
		return 30
	case february:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 31
	}
}
