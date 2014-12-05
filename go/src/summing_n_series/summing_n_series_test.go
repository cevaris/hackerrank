package main

import (
	"testing"
	"math/big"
	"fmt"
)

func BenchmarkPow64_10K(b *testing.B) {
	test := int64(1e4)
	actual := Pow64(test,int64(2))
	expected := big.NewInt(int64(1e8))
	if expected.Cmp(actual) != 0 {
		b.Error("[%v] -> (%v) != (%v)",test, actual, expected)
	}
}

func BenchmarkPow64_100K(b *testing.B) {
	test := int64(1e5)
	actual := Pow64(test,int64(2))
	expected := big.NewInt(int64(1e10))

	if expected.Cmp(actual) != 0 {
		b.Error("[%v] -> (%v) != (%v)",test, actual, expected)
	}
}

func BenchmarkPow64_1e17(b *testing.B) {
	test := int64(1e17)
	actual := Pow64(test,int64(2))
	expected := new(big.Int)
	fmt.Sscan("10000000000000000000000000000000000", expected)
	
	if expected.Cmp(actual) != 0 {
		b.Error("[%v] -> (%v) != (%v)",test, actual, expected)
	}
}

func BenchmarkCalcSum10k(b *testing.B) {
	test := int64(1e4)
	for i := 0; i < b.N; i++ {
		CalcSum(test)
	}
}

func BenchmarkCalcSum1M(b *testing.B) {
	test := int64(1e6)
	for i := 0; i < b.N; i++ {
		CalcSum(test)
	}
}
