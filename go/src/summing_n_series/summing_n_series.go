package main

import (
	"fmt"
	"log"
	"math/big"
)

func Pow(a, b int) *big.Int {
	var m *big.Int
	x := big.NewInt(int64(a))
	y := big.NewInt(int64(b))
	return new(big.Int).Exp(x, y, m)
}

func CalcSum(n int) *big.Int {
	res := big.NewInt(0)
	for i := 1; i <= n; i++ {
		res.Sub(res,Pow(i-1,2)).Add(res,Pow(i,2))
	}
	return res
}

func main() {
	testCases, err := ParseHackerRank()

	if err != nil {
		log.Fatal(err)
	}
	
	resMod := big.NewInt(int64(10e9 + 7))
	for _, t  := range testCases {
		res := CalcSum(t)
		fmt.Println(res.Mod(res, resMod))
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

