package main

import (
	"fmt"
	"sort"
	"testing"
)

var Deduplicated []int

type deduplicateBench func(haystack []int) []int

func BenchmarkDeduplication(b *testing.B) {
	for _, size := range sizes {
		b.Run(
			fmt.Sprintf("slice(%d)", size),
			benchmarkDeduplicate(size, benchDeduplicateSlice),
		)
		b.Run(
			fmt.Sprintf("map(%d)", size),
			benchmarkDeduplicate(size, benchDeduplicateMap))
	}
}

func benchmarkDeduplicate(size int, runF deduplicateBench) func(*testing.B) {
	haystack := testingSlice(size)
	return func(b *testing.B) {
		var f []int

		for n := 0; n < b.N; n++ {
			// Note: we need to copy to prevent runs pre-sorting
			// arrays for each other.
			h := make([]int, len(haystack))
			copy(h, haystack)
			f = runF(h)
		}
		Deduplicated = f
	}
}

func benchDeduplicateSlice(haystack []int) []int {
	// "borrowed" from https://github.com/golang/go/wiki/SliceTricks, thanks!
	// Note sort + slices.Compact is the same thing.
	sort.Ints(haystack)
	j := 0
	for i := 1; i < len(haystack); i++ {
		if haystack[j] == haystack[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		haystack[j] = haystack[i]
	}
	return haystack[:j+1]
}

func benchDeduplicateMap(haystack []int) []int {
	result := make([]int, 0, len(haystack))
	seen := make(map[int]struct{}, len(haystack))

	for _, item := range haystack {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}
