package main

/*
Problem Statement

Given a triangle of numbers where each number is equal to the sum of the three numbers on top of it, find the first even number in a row.

Explanatory Note: The vertex of the triangle (at the top) is 1. The structure of the triangle is shown below. Each number is equal to the sum of the numbers at the following positions: Position X: immediately above it. Position Y: to the immediate left of X. Position Z: to the immediate right of X. If there are no numbers at positions X, Y, or Z, then assume those positions are occupied by a zero (0). This can occur for positions at the edge.

Here are four rows of the triangle:

          1
      1   1  1
   1  2   3  2  1
1  3  6   7  6  3  1

Input Format and Constraints
First line contains a number T (number of test cases).
Each of the next T lines contain a number N (the row number, assuming that the top vertex of the triangle is Row 1).

Output Format
For each test case, display an integer that denotes the position of the first even number.
Note: Assume that the left most number in a row is Position 1.
If there is no even number in a row, print -1.

Constraints
1<=T<=100
3<=N<=1000000000

Sample Input
2
3
4

Sample Output
2
3
*/


import (
	"fmt"
	"os"
	"log"
	// "strings"
	"bufio"
	// "io/ioutil"
	"strconv"
	//"encoding/binary"
	"errors"
)

/*
>>> def pascal(n):
...   line = [1]
...   for k in range(n):
...     line.append(line[k] * (n-k) / (k+1))
...   return line
*/
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

func init(){
	// Enable line numbers in output
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	testCases, _ := parseHackerRank()
	//fmt.Println("test Count", testCases, err)

	for _, t  := range testCases {
		if pascalRow, err := genPascalRow(t+1); err != nil {
			log.Println(err)
		} else {
			if ret, err := firstEven(pascalRow); err != nil {
				fmt.Println(pascalRow, ret)
			} else {
				fmt.Println(pascalRow, ret+1)
			}
		}
	}
	
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
