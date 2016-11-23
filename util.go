package euler

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// FatalOnError will panic if err is != nil
func FatalOnError(err error, msg string, args ...interface{}) {
	if err != nil {
		log.Printf(msg, args...)
		log.Fatalf(" ERROR: %s\n", err)
	}
}

// ReadLines reads a file and splits it into lines
func ReadLines(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

// IntRange returns a list of int from 0..n-1
func IntRange(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}
	return r
}

// IntRangeN ..
func IntRangeN(start, end, incr int) []int {
	if start > end {
		start, end = end, start
	}
	d := end - start
	length := d / incr
	r := make([]int, length)
	for i := 0; i < length; i++ {
		r[i] = start + (i * incr)
	}
	return r
}

// ReverseIntSlice reverses elements of an integer slice in-place.
func ReverseIntSlice(e []int) {
	length := len(e)
	last := length - 1
	for i := 0; i < length/2; i++ {
		e[i], e[last-i] = e[last-i], e[i]
	}
}

// CopyIntSlice returns a copy of the given []int
func CopyIntSlice(e []int) []int {
	r := make([]int, len(e))
	copy(r, e)
	return r
}

// ReverseSortable the elements of the sortable
func ReverseSortable(a sort.Interface) {
	last := a.Len() - 1
	for i := 0; i < a.Len()/2; i++ {
		a.Swap(i, last-i)
	}
}
