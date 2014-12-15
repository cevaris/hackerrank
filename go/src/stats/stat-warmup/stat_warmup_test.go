package main

import (
	"testing"
)

func TestModeFinder(t *testing.T) {
	var s *Stats = createStats()
	var expected, actual int64
	var d map[int]int

	testCondition :=  func(label string, actual, expected int64) {
		if actual != expected {
			t.Error(label,"Expected", expected, "Actual", actual)
		}		
	}

	d = map[int]int{1:1,2:1,3:1,4:1}	
	expected = 1
	actual = s.FindMode(d)
	testCondition("All equal freq",actual, expected)

	d = map[int]int{1:1,2:4,3:1,4:1}	
	expected = 2
	actual = s.FindMode(d)
	testCondition("Single max freq",actual, expected)

	d = map[int]int{1:1,2:1,3:2,4:2}	
	expected = 3
	actual = s.FindMode(d)
	testCondition("Competing max freq",actual, expected)
}
