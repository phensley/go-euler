# go-euler - Project Euler solutions in Go

## Usage

```
$ go-euler
Usage: go-euler {all | ###}
```

```
$ go-euler -h
Usage of go-euler:
  -d, --debug            debug mode
  -e, --elapsed float    show answers that took longer than N.NNN seconds to produce
  -p, --profile string   enable cpu profiling
  -r, --reveal           reveal the answer
  -v, --verbose          verbose mode
```

For certain problems verbose mode (`-v`) will show more detail on the answer, and
confirm the knowns, if any.  It may also increase execution time substantially,
as in the case of problem 54 "Poker hands" where poker odds are computed by
generating several hundred thousand random hands.

```
$ go-euler 054
054      ok      0.007  "Poker hands"
```


```
$ go-euler 054 -v
3H 3D 4D TH KD (one pair)    beats    2S 5C 8H 9H KC (no pair)
5S 7H TC JC AH (no pair)    beats    6C 7C 9C QD KD (no pair)
2C 2H 7C 9S JH (one pair)    beats    5D 6H 9C TS KH (no pair)
3D 5C 7S 7C KD (one pair)    beats    2S 7D JS QD KS (no pair)

.... snip ....

Confirming poker hand odds, generating 250000 random hands ...

.... snip ....

Hand frequencies:
50.38080%       125952   no pair
42.10200%       105255   one pair
 4.69800%        11745   two pairs
 2.11280%         5282   three of a kind
 0.35120%          878   straight
 0.19440%          486   flush
 0.13840%          346   full house
 0.02080%           52   fourOfAKind
 0.00160%            4   straightFlush

054      ok      1.659  "Poker hands"
```

## Solutions

Solving the first 100 problems takes just under 10 seconds.

Depending on my level of interest I've taken a deeper dive on some problems and
not others. Some frequently-reused routines have been developed and optimized
as needed. Not everything has been optimized yet but execution times are
acceptable.

To run all of the problems, use `go-euler all`.  For example, display solutions
which take >= 100 milliseconds to run:

```
$ time go-euler all -e .1
012      ok      0.357  "Highly divisible triangular number"
014      ok      0.314  "Longest Collatz sequence"
023      ok      0.134  "Non-abundant sums"
027      ok      0.382  "Quadratic primes"
034      ok      0.252  "Digit factorials"
043      ok      0.122  "Sub-string divisibility"
047      ok      0.162  "Distinct prime factors"
057      ok      0.136  "Square root convergents"
060      ok      0.221  "Prime pair sets"
070      ok      1.553  "Totient permutation"
073      ok      0.884  "Counting fractions in a range"
074      ok      1.595  "Digit factorial chains"
088      ok      0.233  "Product-sum numbers"
092      ok      0.742  "Square digit chains"
095      ok      0.399  "Amicable chains"
098      ok      0.218  "Anagramic squares"

real    0m8.944s
user    0m9.464s
sys     0m0.346s
```

