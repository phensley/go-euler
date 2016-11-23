package euler084

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoves(t *testing.T) {
	assert.Equal(t, spaceC1, move(spaceA1, 10))
	assert.Equal(t, spaceA2, move(spaceCH3, 7))
}

func TestRolls(t *testing.T) {
	counts := map[pos]int{}
	doubles := 0
	lim := 1000000
	sides := 6
	for i := 0; i < lim; i++ {
		d1, d2 := roll(sides)
		counts[d1]++
		counts[d2]++
		if d1 == d2 {
			doubles++
		}
	}

	avg := int((float64(lim) * 2) / float64(sides))
	for d, v := range counts {
		fmt.Println(d, " is ", avg-v)
	}
	fmt.Println(counts)
	fmt.Println(float64(doubles) / float64(lim))
}

func TestCHDeck(t *testing.T) {
	chDeck := buildCHDeck()
	assert.Equal(t, spaceGO, chDeck[0](spaceCH1))
	assert.Equal(t, spaceJAIL, chDeck[1](spaceCH1))
	assert.Equal(t, spaceC1, chDeck[2](spaceCH1))
	assert.Equal(t, spaceE3, chDeck[3](spaceCH1))
	assert.Equal(t, spaceH2, chDeck[4](spaceCH1))
	assert.Equal(t, spaceR1, chDeck[5](spaceCH1))

	// Railways
	assert.Equal(t, spaceR2, chDeck[6](spaceCH1))
	assert.Equal(t, spaceR3, chDeck[6](spaceCH2))
	assert.Equal(t, spaceR1, chDeck[6](spaceCH3))

	assert.Equal(t, spaceR2, chDeck[7](spaceCH1))
	assert.Equal(t, spaceR3, chDeck[7](spaceCH2))
	assert.Equal(t, spaceR1, chDeck[7](spaceCH3))

	// Utilities
	assert.Equal(t, spaceU1, chDeck[8](spaceCH1))
	assert.Equal(t, spaceU2, chDeck[8](spaceCH2))

	// Back 3 spaces
	assert.Equal(t, spaceT1, chDeck[9](spaceCH1))
	assert.Equal(t, spaceD3, chDeck[9](spaceCH2))
	assert.Equal(t, spaceCC3, chDeck[9](spaceCH3))

	// Ignore
	for i := 10; i < 16; i++ {
		assert.Equal(t, spaceCH1, chDeck[i](spaceCH1))
		assert.Equal(t, spaceCH2, chDeck[i](spaceCH2))
		assert.Equal(t, spaceCH3, chDeck[i](spaceCH3))
	}
}

func TestCCDeck(t *testing.T) {
	ccDeck := buildCCDeck()
	assert.Equal(t, spaceGO, ccDeck[0](spaceCC1))
	assert.Equal(t, spaceJAIL, ccDeck[1](spaceCC1))
	for i := 2; i < 16; i++ {
		assert.Equal(t, spaceCC1, ccDeck[i](spaceCC1))
		assert.Equal(t, spaceCC2, ccDeck[i](spaceCC2))
		assert.Equal(t, spaceCC3, ccDeck[i](spaceCC3))
	}
}
