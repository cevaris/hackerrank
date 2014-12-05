package main

import (
	"fmt"
	"log"
	"math/big"
)

func Pow64(a, b int64) *big.Int {
	var m *big.Int
	x := big.NewInt(a)
	y := big.NewInt(b)
	return new(big.Int).Exp(x, y, m)
}

func CalcSum(n int64) *big.Int {
	return Pow64(n,2)
}

func main() {
	testCases, err := ParseHackerRank()

	if err != nil {
		log.Fatal(err)
	}
	
	resMod := big.NewInt(int64(1e9 + 7))
	for _, t  := range testCases {
		res := CalcSum(t)
		fmt.Println(res.Mod(res, resMod))
	}
}


//////////////////////////////////////////////////////
// Reading/Parsing ///////////////////////////////////

func ParseInt64(ref *int64) error {
	_, err := fmt.Scanf("%d", ref)
	return err
}

func ParseHackerRank() ([]int64, error) {
	var TestCount int64
	if err := ParseInt64(&TestCount); err != nil {
		return nil, err
	}
	// Alloc memory for test cases and fill
	TestCases := make([]int64, TestCount)
	for i, _ := range TestCases {
		if err := ParseInt64(&TestCases[i]); err != nil {
			log.Println(err)
		}
	}
	return TestCases, nil
}

