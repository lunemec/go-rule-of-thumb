package main

import (
	"iter"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var (
	sizes   = []int{10, 100, 500, 1000}
	needles = []int{10, 100, 500, 1000}
)

func testingSlice(size int) []int {
	var ts = make([]int, size)
	for i := 0; i < size; i++ {
		ts[i] = rand.Intn(size)
	}
	return ts
}

func testingIter(size int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < size; i++ {
			if !yield(rand.Intn(size)) {
				return
			}
		}
	}
}

func testingString(size int) string {
	var builder strings.Builder
	for i := 0; i < size; i++ {
		builder.WriteString(strconv.Itoa(rand.Intn(size)))
	}
	return builder.String()
}

func TestIter(t *testing.T) {
	size := 10
	var accum []int

	for val := range testingIter(size) {
		accum = append(accum, val)
	}

	if len(accum) != size {
		t.Errorf("expect size of accum to be: %v, got: %v", size, len(accum))
	}
}
