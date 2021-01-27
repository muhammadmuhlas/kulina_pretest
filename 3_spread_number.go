package main

import (
	"strconv"
	"math"
	"fmt"
)

// SpreadNumber will spread number into decimal
// First, slice an int and then for each element
// multiply with power of ten times current length respectively

func SpreadNumber(n int) {
	s := strconv.Itoa(n)
	for k, v := range []rune(s) {
		ss, _ := strconv.Atoi(string(v))
		fmt.Printf("%.f\n", float64(ss) * math.Pow(10, float64(len(s)-k-1)))
	}
}
