package main

import (
	"fmt"
	"log"
	"errors"
)

func genPascalRow(n int) ([]int, error) {
	pascalRow := make([]int, n+1)
	//pascalRow := []int{1}
	pascalRow[0] = 1
	
	for k := 0; k < n; k++ {
		//fmt.Print(n,",", k, ",", pascalRow[k] * (n - k) / (k+1), " ")
		pascalRow[k+1] = pascalRow[k] * (n - k) / (k+1)
		//pascalRow = append(pascalRow, pascalRow[k] * (n - k) / (k+1))
	}
	//fmt.Println()
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
		if pascalRow, err := genPascalRow(t); err != nil {
			log.Println(err)
		} else {
			if ret, err := firstEven(pascalRow); err != nil {
				fmt.Println(t, pascalRow,ret)
			} else {
				// Need to handle the offbyone answer
				fmt.Println(t, pascalRow, ret+1)
			}
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

