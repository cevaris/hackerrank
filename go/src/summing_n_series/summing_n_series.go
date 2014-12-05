package main

import (
	"fmt"
	"log"
)

const resMod = 10e9 + 7

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}


func Sum(arr []int) int {
	v := 0
	for _, val := range arr {
		v += val
	}
	return v
}

func NthTerm(n int) int {
	return Pow(n,2) - Pow(n - 1,2)
}

func GenArray(n int) []int {
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = NthTerm(i)
	}
	return arr
}

func main() {
	testCases, err := ParseHackerRank()

	if err != nil {
		log.Fatal(err)
	}
	
	for _, t  := range testCases {
		arr := GenArray(t)
		sum := Sum(arr)
		res := sum % resMod
		fmt.Println(res)
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

