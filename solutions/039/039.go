package euler039

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("039", "Integer right triangles", solve)
}

func solve(ctx *euler.Context) {
	// First, build up a table of squares <= limit
	// All our triangles will obviously have a longest side < the perimeter
	limit := uint64(1000)
	squares := euler.NewBitString((limit * limit) + uint64(1))
	for i := uint64(1); i <= limit; i++ {
		r := i * i
		squares.Set(uint64(r))
	}

	triangles := make(map[int][]triangle)

	// Next, find sums that are also members of the set
	sq := squares.Ones()
	for _, x := range sq {
		for _, y := range sq {
			z := x + y
			if squares.IsSet(z) {
				a := sqrt(x)
				b := sqrt(y)
				c := sqrt(z)
				p := a + b + c
				if uint64(p) > limit {
					continue
				}

				t := triangle{a, b, c}
				triangles[p] = append(triangles[p], t)
			}
		}
	}

	longest := 0
	length := 0
	for p, list := range triangles {
		if len(list) > length {
			length = len(list)
			longest = p
		}
	}

	if euler.Verbose {
		fmt.Println("permiter", longest, "has", length, "solutions")
	}
	answer := fmt.Sprintf("%d", longest)
	ctx.SetAnswer(answer)
}

type triangle struct {
	a int
	b int
	c int
}

func sqrt(n uint64) int {
	return int(math.Sqrt(float64(n)))
}
