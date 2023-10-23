package main

import (
	"math/rand"
	"strconv"
	"strings"
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

func testingString(size int) string {
	var builder strings.Builder
	for i := 0; i < size; i++ {
		builder.WriteString(strconv.Itoa(rand.Intn(size)))
	}
	return builder.String()
}
