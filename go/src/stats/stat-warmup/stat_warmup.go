package main

import (
	"fmt"
	"strings"
	"log"
	"strconv"
	"bufio"
	"os"
)

func main() {
	testCases, err := parseHackerRank()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(testCases)

	// for _, t  := range testCases {
	// 	fmt.Println(firstEven(t))
	// }
}

//////////////////////////////////////////////////////
// Reading/Parsing ///////////////////////////////////

func ParseArray(ref []int) error {
	// Read array line
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		return err
	}
	// Separate each value
	strs := strings.Split(string(line)," ")

	// For each value, split
	for i,v := range strs {
		if res, err := strconv.Atoi(v); err != nil {
			return err
		} else {
			ref[i] = res 
		}
	}

	return nil	
}

func ParseInt(ref *int) error {
	_, err := fmt.Scanf("%d", ref)
	return err
}

func parseHackerRank() ([]int, error) {
	var testCount int
	if err := ParseInt(&testCount); err != nil {
		return nil, err
	}
	// Alloc memory for test cases and fill
	testCases := make([]int, testCount)
	ParseArray(testCases)
	return testCases, nil
}
