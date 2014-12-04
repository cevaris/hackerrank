package main

import (
	"fmt"
	"log"
)


func main() {
	testCases, _ := parseHackerRank()

	for _, t  := range testCases {
		if t % 2 == 0 {
			if (t / 2) % 2 == 0 {
				fmt.Println(3)
			} else {
				fmt.Println(4)
			}
		} else {
			fmt.Println(2)
		}
	}
}


func init(){
	// Enable line numbers in output
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}


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

