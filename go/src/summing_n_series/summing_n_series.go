package main

import (
	"fmt"
	"log"
)

func NthTerm(n int) int {
	return n^2 - (n - 1)^2
}

func Sum(arr []int) int {
	v := 0
	for _, val := range arr {
		v += val
	}
	return v
}

func main() {
	testCases, err := ParseHackerRank()

	if err != nil {
		log.Fatal(err)
	}
	
	for _, t  := range testCases {
		fmt.Println(t)
	}
}


//////////////////////////////////////////////////////
// Reading/Parsing ///////////////////////////////////

func ParseInt(ref *int) error {
	_, err := fmt.Scanf("%d", ref)
	return err
}

func ParseHackerRank() ([]int, error) {
	var TestCount int
	if err := ParseInt(&TestCount); err != nil {
		return nil, err
	}
	// Alloc memory for test cases and fill
	TestCases := make([]int, TestCount)
	for i, _ := range TestCases {
		if err := ParseInt(&TestCases[i]); err != nil {
			log.Println(err)
		}
	}
	return TestCases, nil
}

