package main

import (
	"fmt"
	"log"
	//"reflect"
	"math/big"
)

const resMod = 10e9 + 7

// func Pow(a, b int) int {
// 	p := 1
// 	for b > 0 {
// 		if b&1 != 0 {
// 			p *= a
// 		}
// 		b >>= 1
// 		a *= a
// 	}
// 	return p
// }
func Pow(a, b int) *big.Int {
	x := big.NewInt(int64(a))
	y := big.NewInt(int64(b))
	return x.Exp(x, y, nil)
}

func Sum(arr []int) int {
	v := 0
	for _, val := range arr {
		v += val
	}
	return v
}

// func NthTerm(n int) *big.Int {
// 	//return Pow(n,2) - Pow(n - 1,2)
// 	return big.NewInt(3)
// }

//// Need to make iterative, cannot store in memory
// func GenArray(n int) *[]big.Int {
// 	arr := make([]big.Int, n+1)
// 	for i := 1; i <= n; i++ {
// 		arr[i] = NthTerm(i)
// 	}
// 	return arr
// }

func CalcSum(n int) *big.Int {
	res := big.NewInt(0)
	for i := 1; i <= n; i++ {
		// x := res.Add(-Pow(-1+i,2)) + Pow(i,2)
		res.Sub(res,Pow(-1+i,2)).Add(res,Pow(i,2))
		//res.Add(res, big.NewInt(int64(x)))
	}
	return res
}


func main() {
	testCases, err := ParseHackerRank()

	if err != nil {
		log.Fatal(err)
	}

	for _, t  := range testCases {
		//fmt.Println(t,reflect.TypeOf(t),CalcSum(t))
		fmt.Println(CalcSum(t))
	}
	
	// for _, t  := range testCases {
	// 	arr := CalcSum(t)
	// 	sum := Sum(arr)
	// 	res := sum % resMod
	// 	fmt.Println(res)
	// }
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

