package main

import (
	"testing"
)

func assert(truth bool, msg string) {
	if !truth {
		panic(msg)
	}
}

func BenchmarkAssertion(b *testing.B) {
	b.Run(
		"no assert",
		benchmarkNoAssert,
	)
	b.Run(
		"assert",
		benchmarkAssert,
	)
	b.Run(
		"assert(5)",
		benchmarkAssert5,
	)
	b.Run(
		"defer assert",
		benchmarkDeferAssert,
	)
}

var (
	Truth  = true
	Truth2 = true
	Truth3 = true
	Truth4 = true
	Truth5 = true
	Out    struct{}
)

func benchmarkNoAssert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		func() {
			Out = struct{}{}
		}()
	}
}

func benchmarkAssert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		func() {
			assert(Truth, "n must be larger than 0")
			Out = struct{}{}
		}()
	}
}

func benchmarkAssert5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		func() {
			assert(Truth, "n must be larger than 0")
			assert(Truth2, "n must be larger than 0")
			assert(Truth3, "n must be larger than 0")
			assert(Truth4, "n must be larger than 0")
			assert(Truth5, "n must be larger than 0")
			Out = struct{}{}
		}()
	}
}

func benchmarkDeferAssert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		func() {
			defer assert(Truth, "n must be larger than 0")
			Out = struct{}{}
		}()
	}
}
