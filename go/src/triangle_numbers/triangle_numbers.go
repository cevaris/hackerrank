package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"errors"
)

func genPascalRow(n int) ([]int, error) {
	pascalRow := make([]int, n+1)
	pascalRow[0] = 1
	
	for k := 0; k < n; k++ {
		pascalRow[k+1] = pascalRow[k] * (n - k) / (k+1)
	}
	
	return pascalRow, nil
}

func firstEven(arr []int) (int, error) {
	for i, val := range arr {
		if val % 2 == 0 {
			return i, nil
		}
	}
	return -1, errors.New("Even value not found")
}


func main() {
	testCases, _ := parseHackerRank()

	for _, t  := range testCases {
		if pascalRow, err := genPascalRow(t+1); err != nil {
			log.Println(err)
		} else {
			// Need to handle the offbyone answer
			if ret, err := firstEven(pascalRow); err != nil {
				fmt.Println(ret)
			} else {
				fmt.Println(ret+1)
			}
		}
	}
	
}


func init(){
	// Enable line numbers in output
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}


func StdinReader() (*bufio.Reader) {
	return bufio.NewReader(os.Stdin)
}

func parseIntLine() (int, error) {
	reader := StdinReader()
	numStr, _, err := reader.ReadLine()
	if err != nil {
		return 0, err
	}

	numVal, err := strconv.Atoi(string(numStr))
	if err != nil {
		return 0, err
	}
	
	return numVal, nil
}

func parseHackerRank() ([]int, error) {
	testCount, err := parseIntLine()
	if err != nil {
		return nil, err
	}
	// Alloc memory for test cases and fill
	testCases := make([]int, testCount)
	for i, _ := range testCases {
		if testCases[i], err = parseIntLine(); err != nil {
			log.Println(err)
		}
	}
	return testCases, nil
}

