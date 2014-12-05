package main

import (
	"testing"
	"math/big"
	"fmt"
)

func BenchmarkPow10K(b *testing.B) {
	test := int(1e4)
	actual := Pow(test,2)
	expected := big.NewInt(int64(1e8))
	if expected.Cmp(actual) != 0 {
		b.Error("[%v] -> (%v) != (%v)",test, actual, expected)
	}
}

func BenchmarkPow100K(b *testing.B) {
	test := int(1e5)
	actual := Pow(test,2)
	expected := big.NewInt(int64(1e10))

	if expected.Cmp(actual) != 0 {
		b.Error("[%v] -> (%v) != (%v)",test, actual, expected)
	}
}

func BenchmarkPow1e17(b *testing.B) {
	test := int(1e17)
	actual := Pow(test,2)
	expected := new(big.Int)
	fmt.Sscan("10000000000000000000000000000000000", expected)
	
	if expected.Cmp(actual) != 0 {
		b.Error("[%v] -> (%v) != (%v)",test, actual, expected)
	}
}
