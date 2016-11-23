package euler059

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler059 *.txt

func init() {
	euler.Register("059", "XOR decryption", solve)
}

func solve(ctx *euler.Context) {
	ciphertext := readCiphertext()

	if euler.Verbose {
		// For fun, try some other potential passphrase lengths
		for n := 2; n < 6; n++ {
			decrypt(n, ciphertext)
		}
	}

	// Passphrase is known to be length 3
	sum := decrypt(3, ciphertext)
	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

var (
	// Most frequently occurring English characters, in order of occurrence.
	// Space occurs slightly more frequently than 'e' in texts.
	// Letter frequencies from analysis of 9,400+ English texts
	// from Project Gutenberg 2003.
	lowercase = []byte{' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r'}

	// Word/character sequences likely to appear in English
	words = map[string]int{
		"the": 1, "of": 1, "be": 1, "for": 1, "a": 1, "and": 1, "not": 1,
		"to": 1, "is": 1, "as": 1,
	}
)

// Simple scoring system, just an aggregate count of occurrances with
// penalties for non-ASCII characters
func score(text string) int {
	score := 0

	// Subtract score for any invalid ASCII characters, e.g. NULL, controls, etc
	length := len(text)
	for i := 0; i < length; i++ {
		ch := text[i]
		if ch <= 0x08 || (ch >= 0x0e && ch <= 0x1f) || ch > 0x7E {
			score--
		}
	}

	// Add score for any common English words
	fragments := strings.Split(text, " ")
	for i := 0; i < len(fragments); i++ {
		frag := fragments[i]
		score += words[frag]
	}
	return score
}

// Decipher the ciphertext using the passphrase
func decipher(passphrase, ciphertext []byte) string {
	plaintext := []byte{}
	for i := 0; i < len(ciphertext); i++ {
		idx := i % len(passphrase)
		ch := passphrase[idx] ^ ciphertext[i]
		plaintext = append(plaintext, ch)
	}
	return string(plaintext)
}

func decrypt(passlen int, ciphertext []byte) int {
	// Total characters in the ciphertext
	total := len(ciphertext)

	// Gather frequencies for characters in the ciphertext.  We keep these
	// in N separate groups, one for each letter in the passphrase.
	freqs := make([]map[byte]int, passlen)
	i := 0
	for i < passlen {
		freqs[i] = make(map[byte]int)
		i++
	}

	// Compute the letter frequencies grouped by character in the passphrase
	for i = 0; i < total; i++ {
		c := ciphertext[i]
		idx := i % passlen
		freqs[idx][c]++
	}

	// Convert the frequencies into a slice of letter{character, frequency}
	inverse := []letters{}
	for _, freq := range freqs {
		// Associate each letter with its frequency and append to a slice
		o := letters{}
		for c, f := range freq {
			o = append(o, letter{c, f})
		}
		// Reverse sort the letters slice by frequency
		sort.Sort(sort.Reverse(o))
		inverse = append(inverse, o)
	}

	// Decode the letters by frequency and collect these candidates in N groups,
	// one for each character in the passphrase
	grouped := make([][]byte, passlen)
	for _, c := range lowercase {
		for j := 0; j < passlen; j++ {
			for k, inv := range inverse {
				ch := inv[j].c ^ c
				grouped[k] = append(grouped[k], ch)
			}
		}
	}

	// Generate the indices for combinations of candidate letters. We need
	// a list of indices of length `passlen` each having a base of `len(lowercase)`
	bases := []int{}
	i = passlen
	for i > 0 {
		bases = append(bases, len(lowercase))
		i--
	}

	// Build up candidate phrases by selecting a letter from each group
	indices := euler.NewMixedRadixCounter(bases).All()
	phrases := [][]byte{}
	seen := make(map[string]*struct{})
	for _, index := range indices {
		phrase := []byte{}
		// fmt.Println(index, grouped)
		for x, y := range index {
			phrase = append(phrase, grouped[x][y])
		}

		key := string(phrase)
		if seen[key] == nil {
			seen[key] = &struct{}{}
			phrases = append(phrases, phrase)
		}
	}

	// Decipher and score the texts by checking if they look like English
	scored := texts{}
	for _, phrase := range phrases {
		result := decipher(phrase, ciphertext)
		score := score(result)
		scored = append(scored, text{result, string(phrase), score})
		if euler.Verbose && score > 0 {
			fmt.Printf("Tried '%s' score %d ..\n", phrase, score)
		}
	}

	// Reverse-sorting the plaintexts by score
	sort.Sort(sort.Reverse(scored))

	// Use the highest score as our best guess at the answer
	if euler.Verbose {
		fmt.Printf("Passphrase length %d: '%v' produced highest score of %d\n",
			passlen, scored[0].phrase, scored[0].score)
		fmt.Println(scored[0].text)
		fmt.Println("------------------------------------------------")
	}

	// Compute the answer by summing the ascii character values in the plaintext
	answer := 0
	for _, c := range []byte(scored[0].text) {
		answer += int(c)
	}
	return answer
}

func readCiphertext() []byte {
	res := []byte{}
	input := strings.TrimSpace(string(rawfiles["p059_cipher.txt"]))
	for _, s := range strings.Split(input, ",") {
		code, err := strconv.ParseInt(s, 10, 8)
		failIf(err)
		res = append(res, byte(code))
	}
	return res
}

func failIf(err error) {
	if err != nil {
		log.Fatalln("unexpected error", err)
	}
}

type letter struct {
	c byte
	f int
}

func (o letter) String() string {
	return fmt.Sprintf("0x%x %d", o.c, o.f)
}

type letters []letter

func (o letters) Len() int { return len(o) }

func (o letters) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

func (o letters) Less(i, j int) bool { return o[i].f < o[j].f }

type text struct {
	text   string
	phrase string
	score  int
}

type texts []text

func (t texts) Len() int { return len(t) }

func (t texts) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (t texts) Less(i, j int) bool { return t[i].score < t[j].score }
