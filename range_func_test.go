package main

import (
	"fmt"
	"testing"
)

var Accum int

func BenchmarkRangeFunc(b *testing.B) {
	for _, size := range sizes {
		for _, iterations := range sizes {
			b.Run(
				fmt.Sprintf("slice(%d) iterations(%d)", size, iterations),
				benchmarkSliceIterate(size, iterations),
			)
			b.Run(
				fmt.Sprintf("iter func(%d) iterations(%d)", size, iterations),
				benchmarkRangeFuncIterate(size, iterations))
		}
	}
}

func benchmarkSliceIterate(size, iterations int) func(*testing.B) {
	return func(b *testing.B) {
		var acc int

		for n := 0; n < b.N; n++ {
			for i := 0; i <= iterations; i++ {
				// Here we have to allocate the iteration slice
				// as that is the main benefit of range over func
				// - no upfront allocation.
				iterSlice := testingSlice(size)
				for _, val := range iterSlice {
					acc += val
				}
			}
		}
		Accum = acc
	}
}

func benchmarkRangeFuncIterate(size, iterations int) func(*testing.B) {
	return func(b *testing.B) {
		var acc int

		for n := 0; n < b.N; n++ {
			for i := 0; i <= iterations; i++ {
				for val := range testingIter(size) {
					acc += val
				}
			}
		}
		Accum = acc

	}
}
