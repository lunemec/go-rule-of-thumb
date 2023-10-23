package main

import (
	"fmt"
	"testing"
)

var Append []int

type appendBench func(first []int, second []int) []int

func BenchmarkAppend(b *testing.B) {
	for _, sizeFirst := range sizes {
		for _, sizeSecond := range sizes {
			b.Run(
				fmt.Sprintf("append_expand(%d)(%d)", sizeFirst, sizeSecond),
				benchmarkAppend(sizeFirst, sizeSecond, appendExpand),
			)
			b.Run(
				fmt.Sprintf("append_for(%d)(%d)", sizeFirst, sizeSecond),
				benchmarkAppend(sizeFirst, sizeSecond, appendFor),
			)

		}
	}
}

func benchmarkAppend(sizeFirst, sizeSecond int, runF appendBench) func(*testing.B) {
	first := testingSlice(sizeFirst)
	second := testingSlice(sizeSecond)

	return func(b *testing.B) {
		var f []int

		for n := 0; n < b.N; n++ {
			f = runF(first, second)
		}
		Append = f
	}
}

func appendExpand(first, second []int) []int {
	return append(first, second...)
}

func appendFor(first, second []int) []int {
	for _, secondVal := range second {
		first = append(first, secondVal)
	}
	return first
}
