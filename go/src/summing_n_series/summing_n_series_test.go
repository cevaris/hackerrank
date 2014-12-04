package main

import (
	"testing"
)


func IntArray(n int) []int {
	arr := make([]int, n+1)
	for i :=0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func TestSum(t *testing.T) {

	test     := []int{1,2,3,4,5}
	expected := 15
	actual   := Sum(test)
	if expected != actual {
		t.Errorf("(%v) != (%v)", actual, expected)
	}
	
}

func TestSumEmpty(t *testing.T) {

	test     := []int{}
	expected := 0
	actual   := Sum(test)
	if expected != actual {
		t.Errorf("(%v) != (%v)", actual, expected)
	}
	
}


func BenchmarkSum10K(b *testing.B) {
	arr := IntArray(10000)
	for i := 0; i < b.N; i++ {
		Sum(arr)
	}
}

func BenchmarkSum100K(b *testing.B) {
	arr := IntArray(100000)
	for i := 0; i < b.N; i++ {
		Sum(arr)
	}
}

func BenchmarkSum1M(b *testing.B) {
	arr := IntArray(1000000)
	for i := 0; i < b.N; i++ {
		Sum(arr)
	}
}
