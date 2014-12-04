package main

import (
	"fmt"
	"log"
)

func firstEven(t int) int {
	switch {
	case t < 3:
		return -1
	case t % 2 != 0:
		return 2
	case (t / 2) % 2 == 0:
		return 3
	default:
		return 4
	}
}

func main() {
	testCases, err := parseHackerRank()

	if err != nil {
		log.Fatal(err)
	}
	
	for _, t  := range testCases {
		fmt.Println(firstEven(t))
	}
}


//////////////////////////////////////////////////////
// Reading/Parsing ///////////////////////////////////

func ParseInt(ref *int) error {
	_, err := fmt.Scanf("%d", ref)
	return err
}

func parseHackerRank() ([]int, error) {
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

