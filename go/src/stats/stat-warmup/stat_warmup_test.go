package main

import (
	"testing"
)

func TestModeFinder(t *testing.T) {
	var d map[int]int = map[int]int{1:1,2:1,3:1,4:1}
	var s *Stats = createStats()
	var expected int64 = 1
	var actual = s.FindMode(d)
	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}
}
