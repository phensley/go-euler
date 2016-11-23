package euler098

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler098 *.txt

func init() {
	euler.Register("098", "Anagramic squares", solve)
}

func solve(ctx *euler.Context) {
	// Map words to their anagram, if any.
	words := readWords()
	anagrams := anagramlist{}
	for i, w1 := range words {
		for _, w2 := range words[i+1:] {
			if w1.length != w2.length {
				continue
			}

			// The two words are anagrams of one another, link them up.
			if w1.sorted == w2.sorted {
				index := indexAnagram(&w1, &w2)
				anagrams = append(anagrams, anagram{w1, w2, index})
			}
		}
	}

	// Given the maximum length of a word, compute the largest
	// possible square we'll see
	limit := float64(euler.SmallFactorial(9))
	sqrt := uint64(math.Sqrt(limit))
	squares := euler.NewBitString(uint64(limit))
	for n := uint64(2); n <= sqrt; n++ {
		squares.Set(n * n)
	}

	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	largest := 0

	sort.Sort(anagrams)
	for _, a := range anagrams {

		length := a.w1.length
		perm2 := make([]int, length)

		// Check all digit permutations of length == to our word length.
		// Once we find one that is square, we encode the same digits into
		// our anagram and check if it is square.
		perms := euler.NewPartialPermutations(digits, a.w1.length)
		for perms.Next() {
			perm1 := perms.Get()

			// Skip those with leading digit zero.
			if perm1[0] == 0 {
				continue
			}

			// If this permutation is not a square, skip it.
			p := euler.DigitsToInt(perm1)
			if !squares.IsSet(uint64(p)) {
				continue
			}

			// Use the word's index, which encodes the digit swaps
			// transforming a word into its anagram.
			for i, j := range a.index {
				perm2[j] = perm1[i]
			}

			// Skip leading digit zero
			if perm2[0] == 0 {
				continue
			}

			// Check if the anagram's digits are square
			q := euler.DigitsToInt(perm2)
			if !squares.IsSet(uint64(q)) {
				continue
			}

			// Check if either of the squares produced is the largest we've seen.
			if p > largest {
				largest = p
			}
			if q > largest {
				largest = q
			}
			if euler.Verbose {
				fmt.Printf("%s %d  ->  %s %d\n", a.w1.normal, p, a.w2.normal, q)
			}
		}
	}

	answer := fmt.Sprintf("%d", largest)
	ctx.SetAnswer(answer)
}

// Create an index that maps the character positions in the
// first word to those in the second.
func indexAnagram(w1, w2 *word) []int {
	length := w1.length
	r := make([]int, length)
	for i := 0; i < length; i++ {
		r[i] = -1
	}

	// Map the index of each letter swap transforming W1 into W2.
	//  ACT -> CAT  would have indices [1,0,2]
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if w1.normal[i] == w2.normal[j] && r[i] == -1 {
				r[i] = j
				break
			}
		}
	}
	if euler.Verbose {
		fmt.Printf("%s -> %s  swap indices:  %v\n", w1.normal, w2.normal, r)
	}
	return r
}

func readWords() wordlist {
	raw := strings.Split(string(rawfiles["p098_words.txt"]), ",")
	words := make(wordlist, len(raw))
	for i, w := range raw {
		w = w[1 : len(w)-1]
		b := bytes(w)
		sort.Sort(b)
		words[i] = word{len(w), w, string(b)}
	}
	sort.Sort(words)
	return words
}

type anagram struct {
	w1    word
	w2    word
	index []int
}

// For debugging
func (a anagram) String() string {
	return fmt.Sprintf("w1=%s w2=%s index=%v", a.w1.normal, a.w2.normal, a.index)
}

type word struct {
	length int
	normal string
	sorted string
}

// Lots of sorting for this problem

type wordlist []word

func (w wordlist) Len() int { return len(w) }

func (w wordlist) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func (w wordlist) Less(i, j int) bool { return w[i].sorted < w[j].sorted }

type bytes []byte

func (b bytes) Len() int { return len(b) }

func (b bytes) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b bytes) Less(i, j int) bool { return b[i] < b[j] }

type anagramlist []anagram

func (a anagramlist) Len() int { return len(a) }

func (a anagramlist) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a anagramlist) Less(i, j int) bool { return a[i].w1.length < a[j].w1.length }
