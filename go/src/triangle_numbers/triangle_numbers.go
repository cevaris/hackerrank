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
)

func StdinReader() (*bufio.Reader) {
	return bufio.NewReader(os.Stdin)
}
// func parseIntLine()

// func parseHackerRank() 

func main() {
	reader := StdinReader()
	//reader := bufio.NewReader(os.Stdin)
	//defer reader.Close()

	// Initial read of test case count
	testCountStr, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err,"Error Reading Test Count")
	}

	testCount, err := strconv.Atoi(string(testCountStr))
	//testCount, err := binary.ReadVarint(testCountStr)
	if err != nil {
		log.Fatal(err,"Error Converting Test Count to Int")
	}

	

	testCases := make([]int, testCount)

	fmt.Println("test Count", testCount, testCases)

	// for n := range testCount {
	// 	fmt.Println(scanner.Text())
	// }

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	
	// reader, err := bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(reader)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// content, err := ioutil.ReadAll(os.Stdin)
	// if err == nil {
	// 	log.Fatal(err)
	// }
	// lines := strings.Split(string(content), "\n")

	// fmt.Println("Input", lines)


	// pascalRow := []int{1}
	// fmt.Println("Data", pascalRow)
}
