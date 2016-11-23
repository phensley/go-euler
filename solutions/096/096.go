package euler096

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler096 *.txt

func init() {
	euler.Register("096", "Su Doku", solve)
}

const (
	known = `003020600
900305001
001806400
008102900
700000008
006708200
002609500
800203009
005010300`
)

var allbits bits

func init() {
	for i := uint(1); i <= 9; i++ {
		allbits.set(i)
	}
}

func solve(ctx *euler.Context) {
	p := readPuzzle(strings.Split(known, "\n"))
	if !p.solve() {
		panic(fmt.Sprintf("Failed to solve known puzzle!\n"))
	}

	sum := uint(0)
	for n, p := range loadPuzzles() {
		if p.solve() {
			if euler.Verbose {
				fmt.Printf("Puzzle %d:\n", n+1)
				fmt.Println(p)
			}

			// Sum the numbers in the first 3 columns of the first row
			n := (p.grid[0][0] * 100) + (p.grid[0][1] * 10) + p.grid[0][2]
			sum += n

		} else if euler.Verbose {
			fmt.Printf("Puzzle %d FAILED: ", n+1)
			fmt.Println(p)
		}
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

// Compact version of euler.Bitstring
type bits uint16

func (b *bits) set(n uint) {
	*b = *b | 1<<n
}

func (b *bits) clear(n uint) {
	*b &^= 1 << n
}

func (b *bits) isSet(n uint) bool {
	return *b&(1<<n) != 0
}

func (b *bits) count() uint {
	v := *b
	c := uint(0)
	for v > 0 {
		c++
		v &= v - 1
	}
	return c
}

type puzzle struct {
	// Cells which may be empty or completed
	grid [9][9]uint

	// Record numbers in completed cells in boxes, rows and columns
	boxes [3][3]bits
	rows  [9]bits
	cols  [9]bits

	// Track the total completed cells
	total int
}

// Set a cell's value, possibly overwriting it, and update the
// puzzle state.
func (p *puzzle) set(r, c, n uint) {
	i := r / 3
	j := c / 3
	if v := p.grid[r][c]; v != 0 {
		p.boxes[i][j].clear(v)
		p.rows[r].clear(v)
		p.cols[c].clear(v)
		p.total--
	}

	p.grid[r][c] = n
	p.boxes[i][j].set(n)
	p.rows[r].set(n)
	p.cols[c].set(n)
	p.total++
}

// Undo a cell, setting it to empty, and update the puzzle state.
func (p *puzzle) undo(r, c uint) {
	n := p.grid[r][c]
	p.boxes[r/3][c/3].clear(n)
	p.rows[r].clear(n)
	p.cols[c].clear(n)
	p.total--
	p.grid[r][c] = 0
}

// Find the empty cell having the fewest possibilities. For example,
// the middle cell X can only hold a 9. By solving that cell before
// others we can eliminate more branches.
//
//                5
//    --------+-------+-------+
//    |       | 8     |       |
//    | 2   3 |   X   | 4   1 |
//    |       |   6 7 |       |
//    +-------+-------+-------|
//                1
//
func (p *puzzle) findCell(pChoices *bits, pRow, pCol *uint) bool {
	var maxChoices uint
	for r := uint(0); r < 9; r++ {
		row := p.rows[r]

		// Skip completed rows
		if row == allbits {
			continue
		}

		for c := uint(0); c < 9; c++ {
			col := p.cols[c]

			// Skip completed cols and individual cells
			if col == allbits || p.grid[r][c] != 0 {
				continue
			}

			// For a given empty cell, collect all completed numbers in its
			// same row, col and box.
			choices := row | col | p.boxes[r/3][c/3]

			// Skip if count is not our best choice
			num := choices.count()
			if num <= maxChoices {
				continue
			}

			// Check for conflict
			if choices == allbits {
				return false
			}

			// Track the cell coordinates (row,col) with the most
			// possibilities eliminated
			maxChoices = num
			*pRow = r
			*pCol = c
			*pChoices = choices

			if num == 8 {
				return true
			}
		}
	}

	return true
}

// Dumb, backtracking, brute-force search. Has the benefit of (a) using very fast operations
// to record the puzzle state and test possibilities, (b) applying a very simple heuristic,
// solving empty cells with the fewest possibilities first, and (c) no copies of the data
// structure are required as at each level we modify a single empty cell, and if that branch
// fails to find a solution we backtrack by undoing that single cell.
func (p *puzzle) solve() bool {
	if p.total == 81 {
		return true
	}

	var choices bits
	var row, col uint

	// Find the empty cell with the fewest possibilities and try each
	// in succession. For example, if a given cell can only contain 2 or 3.
	// We set 2 and start a deeper search. If that fails, we backtrack and
	// try 3, and so on.
	if !p.findCell(&choices, &row, &col) {
		return false
	}

	for n := uint(1); n <= 9; n++ {
		// Try each of the possibilities for this cell, recursing
		// one level deeper.
		if !choices.isSet(n) {
			p.set(row, col, n)
			if p.solve() {
				return true
			}
		}
	}

	// If this level's search failed, undo the cell.
	p.undo(row, col)
	return false
}

func (p *puzzle) String() string {
	buf := bytes.Buffer{}
	for r := uint(0); r < 9; r++ {
		if r > 0 && r%3 == 0 {
			buf.WriteRune('\n')
		}
		for c := uint(0); c < 9; c++ {
			if c > 0 && c%3 == 0 {
				buf.WriteRune(' ')
			}
			n := p.grid[r][c]
			if n == 0 {
				buf.WriteString("_ ")
			} else {
				buf.WriteRune('0' + rune(n))
				buf.WriteRune(' ')
			}
		}
		buf.WriteRune('\n')
	}
	return string(buf.Bytes())
}

func loadPuzzles() []*puzzle {
	data := string(rawfiles["p096_sudoku.txt"])
	lines := strings.Split(data, "\n")
	i := 0
	res := []*puzzle{}
	for i < len(lines) {
		p := readPuzzle(lines[i+1 : i+10])
		res = append(res, p)
		i += 10
	}
	return res
}

func readPuzzle(rows []string) *puzzle {
	if len(rows) != 9 {
		panic(fmt.Sprintf("cannot construct sudoku board from %d rows", len(rows)))
	}

	p := &puzzle{}
	for r := uint(0); r < 9; r++ {
		col := strings.Split(strings.Trim(rows[r], " \r\n"), "")
		if len(col) != 9 {
			panic(fmt.Sprintf("cannot construct sudoku row from column of length %d: %v", len(col), col))
		}

		for c := uint(0); c < 9; c++ {
			n, err := strconv.ParseInt(col[c], 10, 8)
			if err != nil {
				panic(fmt.Sprintf("invalid digit %s found at (row,col) (%d,%d)", col[c], r, c))
			}
			if n != 0 {
				p.set(r, c, uint(n))
			}
		}
	}
	return p
}
